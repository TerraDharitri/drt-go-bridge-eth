package factory

import (
	"errors"
	"fmt"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/TerraDharitri/drt-go-bridge-eth/clients/chain"
	"github.com/TerraDharitri/drt-go-bridge-eth/config"
	"github.com/TerraDharitri/drt-go-bridge-eth/core"
	"github.com/TerraDharitri/drt-go-bridge-eth/status"
	"github.com/TerraDharitri/drt-go-bridge-eth/testsCommon"
	bridgeTests "github.com/TerraDharitri/drt-go-bridge-eth/testsCommon/bridge"
	p2pMocks "github.com/TerraDharitri/drt-go-bridge-eth/testsCommon/p2p"
	"github.com/TerraDharitri/drt-go-chain-core/core/check"
	"github.com/TerraDharitri/drt-go-chain/testscommon/statusHandler"
	logger "github.com/TerraDharitri/drt-go-chain-logger"
	"github.com/TerraDharitri/drt-go-sdk/blockchain"
	sdkCore "github.com/TerraDharitri/drt-go-sdk/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createMockEthDharitrIBridgeArgs() ArgsEthereumToDharitrIBridge {
	stateMachineConfig := config.ConfigStateMachine{
		StepDurationInMillis:       1000,
		IntervalForLeaderInSeconds: 60,
	}

	cfg := config.Config{
		Eth: config.EthereumConfig{
			Chain:                        chain.Ethereum,
			NetworkAddress:               "http://127.0.0.1:8545",
			SafeContractAddress:          "5DdDe022a65F8063eE9adaC54F359CBF46166068",
			PrivateKeyFile:               "testdata/grace.sk",
			IntervalToResendTxsInSeconds: 0,
			GasLimitBase:                 200000,
			GasLimitForEach:              30000,
			GasStation: config.GasStationConfig{
				Enabled:                    true,
				URL:                        "",
				PollingIntervalInSeconds:   1,
				RequestRetryDelayInSeconds: 1,
				MaxFetchRetries:            3,
				RequestTimeInSeconds:       1,
				MaximumAllowedGasPrice:     100,
				GasPriceSelector:           "FastGasPrice",
				GasPriceMultiplier:         1,
			},
			MaxRetriesOnQuorumReached:          1,
			IntervalToWaitForTransferInSeconds: 1,
			ClientAvailabilityAllowDelta:       10,
		},
		DharitrI: config.DharitrIConfig{
			PrivateKeyFile:                  "testdata/grace.pem",
			IntervalToResendTxsInSeconds:    60,
			NetworkAddress:                  "http://127.0.0.1:8079",
			MultisigContractAddress:         "drt1qqqqqqqqqqqqqpgqgftcwj09u0nhmskrw7xxqcqh8qmzwyexd8ssr4um9c",
			SafeContractAddress:             "drt1qqqqqqqqqqqqqpgqgftcwj09u0nhmskrw7xxqcqh8qmzwyexd8ssr4um9c",
			GasMap:                          testsCommon.CreateTestDharitrIGasMap(),
			MaxRetriesOnQuorumReached:       1,
			MaxRetriesOnWasTransferProposed: 1,
			ClientAvailabilityAllowDelta:    10,
			Proxy: config.ProxyConfig{
				CacherExpirationSeconds: 600,
				RestAPIEntityType:       "observer",
				MaxNoncesDelta:          10,
				FinalityCheck:           true,
			},
		},
		Relayer: config.ConfigRelayer{
			RoleProvider: config.RoleProviderConfig{
				PollingIntervalInMillis: 1000,
			},
		},
		StateMachine: map[string]config.ConfigStateMachine{
			"EthereumToDharitrI": stateMachineConfig,
			"DharitrIToEthereum": stateMachineConfig,
		},
	}
	configs := config.Configs{
		GeneralConfig:   cfg,
		ApiRoutesConfig: config.ApiRoutesConfig{},
		FlagsConfig: config.ContextFlagsConfig{
			RestApiInterface: core.WebServerOffString,
		},
	}

	argsProxy := blockchain.ArgsProxy{
		ProxyURL:            cfg.DharitrI.NetworkAddress,
		CacheExpirationTime: time.Minute,
		EntityType:          sdkCore.ObserverNode,
	}
	proxy, _ := blockchain.NewProxy(argsProxy)
	return ArgsEthereumToDharitrIBridge{
		Configs:                       configs,
		Messenger:                     &p2pMocks.MessengerStub{},
		StatusStorer:                  testsCommon.NewStorerMock(),
		Proxy:                         proxy,
		DharitrIClientStatusHandler: &testsCommon.StatusHandlerStub{},
		Erc20ContractsHolder:          &bridgeTests.ERC20ContractsHolderStub{},
		ClientWrapper:                 &bridgeTests.EthereumClientWrapperStub{},
		TimeForBootstrap:              minTimeForBootstrap,
		TimeBeforeRepeatJoin:          minTimeBeforeRepeatJoin,
		MetricsHolder:                 status.NewMetricsHolder(),
		AppStatusHandler:              &statusHandler.AppStatusHandlerStub{},
	}
}

func TestNewEthDrtBridgeComponents(t *testing.T) {
	t.Parallel()

	t.Run("nil Proxy", func(t *testing.T) {
		t.Parallel()
		args := createMockEthDharitrIBridgeArgs()
		args.Proxy = nil

		components, err := NewEthDrtBridgeComponents(args)
		assert.True(t, errors.Is(err, errNilProxy))
		assert.Nil(t, components)
	})
	t.Run("nil Messenger", func(t *testing.T) {
		t.Parallel()
		args := createMockEthDharitrIBridgeArgs()
		args.Messenger = nil

		components, err := NewEthDrtBridgeComponents(args)
		assert.Equal(t, errNilMessenger, err)
		assert.Nil(t, components)
	})
	t.Run("nil ClientWrapper", func(t *testing.T) {
		t.Parallel()
		args := createMockEthDharitrIBridgeArgs()
		args.ClientWrapper = nil

		components, err := NewEthDrtBridgeComponents(args)
		assert.Equal(t, errNilEthClient, err)
		assert.Nil(t, components)
	})
	t.Run("nil StatusStorer", func(t *testing.T) {
		t.Parallel()
		args := createMockEthDharitrIBridgeArgs()
		args.StatusStorer = nil

		components, err := NewEthDrtBridgeComponents(args)
		assert.Equal(t, errNilStatusStorer, err)
		assert.Nil(t, components)
	})
	t.Run("nil Erc20ContractsHolder", func(t *testing.T) {
		t.Parallel()
		args := createMockEthDharitrIBridgeArgs()
		args.Erc20ContractsHolder = nil

		components, err := NewEthDrtBridgeComponents(args)
		assert.Equal(t, errNilErc20ContractsHolder, err)
		assert.Nil(t, components)
	})
	t.Run("err on createDharitrIKeysAndAddresses, empty pk file", func(t *testing.T) {
		t.Parallel()
		args := createMockEthDharitrIBridgeArgs()
		args.Configs.GeneralConfig.DharitrI.PrivateKeyFile = ""

		components, err := NewEthDrtBridgeComponents(args)
		assert.NotNil(t, err)
		assert.Nil(t, components)
	})
	t.Run("err on createDharitrIKeysAndAddresses, empty multisig address", func(t *testing.T) {
		t.Parallel()
		args := createMockEthDharitrIBridgeArgs()
		args.Configs.GeneralConfig.DharitrI.MultisigContractAddress = ""

		components, err := NewEthDrtBridgeComponents(args)
		assert.NotNil(t, err)
		assert.Nil(t, components)
	})
	t.Run("err on createDharitrIClient", func(t *testing.T) {
		t.Parallel()
		args := createMockEthDharitrIBridgeArgs()
		args.Configs.GeneralConfig.DharitrI.GasMap = config.DharitrIGasMapConfig{}

		components, err := NewEthDrtBridgeComponents(args)
		assert.NotNil(t, err)
		assert.Nil(t, components)
	})
	t.Run("err on createDharitrIRoleProvider", func(t *testing.T) {
		t.Parallel()
		args := createMockEthDharitrIBridgeArgs()
		args.Configs.GeneralConfig.Relayer.RoleProvider.PollingIntervalInMillis = 0

		components, err := NewEthDrtBridgeComponents(args)
		assert.NotNil(t, err)
		assert.Nil(t, components)
	})
	t.Run("err on createEthereumClient, empty eth config", func(t *testing.T) {
		t.Parallel()
		args := createMockEthDharitrIBridgeArgs()
		args.Configs.GeneralConfig.Eth = config.EthereumConfig{}

		components, err := NewEthDrtBridgeComponents(args)
		assert.NotNil(t, err)
		assert.Nil(t, components)
	})
	t.Run("err on createEthereumClient, invalid gas price selector", func(t *testing.T) {
		t.Parallel()
		args := createMockEthDharitrIBridgeArgs()
		args.Configs.GeneralConfig.Eth.GasStation.GasPriceSelector = core.WebServerOffString

		components, err := NewEthDrtBridgeComponents(args)
		assert.NotNil(t, err)
		assert.Nil(t, components)
	})
	t.Run("err missing state machine config", func(t *testing.T) {
		t.Parallel()
		args := createMockEthDharitrIBridgeArgs()
		args.Configs.GeneralConfig.StateMachine = make(map[string]config.ConfigStateMachine)

		components, err := NewEthDrtBridgeComponents(args)
		assert.True(t, errors.Is(err, errMissingConfig))
		assert.True(t, strings.Contains(err.Error(), args.Configs.GeneralConfig.Eth.Chain.PeerChainToDharitrIName()))
		assert.Nil(t, components)
	})
	t.Run("invalid time for bootstrap", func(t *testing.T) {
		t.Parallel()
		args := createMockEthDharitrIBridgeArgs()
		args.TimeForBootstrap = minTimeForBootstrap - 1

		components, err := NewEthDrtBridgeComponents(args)
		assert.True(t, errors.Is(err, errInvalidValue))
		assert.True(t, strings.Contains(err.Error(), "for TimeForBootstrap"))
		assert.Nil(t, components)
	})
	t.Run("invalid time before retry", func(t *testing.T) {
		t.Parallel()
		args := createMockEthDharitrIBridgeArgs()
		args.TimeBeforeRepeatJoin = minTimeBeforeRepeatJoin - 1

		components, err := NewEthDrtBridgeComponents(args)
		assert.True(t, errors.Is(err, errInvalidValue))
		assert.True(t, strings.Contains(err.Error(), "for TimeBeforeRepeatJoin"))
		assert.Nil(t, components)
	})
	t.Run("nil MetricsHolder", func(t *testing.T) {
		t.Parallel()
		args := createMockEthDharitrIBridgeArgs()
		args.MetricsHolder = nil

		components, err := NewEthDrtBridgeComponents(args)
		assert.Equal(t, errNilMetricsHolder, err)
		assert.Nil(t, components)
	})
	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		args := createMockEthDharitrIBridgeArgs()

		components, err := NewEthDrtBridgeComponents(args)
		require.Nil(t, err)
		require.NotNil(t, components)
		require.Equal(t, 7, len(components.closableHandlers))
		require.False(t, check.IfNil(components.toDharitrIStatusHandler))
		require.False(t, check.IfNil(components.fromDharitrIStatusHandler))
	})
}

func TestEthDharitrIBridgeComponents_StartAndCloseShouldWork(t *testing.T) {
	t.Parallel()

	args := createMockEthDharitrIBridgeArgs()
	components, err := NewEthDrtBridgeComponents(args)
	assert.Nil(t, err)

	err = components.Start()
	assert.Nil(t, err)
	assert.Equal(t, 7, len(components.closableHandlers))

	time.Sleep(time.Second * 2) // allow go routines to start

	err = components.Close()
	assert.Nil(t, err)
}

func TestEthDharitrIBridgeComponents_Start(t *testing.T) {
	t.Parallel()

	t.Run("messenger errors on bootstrap", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("expected error")
		args := createMockEthDharitrIBridgeArgs()
		args.Messenger = &p2pMocks.MessengerStub{
			BootstrapCalled: func() error {
				return expectedErr
			},
		}
		components, _ := NewEthDrtBridgeComponents(args)

		err := components.Start()
		assert.Equal(t, expectedErr, err)
	})
	t.Run("broadcaster errors on RegisterOnTopics", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("expected error")
		args := createMockEthDharitrIBridgeArgs()
		components, _ := NewEthDrtBridgeComponents(args)
		components.broadcaster = &testsCommon.BroadcasterStub{
			RegisterOnTopicsCalled: func() error {
				return expectedErr
			},
		}

		err := components.Start()
		assert.Equal(t, expectedErr, err)
	})
}

func TestEthDharitrIBridgeComponents_Close(t *testing.T) {
	t.Parallel()

	t.Run("nil closable should not panic", func(t *testing.T) {
		t.Parallel()

		defer func() {
			r := recover()
			if r != nil {
				assert.Fail(t, fmt.Sprintf("should have not failed %v", r))
			}
		}()

		components := &ethDrtBridgeComponents{
			baseBridgeComponents: &baseBridgeComponents{
				baseLogger: logger.GetOrCreate("test"),
			},
		}
		components.addClosableComponent(nil)

		err := components.Close()
		assert.Nil(t, err)
	})
	t.Run("one component errors, should return error", func(t *testing.T) {
		t.Parallel()

		components := &ethDrtBridgeComponents{
			baseBridgeComponents: &baseBridgeComponents{
				baseLogger: logger.GetOrCreate("test"),
			},
		}

		expectedErr := errors.New("expected error")

		numCalls := 0
		components.addClosableComponent(&testsCommon.CloserStub{
			CloseCalled: func() error {
				numCalls++
				return nil
			},
		})
		components.addClosableComponent(&testsCommon.CloserStub{
			CloseCalled: func() error {
				numCalls++
				return expectedErr
			},
		})
		components.addClosableComponent(&testsCommon.CloserStub{
			CloseCalled: func() error {
				numCalls++
				return nil
			},
		})

		err := components.Close()
		assert.Equal(t, expectedErr, err)
		assert.Equal(t, 3, numCalls)
	})
}

func TestEthDharitrIBridgeComponents_startBroadcastJoinRetriesLoop(t *testing.T) {
	t.Parallel()

	t.Run("close before minTimeBeforeRepeatJoin", func(t *testing.T) {
		t.Parallel()

		numberOfCalls := uint32(0)
		args := createMockEthDharitrIBridgeArgs()
		components, _ := NewEthDrtBridgeComponents(args)

		components.broadcaster = &testsCommon.BroadcasterStub{
			BroadcastJoinTopicCalled: func() {
				atomic.AddUint32(&numberOfCalls, 1)
			},
		}

		err := components.Start()
		assert.Nil(t, err)
		time.Sleep(time.Second * 3)

		err = components.Close()
		assert.Nil(t, err)
		assert.Equal(t, uint32(1), atomic.LoadUint32(&numberOfCalls)) // one call expected from Start
	})
	t.Run("broadcast should be called again", func(t *testing.T) {
		t.Parallel()

		numberOfCalls := uint32(0)
		args := createMockEthDharitrIBridgeArgs()
		components, _ := NewEthDrtBridgeComponents(args)
		components.timeBeforeRepeatJoin = time.Second * 3
		components.broadcaster = &testsCommon.BroadcasterStub{
			BroadcastJoinTopicCalled: func() {
				atomic.AddUint32(&numberOfCalls, 1)
			},
		}

		err := components.Start()
		assert.Nil(t, err)
		time.Sleep(time.Second * 7)

		err = components.Close()
		assert.Nil(t, err)
		assert.Equal(t, uint32(3), atomic.LoadUint32(&numberOfCalls)) // 3 calls expected: Start + 2 times from loop
	})
}

func TestEthDharitrIBridgeComponents_EthRelayerAddresses(t *testing.T) {
	t.Parallel()

	args := createMockEthDharitrIBridgeArgs()
	components, _ := NewEthDrtBridgeComponents(args)

	assert.Equal(t, "0x3FE464Ac5aa562F7948322F92020F2b668D543d8", components.PeerChainRelayerAddress())
}
