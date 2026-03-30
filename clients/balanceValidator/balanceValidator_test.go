package balanceValidator

import (
	"context"
	"errors"
	"math/big"
	"testing"

	"github.com/TerraDharitri/drt-go-bridge-eth/clients"
	bridgeCore "github.com/TerraDharitri/drt-go-bridge-eth/core"
	"github.com/TerraDharitri/drt-go-bridge-eth/core/batchProcessor"
	"github.com/TerraDharitri/drt-go-bridge-eth/testsCommon/bridge"
	"github.com/TerraDharitri/drt-go-chain/testscommon"
	"github.com/stretchr/testify/assert"
)

var (
	ethToken = []byte("peer token")
	drtToken = []byte("drt token")
	amount   = big.NewInt(100)
	amount2  = big.NewInt(200)
)

func createMockArgsBalanceValidator() ArgsBalanceValidator {
	return ArgsBalanceValidator{
		Log:              &testscommon.LoggerStub{},
		DharitrIClient: &bridge.DharitrIClientStub{},
		PeerChainClient:  &bridge.PeerChainClientStub{},
	}
}

type testConfiguration struct {
	isNativeOnEth      bool
	isMintBurnOnEth    bool
	totalBalancesOnEth *big.Int
	burnBalancesOnEth  *big.Int
	mintBalancesOnEth  *big.Int

	isNativeOnDrt      bool
	isMintBurnOnDrt    bool
	totalBalancesOnDrt *big.Int
	burnBalancesOnDrt  *big.Int
	mintBalancesOnDrt  *big.Int

	errorsOnCalls map[string]error

	peerToken []byte
	drtToken  []byte
	amount    *big.Int
	direction batchProcessor.Direction

	lastExecutedEthBatch       uint64
	pendingDrtBatchId          uint64
	amountsOnDrtPendingBatches map[uint64][]*big.Int
	amountsOnEthPendingBatches map[uint64][]*big.Int
}

func (cfg *testConfiguration) deepClone() testConfiguration {
	result := testConfiguration{
		isNativeOnEth:              cfg.isNativeOnEth,
		isMintBurnOnEth:            cfg.isMintBurnOnEth,
		isNativeOnDrt:              cfg.isNativeOnDrt,
		isMintBurnOnDrt:            cfg.isMintBurnOnDrt,
		errorsOnCalls:              make(map[string]error),
		peerToken:                  cfg.peerToken,
		drtToken:                   make([]byte, len(cfg.drtToken)),
		direction:                  cfg.direction,
		lastExecutedEthBatch:       cfg.lastExecutedEthBatch,
		pendingDrtBatchId:          cfg.pendingDrtBatchId,
		amountsOnDrtPendingBatches: make(map[uint64][]*big.Int),
		amountsOnEthPendingBatches: make(map[uint64][]*big.Int),
	}
	if cfg.totalBalancesOnEth != nil {
		result.totalBalancesOnEth = big.NewInt(0).Set(cfg.totalBalancesOnEth)
	}
	if cfg.burnBalancesOnEth != nil {
		result.burnBalancesOnEth = big.NewInt(0).Set(cfg.burnBalancesOnEth)
	}
	if cfg.mintBalancesOnEth != nil {
		result.mintBalancesOnEth = big.NewInt(0).Set(cfg.mintBalancesOnEth)
	}
	if cfg.totalBalancesOnDrt != nil {
		result.totalBalancesOnDrt = big.NewInt(0).Set(cfg.totalBalancesOnDrt)
	}
	if cfg.burnBalancesOnDrt != nil {
		result.burnBalancesOnDrt = big.NewInt(0).Set(cfg.burnBalancesOnDrt)
	}
	if cfg.mintBalancesOnDrt != nil {
		result.mintBalancesOnDrt = big.NewInt(0).Set(cfg.mintBalancesOnDrt)
	}
	if cfg.amount != nil {
		result.amount = big.NewInt(0).Set(cfg.amount)
	}

	for key, err := range cfg.errorsOnCalls {
		result.errorsOnCalls[key] = err
	}
	copy(result.drtToken, cfg.drtToken)
	for nonce, values := range cfg.amountsOnDrtPendingBatches {
		result.amountsOnDrtPendingBatches[nonce] = make([]*big.Int, 0, len(values))
		for _, value := range values {
			result.amountsOnDrtPendingBatches[nonce] = append(result.amountsOnDrtPendingBatches[nonce], big.NewInt(0).Set(value))
		}
	}
	for nonce, values := range cfg.amountsOnEthPendingBatches {
		result.amountsOnEthPendingBatches[nonce] = make([]*big.Int, 0, len(values))
		for _, value := range values {
			result.amountsOnEthPendingBatches[nonce] = append(result.amountsOnEthPendingBatches[nonce], big.NewInt(0).Set(value))
		}
	}

	return result
}

type testResult struct {
	checkRequiredBalanceOnEthCalled bool
	checkRequiredBalanceOnDrtCalled bool
	error                           error
}

func TestNewBalanceValidator(t *testing.T) {
	t.Parallel()

	t.Run("nil logger should error", func(t *testing.T) {
		t.Parallel()

		args := createMockArgsBalanceValidator()
		args.Log = nil
		instance, err := NewBalanceValidator(args)
		assert.Nil(t, instance)
		assert.Equal(t, ErrNilLogger, err)
	})
	t.Run("nil DharitrI client should error", func(t *testing.T) {
		t.Parallel()

		args := createMockArgsBalanceValidator()
		args.DharitrIClient = nil
		instance, err := NewBalanceValidator(args)
		assert.Nil(t, instance)
		assert.Equal(t, ErrNilDharitrIClient, err)
	})
	t.Run("nil Ethereum client should error", func(t *testing.T) {
		t.Parallel()

		args := createMockArgsBalanceValidator()
		args.PeerChainClient = nil
		instance, err := NewBalanceValidator(args)
		assert.Nil(t, instance)
		assert.Equal(t, ErrNilEthereumClient, err)
	})
	t.Run("should work", func(t *testing.T) {
		t.Parallel()

		args := createMockArgsBalanceValidator()
		instance, err := NewBalanceValidator(args)
		assert.NotNil(t, instance)
		assert.Nil(t, err)
	})
}

func TestBalanceValidator_IsInterfaceNil(t *testing.T) {
	t.Parallel()

	var instance *balanceValidator
	assert.True(t, instance.IsInterfaceNil())

	instance = &balanceValidator{}
	assert.False(t, instance.IsInterfaceNil())
}

func TestBridgeExecutor_CheckToken(t *testing.T) {
	t.Parallel()

	expectedError := errors.New("expected error")
	t.Run("unknown direction should error", func(t *testing.T) {
		t.Parallel()

		cfg := testConfiguration{
			direction: "",
		}
		result := validatorTester(cfg)
		assert.ErrorIs(t, result.error, ErrInvalidDirection)
	})
	t.Run("query operations error", func(t *testing.T) {
		t.Parallel()

		t.Run("on isMintBurnOnEthereum", func(t *testing.T) {
			cfg := testConfiguration{
				direction: batchProcessor.FromDharitrI,
				errorsOnCalls: map[string]error{
					"MintBurnTokensEth": expectedError,
				},
			}
			result := validatorTester(cfg)
			assert.Equal(t, expectedError, result.error)
			assert.True(t, result.checkRequiredBalanceOnEthCalled)
			assert.False(t, result.checkRequiredBalanceOnDrtCalled)
		})
		t.Run("on isMintBurnOnDharitrI", func(t *testing.T) {
			cfg := testConfiguration{
				direction: batchProcessor.ToDharitrI,
				errorsOnCalls: map[string]error{
					"IsMintBurnTokenDrt": expectedError,
				},
			}
			result := validatorTester(cfg)
			assert.Equal(t, expectedError, result.error)
			assert.False(t, result.checkRequiredBalanceOnEthCalled)
			assert.True(t, result.checkRequiredBalanceOnDrtCalled)
		})
		t.Run("on isNativeOnEthereum", func(t *testing.T) {
			cfg := testConfiguration{
				direction: batchProcessor.ToDharitrI,
				errorsOnCalls: map[string]error{
					"NativeTokensEth": expectedError,
				},
			}
			result := validatorTester(cfg)
			assert.Equal(t, expectedError, result.error)
			assert.False(t, result.checkRequiredBalanceOnEthCalled)
			assert.True(t, result.checkRequiredBalanceOnDrtCalled)
		})
		t.Run("on isNativeOnDharitrI", func(t *testing.T) {
			cfg := testConfiguration{
				direction: batchProcessor.FromDharitrI,
				errorsOnCalls: map[string]error{
					"IsNativeTokenDrt": expectedError,
				},
			}
			result := validatorTester(cfg)
			assert.Equal(t, expectedError, result.error)
			assert.True(t, result.checkRequiredBalanceOnEthCalled)
			assert.False(t, result.checkRequiredBalanceOnDrtCalled)
		})
		t.Run("on computeEthAmount, TotalBalances", func(t *testing.T) {
			cfg := testConfiguration{
				direction:       batchProcessor.FromDharitrI,
				isMintBurnOnDrt: true,
				isNativeOnEth:   true,
				errorsOnCalls: map[string]error{
					"TotalBalancesEth": expectedError,
				},
			}
			result := validatorTester(cfg)
			assert.Equal(t, expectedError, result.error)
			assert.True(t, result.checkRequiredBalanceOnEthCalled)
			assert.False(t, result.checkRequiredBalanceOnDrtCalled)
		})
		t.Run("on computeEthAmount, BurnBalances", func(t *testing.T) {
			cfg := testConfiguration{
				direction:       batchProcessor.FromDharitrI,
				isNativeOnDrt:   true,
				isMintBurnOnEth: true,
				errorsOnCalls: map[string]error{
					"BurnBalancesEth": expectedError,
				},
			}
			result := validatorTester(cfg)
			assert.Equal(t, expectedError, result.error)
			assert.True(t, result.checkRequiredBalanceOnEthCalled)
			assert.False(t, result.checkRequiredBalanceOnDrtCalled)
		})
		t.Run("on computeEthAmount, MintBalances", func(t *testing.T) {
			cfg := testConfiguration{
				direction:       batchProcessor.FromDharitrI,
				isNativeOnDrt:   true,
				isMintBurnOnEth: true,
				errorsOnCalls: map[string]error{
					"MintBalancesEth": expectedError,
				},
			}
			result := validatorTester(cfg)
			assert.Equal(t, expectedError, result.error)
			assert.True(t, result.checkRequiredBalanceOnEthCalled)
			assert.False(t, result.checkRequiredBalanceOnDrtCalled)
		})
		t.Run("on computeEthAmount, GetLastExecutedEthBatchID", func(t *testing.T) {
			cfg := testConfiguration{
				direction:       batchProcessor.FromDharitrI,
				isNativeOnDrt:   true,
				isMintBurnOnEth: true,
				errorsOnCalls: map[string]error{
					"GetLastExecutedEthBatchIDDrt": expectedError,
				},
			}
			result := validatorTester(cfg)
			assert.Equal(t, expectedError, result.error)
			assert.True(t, result.checkRequiredBalanceOnEthCalled)
			assert.False(t, result.checkRequiredBalanceOnDrtCalled)
		})
		t.Run("on computeEthAmount, GetBatch", func(t *testing.T) {
			cfg := testConfiguration{
				direction:       batchProcessor.FromDharitrI,
				isNativeOnDrt:   true,
				isMintBurnOnEth: true,
				errorsOnCalls: map[string]error{
					"GetBatchEth": expectedError,
				},
			}
			result := validatorTester(cfg)
			assert.Equal(t, expectedError, result.error)
			assert.True(t, result.checkRequiredBalanceOnEthCalled)
			assert.False(t, result.checkRequiredBalanceOnDrtCalled)
		})
		t.Run("on computeDrtAmount, TotalBalances", func(t *testing.T) {
			cfg := testConfiguration{
				direction:       batchProcessor.ToDharitrI,
				isNativeOnDrt:   true,
				isMintBurnOnEth: true,
				errorsOnCalls: map[string]error{
					"TotalBalancesDrt": expectedError,
				},
			}
			result := validatorTester(cfg)
			assert.Equal(t, expectedError, result.error)
			assert.False(t, result.checkRequiredBalanceOnEthCalled)
			assert.True(t, result.checkRequiredBalanceOnDrtCalled)
		})
		t.Run("on computeDrtAmount, BurnBalances", func(t *testing.T) {
			cfg := testConfiguration{
				direction:       batchProcessor.ToDharitrI,
				isMintBurnOnDrt: true,
				isNativeOnEth:   true,
				errorsOnCalls: map[string]error{
					"BurnBalancesDrt": expectedError,
				},
			}
			result := validatorTester(cfg)
			assert.Equal(t, expectedError, result.error)
			assert.False(t, result.checkRequiredBalanceOnEthCalled)
			assert.True(t, result.checkRequiredBalanceOnDrtCalled)
		})
		t.Run("on computeDrtAmount, MintBalances", func(t *testing.T) {
			cfg := testConfiguration{
				direction:       batchProcessor.ToDharitrI,
				isMintBurnOnDrt: true,
				isNativeOnEth:   true,
				errorsOnCalls: map[string]error{
					"MintBalancesDrt": expectedError,
				},
			}
			result := validatorTester(cfg)
			assert.Equal(t, expectedError, result.error)
			assert.False(t, result.checkRequiredBalanceOnEthCalled)
			assert.True(t, result.checkRequiredBalanceOnDrtCalled)
		})
		t.Run("on computeDrtAmount, GetLastDrtBatchID", func(t *testing.T) {
			cfg := testConfiguration{
				direction:       batchProcessor.ToDharitrI,
				isMintBurnOnDrt: true,
				isNativeOnEth:   true,
				errorsOnCalls: map[string]error{
					"GetLastDrtBatchID": expectedError,
				},
			}
			result := validatorTester(cfg)
			assert.Equal(t, expectedError, result.error)
			assert.False(t, result.checkRequiredBalanceOnEthCalled)
			assert.True(t, result.checkRequiredBalanceOnDrtCalled)
		})
		t.Run("on computeDrtAmount, GetBatch", func(t *testing.T) {
			cfg := testConfiguration{
				direction:       batchProcessor.ToDharitrI,
				isMintBurnOnDrt: true,
				isNativeOnEth:   true,
				errorsOnCalls: map[string]error{
					"GetBatchDrt": expectedError,
				},
			}
			result := validatorTester(cfg)
			assert.Equal(t, expectedError, result.error)
			assert.False(t, result.checkRequiredBalanceOnEthCalled)
			assert.True(t, result.checkRequiredBalanceOnDrtCalled)
		})
		t.Run("on computeDrtAmount, WasExecuted", func(t *testing.T) {
			cfg := testConfiguration{
				direction:       batchProcessor.ToDharitrI,
				isMintBurnOnDrt: true,
				isNativeOnEth:   true,
				errorsOnCalls: map[string]error{
					"WasExecutedEth": expectedError,
				},
			}
			result := validatorTester(cfg)
			assert.Equal(t, expectedError, result.error)
			assert.False(t, result.checkRequiredBalanceOnEthCalled)
			assert.True(t, result.checkRequiredBalanceOnDrtCalled)
		})
	})
	t.Run("invalid setup", func(t *testing.T) {
		t.Parallel()

		t.Run("on Ethereum is not native nor is mint/burn, should error", func(t *testing.T) {
			cfg := testConfiguration{
				direction:       batchProcessor.ToDharitrI,
				isMintBurnOnDrt: true,
			}
			result := validatorTester(cfg)
			assert.ErrorIs(t, result.error, ErrInvalidSetup)
			assert.Contains(t, result.error.Error(), "isNativeOnEthereum = false, isMintBurnOnEthereum = false")
			assert.False(t, result.checkRequiredBalanceOnEthCalled)
			assert.True(t, result.checkRequiredBalanceOnDrtCalled)
		})
		t.Run("on DharitrI is not native nor is mint/burn, should error", func(t *testing.T) {
			cfg := testConfiguration{
				direction:     batchProcessor.ToDharitrI,
				isNativeOnEth: true,
			}
			result := validatorTester(cfg)
			assert.ErrorIs(t, result.error, ErrInvalidSetup)
			assert.Contains(t, result.error.Error(), "isNativeOnDharitrI = false, isMintBurnOnDharitrI = false")
			assert.False(t, result.checkRequiredBalanceOnEthCalled)
			assert.True(t, result.checkRequiredBalanceOnDrtCalled)
		})
		t.Run("native on both chains, should error", func(t *testing.T) {
			cfg := testConfiguration{
				direction:     batchProcessor.ToDharitrI,
				isNativeOnEth: true,
				isNativeOnDrt: true,
			}
			result := validatorTester(cfg)
			assert.ErrorIs(t, result.error, ErrInvalidSetup)
			assert.Contains(t, result.error.Error(), "isNativeOnEthereum = true, isNativeOnDharitrI = true")
			assert.False(t, result.checkRequiredBalanceOnEthCalled)
			assert.True(t, result.checkRequiredBalanceOnDrtCalled)
		})
	})
	t.Run("bad values on mint & burn", func(t *testing.T) {
		t.Parallel()

		t.Run("on Ethereum, native", func(t *testing.T) {
			t.Parallel()

			cfg := testConfiguration{
				direction:         batchProcessor.ToDharitrI,
				isMintBurnOnEth:   true,
				isNativeOnEth:     true,
				isMintBurnOnDrt:   true,
				burnBalancesOnEth: big.NewInt(37),
				mintBalancesOnEth: big.NewInt(38),
			}
			result := validatorTester(cfg)
			assert.ErrorIs(t, result.error, ErrNegativeAmount)
			assert.Contains(t, result.error.Error(), "ethAmount: -1")
			assert.False(t, result.checkRequiredBalanceOnEthCalled)
			assert.True(t, result.checkRequiredBalanceOnDrtCalled)
		})
		t.Run("on Ethereum, non-native", func(t *testing.T) {
			t.Parallel()

			cfg := testConfiguration{
				direction:         batchProcessor.ToDharitrI,
				isMintBurnOnEth:   true,
				isNativeOnDrt:     true,
				burnBalancesOnEth: big.NewInt(38),
				mintBalancesOnEth: big.NewInt(37),
			}
			result := validatorTester(cfg)
			assert.ErrorIs(t, result.error, ErrNegativeAmount)
			assert.Contains(t, result.error.Error(), "ethAmount: -1")
			assert.False(t, result.checkRequiredBalanceOnEthCalled)
			assert.True(t, result.checkRequiredBalanceOnDrtCalled)
		})
		t.Run("on DharitrI, native", func(t *testing.T) {
			t.Parallel()

			cfg := testConfiguration{
				direction:         batchProcessor.ToDharitrI,
				isMintBurnOnEth:   true,
				isMintBurnOnDrt:   true,
				isNativeOnDrt:     true,
				burnBalancesOnDrt: big.NewInt(37),
				mintBalancesOnDrt: big.NewInt(38),
			}
			result := validatorTester(cfg)
			assert.ErrorIs(t, result.error, ErrNegativeAmount)
			assert.Contains(t, result.error.Error(), "drtAmount: -1")
			assert.False(t, result.checkRequiredBalanceOnEthCalled)
			assert.True(t, result.checkRequiredBalanceOnDrtCalled)
		})
		t.Run("on DharitrI, non-native", func(t *testing.T) {
			t.Parallel()

			cfg := testConfiguration{
				direction:         batchProcessor.ToDharitrI,
				isNativeOnEth:     true,
				isMintBurnOnDrt:   true,
				burnBalancesOnDrt: big.NewInt(38),
				mintBalancesOnDrt: big.NewInt(37),
			}
			result := validatorTester(cfg)
			assert.ErrorIs(t, result.error, ErrNegativeAmount)
			assert.Contains(t, result.error.Error(), "drtAmount: -1")
			assert.False(t, result.checkRequiredBalanceOnEthCalled)
			assert.True(t, result.checkRequiredBalanceOnDrtCalled)
		})
	})
	t.Run("scenarios", func(t *testing.T) {
		t.Parallel()

		t.Run("Eth -> DrT", func(t *testing.T) {
			t.Parallel()

			t.Run("native on DrT, mint-burn on Eth, ok values, no next pending batches", func(t *testing.T) {
				t.Parallel()

				cfg := testConfiguration{
					direction:          batchProcessor.ToDharitrI,
					isMintBurnOnEth:    true,
					isNativeOnDrt:      true,
					burnBalancesOnEth:  big.NewInt(1100),  // initial burn (1000) + burn from this transfer (100)
					mintBalancesOnEth:  big.NewInt(11000), // minted (10000) + initial burn (1000)
					totalBalancesOnDrt: big.NewInt(10000),
					amount:             amount,
					amountsOnEthPendingBatches: map[uint64][]*big.Int{
						1: {amount},
					},
					drtToken:  drtToken,
					peerToken: ethToken,
				}

				result := validatorTester(cfg)
				assert.Nil(t, result.error)
				assert.False(t, result.checkRequiredBalanceOnEthCalled)
				assert.True(t, result.checkRequiredBalanceOnDrtCalled)

				t.Run("mismatch should error", func(t *testing.T) {
					cfg.burnBalancesOnEth.Add(cfg.burnBalancesOnEth, big.NewInt(1))
					result = validatorTester(cfg)
					assert.ErrorIs(t, result.error, ErrBalanceMismatch)
				})
			})
			t.Run("native on DrT, mint-burn on Eth, ok values, with next pending batches", func(t *testing.T) {
				t.Parallel()

				cfg := testConfiguration{
					direction:          batchProcessor.ToDharitrI,
					isMintBurnOnEth:    true,
					isNativeOnDrt:      true,
					burnBalancesOnEth:  big.NewInt(1220),  // initial burn (1000) + burn from this transfer (100) + burn from next batches (120)
					mintBalancesOnEth:  big.NewInt(11000), // minted (10000) + initial burn (1000)
					totalBalancesOnDrt: big.NewInt(10000),
					amount:             amount,
					amountsOnEthPendingBatches: map[uint64][]*big.Int{
						1: {amount},
						2: {big.NewInt(30), big.NewInt(40)},
						3: {big.NewInt(50)},
					},
					drtToken:  drtToken,
					peerToken: ethToken,
				}

				result := validatorTester(cfg)
				assert.Nil(t, result.error)
				assert.False(t, result.checkRequiredBalanceOnEthCalled)
				assert.True(t, result.checkRequiredBalanceOnDrtCalled)

				t.Run("mismatch should error", func(t *testing.T) {
					cfg.burnBalancesOnEth.Add(cfg.burnBalancesOnEth, big.NewInt(1))
					result = validatorTester(cfg)
					assert.ErrorIs(t, result.error, ErrBalanceMismatch)
				})
			})
			t.Run("native on DrT but with mint-burn, mint-burn on Eth, ok values, no next pending batches", func(t *testing.T) {
				t.Parallel()

				cfg := testConfiguration{
					direction:         batchProcessor.ToDharitrI,
					isMintBurnOnEth:   true,
					isNativeOnDrt:     true,
					isMintBurnOnDrt:   true,
					burnBalancesOnEth: big.NewInt(1100),  // initial burn (1000) + burn from this transfer (100)
					mintBalancesOnEth: big.NewInt(11000), // minted (10000) + initial burn (1000)
					burnBalancesOnDrt: big.NewInt(12000),
					mintBalancesOnDrt: big.NewInt(2000), // burn - mint on Drt === mint - burn on Eth
					amount:            amount,
					amountsOnEthPendingBatches: map[uint64][]*big.Int{
						1: {amount},
					},
					drtToken:  drtToken,
					peerToken: ethToken,
				}

				result := validatorTester(cfg)
				assert.Nil(t, result.error)
				assert.False(t, result.checkRequiredBalanceOnEthCalled)
				assert.True(t, result.checkRequiredBalanceOnDrtCalled)

				t.Run("mismatch should error", func(t *testing.T) {
					cfg.burnBalancesOnEth.Add(cfg.burnBalancesOnEth, big.NewInt(1))
					result = validatorTester(cfg)
					assert.ErrorIs(t, result.error, ErrBalanceMismatch)
				})
			})
			t.Run("native on DrT but with mint-burn, mint-burn on Eth, ok values, with next pending batches", func(t *testing.T) {
				t.Parallel()

				cfg := testConfiguration{
					direction:         batchProcessor.ToDharitrI,
					isMintBurnOnEth:   true,
					isNativeOnDrt:     true,
					isMintBurnOnDrt:   true,
					burnBalancesOnEth: big.NewInt(1220),  // initial burn (1000) + burn from this transfer (100) + next batches (120)
					mintBalancesOnEth: big.NewInt(11000), // minted (10000) + initial burn (1000)
					burnBalancesOnDrt: big.NewInt(12000),
					mintBalancesOnDrt: big.NewInt(2000), // burn - mint on Drt === mint - burn on Eth
					amount:            amount,
					amountsOnEthPendingBatches: map[uint64][]*big.Int{
						1: {amount},
						2: {big.NewInt(30), big.NewInt(40)},
						3: {big.NewInt(50)},
					},
					drtToken:  drtToken,
					peerToken: ethToken,
				}

				result := validatorTester(cfg)
				assert.Nil(t, result.error)
				assert.False(t, result.checkRequiredBalanceOnEthCalled)
				assert.True(t, result.checkRequiredBalanceOnDrtCalled)

				t.Run("mismatch should error", func(t *testing.T) {
					cfg.burnBalancesOnEth.Add(cfg.burnBalancesOnEth, big.NewInt(1))
					result = validatorTester(cfg)
					assert.ErrorIs(t, result.error, ErrBalanceMismatch)
				})
			})
			t.Run("native on Eth, mint-burn on DrT, ok values, no next pending batches", func(t *testing.T) {
				t.Parallel()

				cfg := testConfiguration{
					direction:          batchProcessor.ToDharitrI,
					isMintBurnOnDrt:    true,
					isNativeOnEth:      true,
					burnBalancesOnDrt:  big.NewInt(1000),  // initial burn (1000)
					mintBalancesOnDrt:  big.NewInt(11000), // minted (10000) + initial burn (1000)
					totalBalancesOnEth: big.NewInt(10100), // initial (10000) + locked from this transfer (100)
					amount:             amount,
					amountsOnEthPendingBatches: map[uint64][]*big.Int{
						1: {amount},
					},
					drtToken:  drtToken,
					peerToken: ethToken,
				}

				result := validatorTester(cfg)
				assert.Nil(t, result.error)
				assert.False(t, result.checkRequiredBalanceOnEthCalled)
				assert.True(t, result.checkRequiredBalanceOnDrtCalled)

				t.Run("mismatch should error", func(t *testing.T) {
					cfg.burnBalancesOnDrt.Add(cfg.burnBalancesOnDrt, big.NewInt(1))
					result = validatorTester(cfg)
					assert.ErrorIs(t, result.error, ErrBalanceMismatch)
				})
			})
			t.Run("native on Eth, mint-burn on DrT, ok values, with next pending batches", func(t *testing.T) {
				t.Parallel()

				cfg := testConfiguration{
					direction:          batchProcessor.ToDharitrI,
					isMintBurnOnDrt:    true,
					isNativeOnEth:      true,
					burnBalancesOnDrt:  big.NewInt(1000),  // initial burn (1000)
					mintBalancesOnDrt:  big.NewInt(11000), // minted (10000) + initial burn (1000)
					totalBalancesOnEth: big.NewInt(10220), // initial (10000) + locked from this transfer (100) + next batches (120)
					amount:             amount,
					amountsOnEthPendingBatches: map[uint64][]*big.Int{
						1: {amount},
						2: {big.NewInt(30), big.NewInt(40)},
						3: {big.NewInt(50)},
					},
					drtToken:  drtToken,
					peerToken: ethToken,
				}

				result := validatorTester(cfg)
				assert.Nil(t, result.error)
				assert.False(t, result.checkRequiredBalanceOnEthCalled)
				assert.True(t, result.checkRequiredBalanceOnDrtCalled)

				t.Run("mismatch should error", func(t *testing.T) {
					cfg.burnBalancesOnDrt.Add(cfg.burnBalancesOnDrt, big.NewInt(1))
					result = validatorTester(cfg)
					assert.ErrorIs(t, result.error, ErrBalanceMismatch)
				})
			})
			t.Run("native on Eth but with mint-burn, mint-burn on DrT, ok values, no next pending batches", func(t *testing.T) {
				t.Parallel()

				cfg := testConfiguration{
					direction:         batchProcessor.ToDharitrI,
					isMintBurnOnDrt:   true,
					isNativeOnEth:     true,
					isMintBurnOnEth:   true,
					burnBalancesOnDrt: big.NewInt(1000),  // initial burn (1000)
					mintBalancesOnDrt: big.NewInt(11000), // minted (10000) + initial burn (1000)
					burnBalancesOnEth: big.NewInt(12100),
					mintBalancesOnEth: big.NewInt(2000), // burn - mint - transfer on Eth === mint - burn on Drt
					amount:            amount,
					amountsOnEthPendingBatches: map[uint64][]*big.Int{
						1: {amount},
					},
					drtToken:  drtToken,
					peerToken: ethToken,
				}

				result := validatorTester(cfg)
				assert.Nil(t, result.error)
				assert.False(t, result.checkRequiredBalanceOnEthCalled)
				assert.True(t, result.checkRequiredBalanceOnDrtCalled)

				t.Run("mismatch should error", func(t *testing.T) {
					cfg.burnBalancesOnDrt.Add(cfg.burnBalancesOnDrt, big.NewInt(1))
					result = validatorTester(cfg)
					assert.ErrorIs(t, result.error, ErrBalanceMismatch)
				})
			})
			t.Run("native on Eth but with mint-burn, mint-burn on DrT, ok values, with next pending batches", func(t *testing.T) {
				t.Parallel()

				cfg := testConfiguration{
					direction:         batchProcessor.ToDharitrI,
					isMintBurnOnDrt:   true,
					isNativeOnEth:     true,
					isMintBurnOnEth:   true,
					burnBalancesOnDrt: big.NewInt(1000),  // initial burn (1000)
					mintBalancesOnDrt: big.NewInt(11000), // minted (10000) + initial burn (1000)
					burnBalancesOnEth: big.NewInt(12220),
					mintBalancesOnEth: big.NewInt(2000), // burn - mint - transfer on Eth - next transfers === mint - burn on Drt
					amount:            amount,
					amountsOnEthPendingBatches: map[uint64][]*big.Int{
						1: {amount},
						2: {big.NewInt(30), big.NewInt(40)},
						3: {big.NewInt(50)},
					},
					drtToken:  drtToken,
					peerToken: ethToken,
				}

				result := validatorTester(cfg)
				assert.Nil(t, result.error)
				assert.False(t, result.checkRequiredBalanceOnEthCalled)
				assert.True(t, result.checkRequiredBalanceOnDrtCalled)

				t.Run("mismatch should error", func(t *testing.T) {
					cfg.burnBalancesOnDrt.Add(cfg.burnBalancesOnDrt, big.NewInt(1))
					result = validatorTester(cfg)
					assert.ErrorIs(t, result.error, ErrBalanceMismatch)
				})
			})
		})

		t.Run("DrT -> Eth", func(t *testing.T) {
			t.Parallel()

			t.Run("native on DrT, mint-burn on Eth, ok values, no next pending batches", func(t *testing.T) {
				t.Parallel()

				cfg := testConfiguration{
					direction:          batchProcessor.FromDharitrI,
					isMintBurnOnEth:    true,
					isNativeOnDrt:      true,
					burnBalancesOnEth:  big.NewInt(1000),  // initial burn (1000)
					mintBalancesOnEth:  big.NewInt(11000), // minted (10000) + initial burn (1000)
					totalBalancesOnDrt: big.NewInt(10100), // initial (10000) + transfer from this batch (100)
					amount:             amount,
					amountsOnDrtPendingBatches: map[uint64][]*big.Int{
						1: {amount},
					},
					drtToken:  drtToken,
					peerToken: ethToken,
				}

				result := validatorTester(cfg)
				assert.Nil(t, result.error)
				assert.True(t, result.checkRequiredBalanceOnEthCalled)
				assert.False(t, result.checkRequiredBalanceOnDrtCalled)

				t.Run("mismatch should error", func(t *testing.T) {
					cfg.burnBalancesOnEth.Add(cfg.burnBalancesOnEth, big.NewInt(1))
					result = validatorTester(cfg)
					assert.ErrorIs(t, result.error, ErrBalanceMismatch)
				})
			})
			t.Run("native on DrT, mint-burn on Eth, ok values, with next pending batches", func(t *testing.T) {
				t.Parallel()

				cfg := testConfiguration{
					direction:          batchProcessor.FromDharitrI,
					isMintBurnOnEth:    true,
					isNativeOnDrt:      true,
					burnBalancesOnEth:  big.NewInt(1000),  // initial burn (1000)
					mintBalancesOnEth:  big.NewInt(11000), // minted (10000) + initial burn (1000)
					totalBalancesOnDrt: big.NewInt(10220), // initial (10000) + transfer from this batch (100) + next batches (120)
					amount:             amount,
					amountsOnDrtPendingBatches: map[uint64][]*big.Int{
						1: {amount},
						2: {big.NewInt(30), big.NewInt(40)},
						3: {big.NewInt(50)},
					},
					drtToken:  drtToken,
					peerToken: ethToken,
				}

				result := validatorTester(cfg)
				assert.Nil(t, result.error)
				assert.True(t, result.checkRequiredBalanceOnEthCalled)
				assert.False(t, result.checkRequiredBalanceOnDrtCalled)

				t.Run("mismatch should error", func(t *testing.T) {
					cfg.burnBalancesOnEth.Add(cfg.burnBalancesOnEth, big.NewInt(1))
					result = validatorTester(cfg)
					assert.ErrorIs(t, result.error, ErrBalanceMismatch)
				})
			})
			t.Run("native on DrT but with mint-burn, mint-burn on Eth, ok values, no next pending batches", func(t *testing.T) {
				t.Parallel()

				cfg := testConfiguration{
					direction:         batchProcessor.FromDharitrI,
					isMintBurnOnEth:   true,
					isNativeOnDrt:     true,
					isMintBurnOnDrt:   true,
					burnBalancesOnEth: big.NewInt(1000),  // initial burn (1000)
					mintBalancesOnEth: big.NewInt(11000), // minted (10000) + initial burn (1000)
					burnBalancesOnDrt: big.NewInt(12100),
					mintBalancesOnDrt: big.NewInt(2000), // burn - mint - transfer on Drt === mint - burn on Eth
					amount:            amount,
					amountsOnDrtPendingBatches: map[uint64][]*big.Int{
						1: {amount},
					},
					drtToken:  drtToken,
					peerToken: ethToken,
				}

				result := validatorTester(cfg)
				assert.Nil(t, result.error)
				assert.True(t, result.checkRequiredBalanceOnEthCalled)
				assert.False(t, result.checkRequiredBalanceOnDrtCalled)

				t.Run("mismatch should error", func(t *testing.T) {
					cfg.burnBalancesOnEth.Add(cfg.burnBalancesOnEth, big.NewInt(1))
					result = validatorTester(cfg)
					assert.ErrorIs(t, result.error, ErrBalanceMismatch)
				})
			})
			t.Run("native on DrT but with mint-burn, mint-burn on Eth, ok values, with next pending batches", func(t *testing.T) {
				t.Parallel()

				cfg := testConfiguration{
					direction:         batchProcessor.FromDharitrI,
					isMintBurnOnEth:   true,
					isNativeOnDrt:     true,
					isMintBurnOnDrt:   true,
					burnBalancesOnEth: big.NewInt(1000),  // initial burn (1000)
					mintBalancesOnEth: big.NewInt(11000), // minted (10000) + initial burn (1000)
					burnBalancesOnDrt: big.NewInt(12220),
					mintBalancesOnDrt: big.NewInt(2000), // burn - mint - transfer - next batches on Drt === mint - burn on Eth
					amount:            amount,
					amountsOnDrtPendingBatches: map[uint64][]*big.Int{
						1: {amount},
						2: {big.NewInt(30), big.NewInt(40)},
						3: {big.NewInt(50)},
					},
					drtToken:  drtToken,
					peerToken: ethToken,
				}

				result := validatorTester(cfg)
				assert.Nil(t, result.error)
				assert.True(t, result.checkRequiredBalanceOnEthCalled)
				assert.False(t, result.checkRequiredBalanceOnDrtCalled)

				t.Run("mismatch should error", func(t *testing.T) {
					cfg.burnBalancesOnEth.Add(cfg.burnBalancesOnEth, big.NewInt(1))
					result = validatorTester(cfg)
					assert.ErrorIs(t, result.error, ErrBalanceMismatch)
				})
			})
			t.Run("native on Eth, mint-burn on DrT, ok values, no next pending batches", func(t *testing.T) {
				t.Parallel()

				cfg := testConfiguration{
					direction:          batchProcessor.FromDharitrI,
					isMintBurnOnDrt:    true,
					isNativeOnEth:      true,
					burnBalancesOnDrt:  big.NewInt(1100),  // initial burn (1000) + transfer from this batch (100)
					mintBalancesOnDrt:  big.NewInt(11000), // minted (10000) + initial burn (1000)
					totalBalancesOnEth: big.NewInt(10000), // initial (10000)
					amount:             amount,
					amountsOnDrtPendingBatches: map[uint64][]*big.Int{
						1: {amount},
					},
					drtToken:  drtToken,
					peerToken: ethToken,
				}

				result := validatorTester(cfg)
				assert.Nil(t, result.error)
				assert.True(t, result.checkRequiredBalanceOnEthCalled)
				assert.False(t, result.checkRequiredBalanceOnDrtCalled)

				t.Run("mismatch should error", func(t *testing.T) {
					cfg.burnBalancesOnDrt.Add(cfg.burnBalancesOnDrt, big.NewInt(1))
					result = validatorTester(cfg)
					assert.ErrorIs(t, result.error, ErrBalanceMismatch)
				})
			})
			t.Run("native on Eth, mint-burn on DrT, ok values, with next pending batches", func(t *testing.T) {
				t.Parallel()

				cfg := testConfiguration{
					direction:          batchProcessor.FromDharitrI,
					isMintBurnOnDrt:    true,
					isNativeOnEth:      true,
					burnBalancesOnDrt:  big.NewInt(1220),  // initial burn (1000) + transfer from this batch (100) + next batches (120)
					mintBalancesOnDrt:  big.NewInt(11000), // minted (10000) + initial burn (1000)
					totalBalancesOnEth: big.NewInt(10000), // initial (10000)
					amount:             amount,
					amountsOnDrtPendingBatches: map[uint64][]*big.Int{
						1: {amount},
						2: {big.NewInt(30), big.NewInt(40)},
						3: {big.NewInt(50)},
					},
					drtToken:  drtToken,
					peerToken: ethToken,
				}

				result := validatorTester(cfg)
				assert.Nil(t, result.error)
				assert.True(t, result.checkRequiredBalanceOnEthCalled)
				assert.False(t, result.checkRequiredBalanceOnDrtCalled)

				t.Run("mismatch should error", func(t *testing.T) {
					cfg.burnBalancesOnDrt.Add(cfg.burnBalancesOnDrt, big.NewInt(1))
					result = validatorTester(cfg)
					assert.ErrorIs(t, result.error, ErrBalanceMismatch)
				})
			})
			t.Run("native on Eth but with mint-burn, mint-burn on DrT, ok values, no next pending batches", func(t *testing.T) {
				t.Parallel()

				cfg := testConfiguration{
					direction:         batchProcessor.FromDharitrI,
					isMintBurnOnDrt:   true,
					isNativeOnEth:     true,
					isMintBurnOnEth:   true,
					burnBalancesOnDrt: big.NewInt(1100),  // initial burn (1000) + transfer from this batch (100)
					mintBalancesOnDrt: big.NewInt(11000), // minted (10000) + initial burn (1000)
					burnBalancesOnEth: big.NewInt(12000),
					mintBalancesOnEth: big.NewInt(2000), // burn - mint on Eth === mint - burn - transfer on Drt
					amount:            amount,
					amountsOnDrtPendingBatches: map[uint64][]*big.Int{
						1: {amount},
					},
					drtToken:  drtToken,
					peerToken: ethToken,
				}

				result := validatorTester(cfg)
				assert.Nil(t, result.error)
				assert.True(t, result.checkRequiredBalanceOnEthCalled)
				assert.False(t, result.checkRequiredBalanceOnDrtCalled)

				t.Run("mismatch should error", func(t *testing.T) {
					cfg.burnBalancesOnDrt.Add(cfg.burnBalancesOnDrt, big.NewInt(1))
					result = validatorTester(cfg)
					assert.ErrorIs(t, result.error, ErrBalanceMismatch)
				})
			})
			t.Run("native on Eth but with mint-burn, mint-burn on DrT, ok values, with next pending batches", func(t *testing.T) {
				t.Parallel()

				cfg := testConfiguration{
					direction:         batchProcessor.FromDharitrI,
					isMintBurnOnDrt:   true,
					isNativeOnEth:     true,
					isMintBurnOnEth:   true,
					burnBalancesOnDrt: big.NewInt(1220),  // initial burn (1000) + transfer from this batch (100) + transfer from next batches
					mintBalancesOnDrt: big.NewInt(11000), // minted (10000) + initial burn (1000)
					burnBalancesOnEth: big.NewInt(12000),
					mintBalancesOnEth: big.NewInt(2000), // burn - mint on Eth === mint - burn - transfer - next batches on Drt
					amount:            amount,
					amountsOnDrtPendingBatches: map[uint64][]*big.Int{
						1: {amount},
						2: {big.NewInt(30), big.NewInt(40)},
						3: {big.NewInt(50)},
					},
					drtToken:  drtToken,
					peerToken: ethToken,
				}

				result := validatorTester(cfg)
				assert.Nil(t, result.error)
				assert.True(t, result.checkRequiredBalanceOnEthCalled)
				assert.False(t, result.checkRequiredBalanceOnDrtCalled)

				t.Run("mismatch should error", func(t *testing.T) {
					cfg.burnBalancesOnDrt.Add(cfg.burnBalancesOnDrt, big.NewInt(1))
					result = validatorTester(cfg)
					assert.ErrorIs(t, result.error, ErrBalanceMismatch)
				})
			})
		})

		t.Run("DrT <-> Eth", func(t *testing.T) {
			t.Parallel()

			t.Run("from Eth: native on DrT, mint-burn on Eth, ok values, with next pending batches", func(t *testing.T) {
				t.Parallel()

				existingNativeBalanceDrt := int64(10000)
				existingBurnEth := int64(150000)
				existingMintEth := int64(160000)

				cfg := testConfiguration{
					direction:          batchProcessor.ToDharitrI,
					isMintBurnOnEth:    true,
					isNativeOnDrt:      true,
					burnBalancesOnEth:  big.NewInt(existingBurnEth + 100 + 30 + 40 + 50),
					mintBalancesOnEth:  big.NewInt(existingMintEth),
					totalBalancesOnDrt: big.NewInt(existingNativeBalanceDrt + 60 + 80 + 100 + 200),
					amount:             amount,
					pendingDrtBatchId:  1,
					amountsOnEthPendingBatches: map[uint64][]*big.Int{
						1: {amount},
						2: {big.NewInt(30), big.NewInt(40)},
						3: {big.NewInt(50)},
					},
					amountsOnDrtPendingBatches: map[uint64][]*big.Int{
						1: {amount2},
						2: {big.NewInt(60), big.NewInt(80)},
						3: {big.NewInt(100)},
					},
					drtToken:  drtToken,
					peerToken: ethToken,
				}

				result := validatorTester(cfg)
				assert.Nil(t, result.error)
				assert.False(t, result.checkRequiredBalanceOnEthCalled)
				assert.True(t, result.checkRequiredBalanceOnDrtCalled)

				t.Run("mismatch should error", func(t *testing.T) {
					copiedCfg := cfg.deepClone()
					copiedCfg.burnBalancesOnEth.Add(copiedCfg.burnBalancesOnEth, big.NewInt(1))
					result = validatorTester(copiedCfg)
					assert.ErrorIs(t, result.error, ErrBalanceMismatch)
				})
			})
			t.Run("from Eth: native on DrT but with mint-burn, mint-burn on Eth, ok values, with next pending batches", func(t *testing.T) {
				t.Parallel()

				existingBurnDrt := int64(370000) // burn > mint because the token is native
				existingMintDrt := int64(360000)
				existingBurnEth := int64(150000)
				existingMintEth := int64(160000)

				cfg := testConfiguration{
					direction:         batchProcessor.ToDharitrI,
					isMintBurnOnEth:   true,
					isNativeOnDrt:     true,
					isMintBurnOnDrt:   true,
					burnBalancesOnEth: big.NewInt(existingBurnEth + 100 + 30 + 40 + 50),
					mintBalancesOnEth: big.NewInt(existingMintEth),
					burnBalancesOnDrt: big.NewInt(existingBurnDrt + 60 + 80 + 100 + 200),
					mintBalancesOnDrt: big.NewInt(existingMintDrt),
					amount:            amount,
					pendingDrtBatchId: 1,
					amountsOnEthPendingBatches: map[uint64][]*big.Int{
						1: {amount},
						2: {big.NewInt(30), big.NewInt(40)},
						3: {big.NewInt(50)},
					},
					amountsOnDrtPendingBatches: map[uint64][]*big.Int{
						1: {amount2},
						2: {big.NewInt(60), big.NewInt(80)},
						3: {big.NewInt(100)},
					},
					drtToken:  drtToken,
					peerToken: ethToken,
				}

				result := validatorTester(cfg)
				assert.Nil(t, result.error)
				assert.False(t, result.checkRequiredBalanceOnEthCalled)
				assert.True(t, result.checkRequiredBalanceOnDrtCalled)

				t.Run("mismatch should error", func(t *testing.T) {
					copiedCfg := cfg.deepClone()
					copiedCfg.burnBalancesOnEth.Add(copiedCfg.burnBalancesOnEth, big.NewInt(1))
					result = validatorTester(copiedCfg)
					assert.ErrorIs(t, result.error, ErrBalanceMismatch)
				})
			})
			t.Run("from Eth: native on Eth, mint-burn on DrT, ok values, with next pending batches", func(t *testing.T) {
				t.Parallel()

				existingBurnDrt := int64(360000)
				existingMintDrt := int64(370000)
				existingNativeBalanceEth := int64(10000)

				cfg := testConfiguration{
					direction:          batchProcessor.ToDharitrI,
					isMintBurnOnDrt:    true,
					isNativeOnEth:      true,
					burnBalancesOnDrt:  big.NewInt(existingBurnDrt + 200 + 60 + 80 + 100),
					mintBalancesOnDrt:  big.NewInt(existingMintDrt),
					totalBalancesOnEth: big.NewInt(existingNativeBalanceEth + 100 + 30 + 40 + 50),
					amount:             amount,
					pendingDrtBatchId:  1,
					amountsOnEthPendingBatches: map[uint64][]*big.Int{
						1: {amount},
						2: {big.NewInt(30), big.NewInt(40)},
						3: {big.NewInt(50)},
					},
					amountsOnDrtPendingBatches: map[uint64][]*big.Int{
						1: {amount2},
						2: {big.NewInt(60), big.NewInt(80)},
						3: {big.NewInt(100)},
					},
					drtToken:  drtToken,
					peerToken: ethToken,
				}

				result := validatorTester(cfg)
				assert.Nil(t, result.error)
				assert.False(t, result.checkRequiredBalanceOnEthCalled)
				assert.True(t, result.checkRequiredBalanceOnDrtCalled)

				t.Run("mismatch should error", func(t *testing.T) {
					copiedCfg := cfg.deepClone()
					copiedCfg.burnBalancesOnDrt.Add(copiedCfg.burnBalancesOnDrt, big.NewInt(1))
					result = validatorTester(copiedCfg)
					assert.ErrorIs(t, result.error, ErrBalanceMismatch)
				})
			})
			t.Run("from Eth: native on Eth but with mint-burn, mint-burn on DrT, ok values, with next pending batches", func(t *testing.T) {
				t.Parallel()

				existingBurnDrt := int64(360000)
				existingMintDrt := int64(370000)
				existingBurnEth := int64(160000) // burn > mint because the token is native
				existingMintEth := int64(150000)

				cfg := testConfiguration{
					direction:         batchProcessor.ToDharitrI,
					isMintBurnOnDrt:   true,
					isNativeOnEth:     true,
					isMintBurnOnEth:   true,
					burnBalancesOnDrt: big.NewInt(existingBurnDrt + 200 + 60 + 80 + 100),
					mintBalancesOnDrt: big.NewInt(existingMintDrt),
					burnBalancesOnEth: big.NewInt(existingBurnEth + 100 + 30 + 40 + 50),
					mintBalancesOnEth: big.NewInt(existingMintEth),
					amount:            amount,
					pendingDrtBatchId: 1,
					amountsOnEthPendingBatches: map[uint64][]*big.Int{
						1: {amount},
						2: {big.NewInt(30), big.NewInt(40)},
						3: {big.NewInt(50)},
					},
					amountsOnDrtPendingBatches: map[uint64][]*big.Int{
						1: {amount2},
						2: {big.NewInt(60), big.NewInt(80)},
						3: {big.NewInt(100)},
					},
					drtToken:  drtToken,
					peerToken: ethToken,
				}

				result := validatorTester(cfg)
				assert.Nil(t, result.error)
				assert.False(t, result.checkRequiredBalanceOnEthCalled)
				assert.True(t, result.checkRequiredBalanceOnDrtCalled)

				t.Run("mismatch should error", func(t *testing.T) {
					copiedCfg := cfg.deepClone()
					copiedCfg.burnBalancesOnDrt.Add(copiedCfg.burnBalancesOnDrt, big.NewInt(1))
					result = validatorTester(copiedCfg)
					assert.ErrorIs(t, result.error, ErrBalanceMismatch)
				})
			})
			t.Run("from DrT: native on DrT, mint-burn on Eth, ok values, with next pending batches on both chains", func(t *testing.T) {
				t.Parallel()

				existingNativeBalanceDrt := int64(10000)
				existingBurnEth := int64(150000)
				existingMintEth := int64(160000)

				cfg := testConfiguration{
					direction:          batchProcessor.FromDharitrI,
					isMintBurnOnEth:    true,
					isNativeOnDrt:      true,
					burnBalancesOnEth:  big.NewInt(existingBurnEth + 200 + 60 + 80 + 100),
					mintBalancesOnEth:  big.NewInt(existingMintEth),
					totalBalancesOnDrt: big.NewInt(existingNativeBalanceDrt + 30 + 40 + 50 + 100),
					amount:             amount,
					pendingDrtBatchId:  1,
					amountsOnDrtPendingBatches: map[uint64][]*big.Int{
						1: {amount},
						2: {big.NewInt(30), big.NewInt(40)},
						3: {big.NewInt(50)},
					},
					amountsOnEthPendingBatches: map[uint64][]*big.Int{
						1: {amount2},
						2: {big.NewInt(60), big.NewInt(80)},
						3: {big.NewInt(100)},
					},
					drtToken:  drtToken,
					peerToken: ethToken,
				}

				result := validatorTester(cfg)
				assert.Nil(t, result.error)
				assert.True(t, result.checkRequiredBalanceOnEthCalled)
				assert.False(t, result.checkRequiredBalanceOnDrtCalled)

				t.Run("mismatch should error", func(t *testing.T) {
					copiedCfg := cfg.deepClone()
					copiedCfg.burnBalancesOnEth.Add(copiedCfg.burnBalancesOnEth, big.NewInt(1))
					result = validatorTester(copiedCfg)
					assert.ErrorIs(t, result.error, ErrBalanceMismatch)
				})
			})
			t.Run("from DrT: native on DrT but with mint-burn, mint-burn on Eth, ok values, with next pending batches", func(t *testing.T) {
				t.Parallel()

				existingBurnDrt := int64(370000) // burn > mint because the token is native
				existingMintDrt := int64(360000)
				existingBurnEth := int64(150000)
				existingMintEth := int64(160000)

				cfg := testConfiguration{
					direction:         batchProcessor.FromDharitrI,
					isMintBurnOnEth:   true,
					isNativeOnDrt:     true,
					isMintBurnOnDrt:   true,
					burnBalancesOnEth: big.NewInt(existingBurnEth + 200 + 60 + 80 + 100),
					mintBalancesOnEth: big.NewInt(existingMintEth),
					burnBalancesOnDrt: big.NewInt(existingBurnDrt + 30 + 40 + 50 + 100),
					mintBalancesOnDrt: big.NewInt(existingMintDrt),
					amount:            amount,
					pendingDrtBatchId: 1,
					amountsOnDrtPendingBatches: map[uint64][]*big.Int{
						1: {amount},
						2: {big.NewInt(30), big.NewInt(40)},
						3: {big.NewInt(50)},
					},
					amountsOnEthPendingBatches: map[uint64][]*big.Int{
						1: {amount2},
						2: {big.NewInt(60), big.NewInt(80)},
						3: {big.NewInt(100)},
					},
					drtToken:  drtToken,
					peerToken: ethToken,
				}

				result := validatorTester(cfg)
				assert.Nil(t, result.error)
				assert.True(t, result.checkRequiredBalanceOnEthCalled)
				assert.False(t, result.checkRequiredBalanceOnDrtCalled)

				t.Run("mismatch should error", func(t *testing.T) {
					copiedCfg := cfg.deepClone()
					copiedCfg.burnBalancesOnEth.Add(copiedCfg.burnBalancesOnEth, big.NewInt(1))
					result = validatorTester(copiedCfg)
					assert.ErrorIs(t, result.error, ErrBalanceMismatch)
				})
			})
			t.Run("from DrT: native on Eth, mint-burn on DrT, ok values, with next pending batches", func(t *testing.T) {
				t.Parallel()

				existingBurnDrt := int64(360000)
				existingMintDrt := int64(370000)
				existingNativeBalanceEth := int64(10000)

				cfg := testConfiguration{
					direction:          batchProcessor.FromDharitrI,
					isMintBurnOnDrt:    true,
					isNativeOnEth:      true,
					burnBalancesOnDrt:  big.NewInt(existingBurnDrt + 100 + 30 + 40 + 50),
					mintBalancesOnDrt:  big.NewInt(existingMintDrt),
					totalBalancesOnEth: big.NewInt(existingNativeBalanceEth + 200 + 60 + 80 + 100),
					amount:             amount,
					pendingDrtBatchId:  1,
					amountsOnDrtPendingBatches: map[uint64][]*big.Int{
						1: {amount},
						2: {big.NewInt(30), big.NewInt(40)},
						3: {big.NewInt(50)},
					},
					amountsOnEthPendingBatches: map[uint64][]*big.Int{
						1: {amount2},
						2: {big.NewInt(60), big.NewInt(80)},
						3: {big.NewInt(100)},
					},
					drtToken:  drtToken,
					peerToken: ethToken,
				}

				result := validatorTester(cfg)
				assert.Nil(t, result.error)
				assert.True(t, result.checkRequiredBalanceOnEthCalled)
				assert.False(t, result.checkRequiredBalanceOnDrtCalled)

				t.Run("mismatch should error", func(t *testing.T) {
					copiedCfg := cfg.deepClone()
					copiedCfg.burnBalancesOnDrt.Add(copiedCfg.burnBalancesOnDrt, big.NewInt(1))
					result = validatorTester(copiedCfg)
					assert.ErrorIs(t, result.error, ErrBalanceMismatch)
				})
			})
			t.Run("from DrT: native on Eth but with mint-burn, mint-burn on DrT, ok values, with next pending batches", func(t *testing.T) {
				t.Parallel()

				existingBurnDrt := int64(360000)
				existingMintDrt := int64(370000)
				existingBurnEth := int64(160000) // burn > mint because the token is native
				existingMintEth := int64(150000)

				cfg := testConfiguration{
					direction:         batchProcessor.FromDharitrI,
					isMintBurnOnDrt:   true,
					isNativeOnEth:     true,
					isMintBurnOnEth:   true,
					burnBalancesOnDrt: big.NewInt(existingBurnDrt + 100 + 30 + 40 + 50),
					mintBalancesOnDrt: big.NewInt(existingMintDrt),
					burnBalancesOnEth: big.NewInt(existingBurnEth + 200 + 60 + 80 + 100),
					mintBalancesOnEth: big.NewInt(existingMintEth),
					amount:            amount,
					pendingDrtBatchId: 1,
					amountsOnDrtPendingBatches: map[uint64][]*big.Int{
						1: {amount},
						2: {big.NewInt(30), big.NewInt(40)},
						3: {big.NewInt(50)},
					},
					amountsOnEthPendingBatches: map[uint64][]*big.Int{
						1: {amount2},
						2: {big.NewInt(60), big.NewInt(80)},
						3: {big.NewInt(100)},
					},
					drtToken:  drtToken,
					peerToken: ethToken,
				}

				result := validatorTester(cfg)
				assert.Nil(t, result.error)
				assert.True(t, result.checkRequiredBalanceOnEthCalled)
				assert.False(t, result.checkRequiredBalanceOnDrtCalled)

				t.Run("mismatch should error", func(t *testing.T) {
					copiedCfg := cfg.deepClone()
					copiedCfg.burnBalancesOnDrt.Add(copiedCfg.burnBalancesOnDrt, big.NewInt(1))
					result = validatorTester(copiedCfg)
					assert.ErrorIs(t, result.error, ErrBalanceMismatch)
				})
			})
		})
	})
}

func validatorTester(cfg testConfiguration) testResult {
	args := createMockArgsBalanceValidator()

	result := testResult{}

	lastDrtBatchID := uint64(0)
	for key := range cfg.amountsOnDrtPendingBatches {
		if key > lastDrtBatchID {
			lastDrtBatchID = key
		}
	}

	args.DharitrIClient = &bridge.DharitrIClientStub{
		CheckRequiredBalanceCalled: func(ctx context.Context, token []byte, value *big.Int) error {
			result.checkRequiredBalanceOnDrtCalled = true
			return nil
		},
		IsMintBurnTokenCalled: func(ctx context.Context, token []byte) (bool, error) {
			err := cfg.errorsOnCalls["IsMintBurnTokenDrt"]
			if err != nil {
				return false, err
			}

			return cfg.isMintBurnOnDrt, nil
		},
		IsNativeTokenCalled: func(ctx context.Context, token []byte) (bool, error) {
			err := cfg.errorsOnCalls["IsNativeTokenDrt"]
			if err != nil {
				return false, err
			}

			return cfg.isNativeOnDrt, nil
		},
		TotalBalancesCalled: func(ctx context.Context, token []byte) (*big.Int, error) {
			err := cfg.errorsOnCalls["TotalBalancesDrt"]
			if err != nil {
				return nil, err
			}

			return returnBigIntOrZeroIfNil(cfg.totalBalancesOnDrt), nil
		},
		MintBalancesCalled: func(ctx context.Context, token []byte) (*big.Int, error) {
			err := cfg.errorsOnCalls["MintBalancesDrt"]
			if err != nil {
				return nil, err
			}

			return returnBigIntOrZeroIfNil(cfg.mintBalancesOnDrt), nil
		},
		BurnBalancesCalled: func(ctx context.Context, token []byte) (*big.Int, error) {
			err := cfg.errorsOnCalls["BurnBalancesDrt"]
			if err != nil {
				return nil, err
			}

			return returnBigIntOrZeroIfNil(cfg.burnBalancesOnDrt), nil
		},
		GetPendingBatchCalled: func(ctx context.Context) (*bridgeCore.TransferBatch, error) {
			err := cfg.errorsOnCalls["GetPendingBatchDrt"]
			if err != nil {
				return nil, err
			}

			batch := &bridgeCore.TransferBatch{
				ID: cfg.pendingDrtBatchId,
			}
			applyDummyFromDrtDepositsToBatch(cfg, batch)

			return batch, nil
		},
		GetBatchCalled: func(ctx context.Context, batchID uint64) (*bridgeCore.TransferBatch, error) {
			err := cfg.errorsOnCalls["GetBatchDrt"]
			if err != nil {
				return nil, err
			}

			if batchID > getMaxDrtPendingBatchID(cfg) {
				return nil, clients.ErrNoBatchAvailable
			}
			batch := &bridgeCore.TransferBatch{
				ID: batchID,
			}
			applyDummyFromDrtDepositsToBatch(cfg, batch)

			return batch, nil
		},
		GetLastExecutedEthBatchIDCalled: func(ctx context.Context) (uint64, error) {
			err := cfg.errorsOnCalls["GetLastExecutedEthBatchIDDrt"]
			if err != nil {
				return 0, err
			}

			return cfg.lastExecutedEthBatch, nil
		},
		GetLastDrtBatchIDCalled: func(ctx context.Context) (uint64, error) {
			err := cfg.errorsOnCalls["GetLastDrtBatchID"]
			if err != nil {
				return 0, err
			}

			return lastDrtBatchID, nil
		},
	}
	args.PeerChainClient = &bridge.PeerChainClientStub{
		CheckRequiredBalanceCalled: func(ctx context.Context, tokenAddress []byte, value *big.Int) error {
			result.checkRequiredBalanceOnEthCalled = true
			return nil
		},
		MintBurnTokensCalled: func(ctx context.Context, account []byte) (bool, error) {
			err := cfg.errorsOnCalls["MintBurnTokensEth"]
			if err != nil {
				return false, err
			}

			return cfg.isMintBurnOnEth, nil
		},
		NativeTokensCalled: func(ctx context.Context, account []byte) (bool, error) {
			err := cfg.errorsOnCalls["NativeTokensEth"]
			if err != nil {
				return false, err
			}

			return cfg.isNativeOnEth, nil
		},
		TotalBalancesCalled: func(ctx context.Context, account []byte) (*big.Int, error) {
			err := cfg.errorsOnCalls["TotalBalancesEth"]
			if err != nil {
				return nil, err
			}

			return returnBigIntOrZeroIfNil(cfg.totalBalancesOnEth), nil
		},
		MintBalancesCalled: func(ctx context.Context, account []byte) (*big.Int, error) {
			err := cfg.errorsOnCalls["MintBalancesEth"]
			if err != nil {
				return nil, err
			}

			return returnBigIntOrZeroIfNil(cfg.mintBalancesOnEth), nil
		},
		BurnBalancesCalled: func(ctx context.Context, account []byte) (*big.Int, error) {
			err := cfg.errorsOnCalls["BurnBalancesEth"]
			if err != nil {
				return nil, err
			}

			return returnBigIntOrZeroIfNil(cfg.burnBalancesOnEth), nil
		},
		GetBatchCalled: func(ctx context.Context, nonce uint64) (*bridgeCore.TransferBatch, bool, error) {
			err := cfg.errorsOnCalls["GetBatchEth"]
			if err != nil {
				return nil, false, err
			}

			batch := &bridgeCore.TransferBatch{
				ID: nonce,
			}
			applyDummyFromEthDepositsToBatch(cfg, batch)

			return batch, false, nil
		},
		WasExecutedCalled: func(ctx context.Context, batchID uint64) (bool, error) {
			err := cfg.errorsOnCalls["WasExecutedEth"]
			if err != nil {
				return false, err
			}

			_, found := cfg.amountsOnDrtPendingBatches[batchID]
			return !found, nil
		},
	}

	validator, err := NewBalanceValidator(args)
	if err != nil {
		result.error = err
		return result
	}

	result.error = validator.CheckToken(context.Background(), cfg.peerToken, cfg.drtToken, cfg.amount, cfg.direction)

	return result
}

func applyDummyFromDrtDepositsToBatch(cfg testConfiguration, batch *bridgeCore.TransferBatch) {
	if cfg.amountsOnDrtPendingBatches != nil {
		values, found := cfg.amountsOnDrtPendingBatches[batch.ID]
		if found {
			depositCounter := uint64(0)

			for _, deposit := range values {
				batch.Deposits = append(batch.Deposits, &bridgeCore.DepositTransfer{
					Nonce:            depositCounter,
					Amount:           big.NewInt(0).Set(deposit),
					SourceTokenBytes: drtToken,
				})
			}
		}
	}
}

func applyDummyFromEthDepositsToBatch(cfg testConfiguration, batch *bridgeCore.TransferBatch) {
	if cfg.amountsOnEthPendingBatches != nil {
		values, found := cfg.amountsOnEthPendingBatches[batch.ID]
		if found {
			depositCounter := uint64(0)

			for _, deposit := range values {
				batch.Deposits = append(batch.Deposits, &bridgeCore.DepositTransfer{
					Nonce:            depositCounter,
					Amount:           big.NewInt(0).Set(deposit),
					SourceTokenBytes: ethToken,
				})
			}
		}
	}
}

func getMaxDrtPendingBatchID(cfg testConfiguration) uint64 {
	if cfg.amountsOnDrtPendingBatches == nil {
		return 0
	}

	maxBatchIDFound := uint64(0)
	for batchID := range cfg.amountsOnDrtPendingBatches {
		if batchID > maxBatchIDFound {
			maxBatchIDFound = batchID
		}
	}

	return maxBatchIDFound
}

func returnBigIntOrZeroIfNil(value *big.Int) *big.Int {
	if value == nil {
		return big.NewInt(0)
	}

	return big.NewInt(0).Set(value)
}
