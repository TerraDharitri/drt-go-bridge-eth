package framework

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"path"
	"strings"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/TerraDharitri/drt-go-bridge-eth/config"
	"github.com/TerraDharitri/drt-go-bridge-eth/executors/dharitri/module"
	sdkCore "github.com/TerraDharitri/drt-go-sdk/core"
	"github.com/stretchr/testify/require"
)

// framework constants
const (
	LogStepMarker                = "#################################### %s ####################################"
	proxyCacherExpirationSeconds = 600
	proxyMaxNoncesDelta          = 7
	NumRelayers                  = 3
	NumOracles                   = 3
	quorum                       = "03"
)

// TestSetup is the struct that holds all subcomponents for the testing infrastructure
type TestSetup struct {
	testing.TB
	TokensRegistry
	*KeysStore
	Bridge                 *BridgeComponents
	PeerChainHandler       PeerChainHandler
	DharitriHandler      *DharitriHandler
	WorkingDir             string
	ChainSimulator         ChainSimulatorWrapper
	ScCallerKeys           KeysHolder
	ScCallerModuleInstance SCCallerModule
	peerChainType          ChainType

	ctxCancel                   func()
	Ctx                         context.Context
	mutBalances                 sync.RWMutex
	dcdtBalanceForSafe          map[string]*big.Int
	peerChainBalanceTestAddress map[string]*big.Int
	numScCallsInTest            uint32
}

// NewTestSetup creates a new e2e test setup
func NewTestSetup(tb testing.TB, chainType ChainType) *TestSetup {
	log.Info(fmt.Sprintf(LogStepMarker, "starting setup"))

	setup := &TestSetup{
		TB:                          tb,
		TokensRegistry:              NewTokenRegistry(tb),
		WorkingDir:                  tb.TempDir(),
		peerChainType:               chainType,
		dcdtBalanceForSafe:          make(map[string]*big.Int),
		peerChainBalanceTestAddress: make(map[string]*big.Int),
	}
	setup.KeysStore = NewKeysStore(tb, setup.WorkingDir, NumRelayers, NumOracles)

	// create a test context
	setup.Ctx, setup.ctxCancel = context.WithCancel(context.Background())

	switch chainType {
	case ChainTypeEthereum:
		setup.PeerChainHandler = NewEthereumHandler(tb, setup.Ctx, setup.KeysStore, setup.TokensRegistry, quorum)
	case ChainTypeSui:
		argsSuiChainSimulatorWrapper := ArgsSuiChainSimulatorWrapper{
			TB:    tb,
			Owner: setup.OwnerKeys,
		}
		suiChainSimulator := CreateSuiChainSimulatorWrapper(argsSuiChainSimulatorWrapper)
		setup.PeerChainHandler = NewSuiHandler(tb, setup.KeysStore, setup.TokensRegistry, suiChainSimulator, quorum)
	}

	setup.PeerChainHandler.DeployContracts(setup.Ctx)

	setup.createChainSimulatorWrapper()
	setup.DharitriHandler = NewDharitriHandler(tb, setup.Ctx, setup.KeysStore, setup.TokensRegistry, setup.ChainSimulator, quorum)
	setup.DharitriHandler.DeployAndSetContracts(setup.Ctx, chainType)

	return setup
}

func (setup *TestSetup) createChainSimulatorWrapper() {
	// create a new working directory
	tmpDir := path.Join(setup.TempDir(), "test")
	err := os.MkdirAll(tmpDir, os.ModePerm)
	require.NoError(setup, err)

	// start the chain simulator
	args := ArgChainSimulatorWrapper{
		TB:                           setup.TB,
		ProxyCacherExpirationSeconds: proxyCacherExpirationSeconds,
		ProxyMaxNoncesDelta:          proxyMaxNoncesDelta,
	}
	setup.ChainSimulator = CreateChainSimulatorWrapper(args)
	require.NoError(setup, err)
}

// StartRelayersAndScModule will start the bridge and the SC execution module
func (setup *TestSetup) StartRelayersAndScModule() {
	log.Info(fmt.Sprintf(LogStepMarker, "starting relayers & sc execution module"))

	// start relayers
	switch handler := setup.PeerChainHandler.(type) {
	case *EthereumHandler:
		setup.Bridge = NewEthereumBridgeComponents(
			setup.TB,
			setup.WorkingDir,
			setup.ChainSimulator,
			handler.EthChainWrapper,
			handler.Erc20ContractsHolder,
			handler.SimulatedChain,
			NumRelayers,
			handler.SafeAddress.Hex(),
			setup.DharitriHandler.SafeAddress,
			setup.DharitriHandler.MultisigAddress,
		)
	case *SuiHandler:
		setup.Bridge = NewSuiBridgeComponents(
			setup.TB,
			setup.WorkingDir,
			setup.ChainSimulator,
			handler.SuiChainSimulator,
			NumRelayers,
			handler.PackageID,
			setup.DharitriHandler.SafeAddress,
			setup.DharitriHandler.MultisigAddress,
			handler.BridgeObjectID,
			handler.SafeObjectID,
			handler.TreasuryId,
			handler.BridgeInitialSharedVersion,
			handler.SafeInitialSharedVersion,
			handler.TreasuryInitialSharedVersion,
		)
	default:
		panic(fmt.Sprintf("unsupported peer chain handler type: %T", handler))
	}

	setup.startScCallerModule()
}

func (setup *TestSetup) startScCallerModule() {
	cfg := config.ScCallsModuleConfig{
		ScProxyBech32Address:            setup.DharitriHandler.ScProxyAddress.Bech32(),
		ExtraGasToExecute:               60_000_000,  // 60 million: this ensures that a SC call with 0 gas limit is refunded
		MaxGasLimitToUse:                249_999_999, // max cross shard limit
		GasLimitForOutOfGasTransactions: 30_000_000,  // gas to use when a higher than max allowed is encountered
		NetworkAddress:                  setup.ChainSimulator.GetNetworkAddress(),
		ProxyMaxNoncesDelta:             5,
		ProxyFinalityCheck:              false,
		ProxyCacherExpirationSeconds:    60, // 1 minute
		ProxyRestAPIEntityType:          string(sdkCore.Proxy),
		IntervalToResendTxsInSeconds:    1,
		PrivateKeyFile:                  path.Join(setup.WorkingDir, SCCallerFilename),
		PollingIntervalInMillis:         1000, // 1 second
		Filter: config.PendingOperationsFilterConfig{
			AllowedEthAddresses: []string{"*"},
			AllowedDrtAddresses: []string{"*"},
			AllowedTokens:       []string{"*"},
		},
		TransactionChecks: config.TransactionChecksConfig{
			CheckTransactionResults:    true,
			CloseAppOnError:            false,
			ExecutionTimeoutInSeconds:  2,
			TimeInSecondsBetweenChecks: 1,
		},
	}

	var err error
	setup.ScCallerModuleInstance, err = module.NewScCallsModule(cfg, log, nil)
	require.Nil(setup, err)
	log.Info("started SC calls module", "monitoring SC proxy address", setup.DharitriHandler.ScProxyAddress)
}

// IssueAndConfigureTokens will issue and configure the provided tokens on both chains
func (setup *TestSetup) IssueAndConfigureTokens(tokens ...TestTokenParams) {
	log.Info(fmt.Sprintf(LogStepMarker, fmt.Sprintf("issuing %d tokens", len(tokens))))

	require.Greater(setup, len(tokens), 0)

	setup.PeerChainHandler.PauseContractsForTokenChanges(setup.Ctx)
	setup.DharitriHandler.PauseContractsForTokenChanges(setup.Ctx)

	for _, token := range tokens {
		setup.processNumScCallsOperations(token)
		setup.AddToken(token.IssueTokenParams)
		setup.PeerChainHandler.IssueAndWhitelistToken(setup.Ctx, token.IssueTokenParams)
		setup.DharitriHandler.IssueAndWhitelistToken(setup.Ctx, token.IssueTokenParams)

		dcdtBalanceForSafe := setup.DharitriHandler.GetDCDTChainSpecificTokenBalance(setup.Ctx, setup.DharitriHandler.SafeAddress, token.AbstractTokenIdentifier)
		var peerChainBalanceForTestAddr *big.Int
		switch setup.peerChainType {
		case ChainTypeEthereum:
			peerChainBalanceForTestAddr = setup.PeerChainHandler.GetBalance(setup.Ctx, setup.TestKeys.EthAddress.Bytes(), token.AbstractTokenIdentifier)
		case ChainTypeSui:
			peerChainBalanceForTestAddr = setup.PeerChainHandler.GetBalance(setup.Ctx, setup.TestKeys.SuiAddress, token.AbstractTokenIdentifier)
		}

		setup.mutBalances.Lock()
		setup.dcdtBalanceForSafe[token.AbstractTokenIdentifier] = dcdtBalanceForSafe
		setup.peerChainBalanceTestAddress[token.AbstractTokenIdentifier] = peerChainBalanceForTestAddr
		setup.mutBalances.Unlock()

		log.Info("recorded the DCDT balance for safe contract", "token", token.AbstractTokenIdentifier, "balance", dcdtBalanceForSafe.String())
		log.Info("recorded the peer chain balance for test address", "token", token.AbstractTokenIdentifier, "balance", peerChainBalanceForTestAddr.String())
	}

	setup.PeerChainHandler.UnPauseContractsAfterTokenChanges(setup.Ctx)
	setup.DharitriHandler.UnPauseContractsAfterTokenChanges(setup.Ctx)

	for _, token := range tokens {
		setup.DharitriHandler.SubmitAggregatorBatch(setup.Ctx, token.IssueTokenParams)
	}
}

func (setup *TestSetup) processNumScCallsOperations(token TestTokenParams) {
	for _, op := range token.TestOperations {
		if len(op.DrtSCCallData) > 0 || op.DrtForceSCCall {
			atomic.AddUint32(&setup.numScCallsInTest, 1)
		}
	}
}

// GetNumScCallsOperations returns the number of SC calls in this test setup
func (setup *TestSetup) GetNumScCallsOperations() uint32 {
	return atomic.LoadUint32(&setup.numScCallsInTest)
}

// IsTransferDoneFromPeerChain returns true if all provided tokens are bridged from peer chain towards DharitrI
func (setup *TestSetup) IsTransferDoneFromPeerChain(tokens ...TestTokenParams) bool {
	isDone := true
	for _, params := range tokens {
		isDone = isDone && setup.isTransferDoneFromPeerChainForToken(params)
	}

	return isDone
}

func (setup *TestSetup) isTransferDoneFromPeerChainForToken(params TestTokenParams) bool {
	expectedValueOnReceiver := big.NewInt(0)
	expectedValueOnContract := big.NewInt(0)
	for _, operation := range params.TestOperations {
		if operation.ValueToTransferToDrt == nil {
			continue
		}

		if len(operation.DrtSCCallData) > 0 || operation.DrtForceSCCall {
			if !operation.DrtFaultySCCall {
				expectedValueOnContract.Add(expectedValueOnContract, operation.ValueToTransferToDrt)
			}
		} else {
			expectedValueOnReceiver.Add(expectedValueOnReceiver, operation.ValueToTransferToDrt)
		}
	}

	receiverBalance := setup.DharitriHandler.GetDCDTUniversalTokenBalance(setup.Ctx, setup.TestKeys.DrtAddress, params.AbstractTokenIdentifier)
	if receiverBalance.String() != expectedValueOnReceiver.String() {
		return false
	}

	contractBalance := setup.DharitriHandler.GetDCDTUniversalTokenBalance(setup.Ctx, setup.DharitriHandler.TestCallerAddress, params.AbstractTokenIdentifier)
	return contractBalance.String() == expectedValueOnContract.String()
}

// IsTransferDoneFromPeerChainWithRefund returns true if all provided tokens are bridged from peer chain towards DharitrI including refunds
func (setup *TestSetup) IsTransferDoneFromPeerChainWithRefund(tokens ...TestTokenParams) bool {
	isDone := true
	for _, params := range tokens {
		isDone = isDone && setup.isTransferDoneFromPeerChainWithRefundForToken(params)
	}

	return isDone
}

func (setup *TestSetup) isTransferDoneFromPeerChainWithRefundForToken(params TestTokenParams) bool {
	expectedValueOnReceiver := big.NewInt(0)
	for _, operation := range params.TestOperations {
		valueToTransferToDrt := big.NewInt(0)
		if operation.ValueToTransferToDrt != nil {
			valueToTransferToDrt.Set(operation.ValueToTransferToDrt)
		}

		valueToSendFromDrT := big.NewInt(0)
		if operation.ValueToSendFromDrT != nil {
			valueToSendFromDrT.Set(operation.ValueToSendFromDrT)
			// we subtract the fee also
			expectedValueOnReceiver.Sub(expectedValueOnReceiver, feeInt)
		}

		expectedValueOnReceiver.Add(expectedValueOnReceiver, big.NewInt(0).Sub(valueToSendFromDrT, valueToTransferToDrt))
		if len(operation.DrtSCCallData) > 0 || operation.DrtForceSCCall {
			if operation.DrtFaultySCCall {
				// the balance should be bridged back to the receiver on peer chain - fee
				expectedValueOnReceiver.Add(expectedValueOnReceiver, valueToTransferToDrt)
				expectedValueOnReceiver.Sub(expectedValueOnReceiver, feeInt)
			}
		}
	}

	var receiverBalance *big.Int
	switch setup.peerChainType {
	case ChainTypeEthereum:
		receiverBalance = setup.PeerChainHandler.GetBalance(setup.Ctx, setup.TestKeys.EthAddress.Bytes(), params.AbstractTokenIdentifier)
	case ChainTypeSui:
		receiverBalance = setup.PeerChainHandler.GetBalance(setup.Ctx, setup.TestKeys.SuiAddress, params.AbstractTokenIdentifier)
	}

	return receiverBalance.String() == expectedValueOnReceiver.String()
}

// IsTransferDoneFromDharitrI returns true if all provided tokens are bridged from DharitrI towards peer chain
func (setup *TestSetup) IsTransferDoneFromDharitrI(tokens ...TestTokenParams) bool {
	isDone := true
	for _, params := range tokens {
		isDone = isDone && setup.isTransferDoneFromDharitrIForToken(params)
	}

	return isDone
}

func (setup *TestSetup) isTransferDoneFromDharitrIForToken(params TestTokenParams) bool {
	setup.mutBalances.Lock()
	initialBalanceForSafe := setup.dcdtBalanceForSafe[params.AbstractTokenIdentifier]
	expectedReceiver := big.NewInt(0).Set(setup.peerChainBalanceTestAddress[params.AbstractTokenIdentifier])
	expectedReceiver.Add(expectedReceiver, params.PeerChainTestAddrExtraBalance)
	setup.mutBalances.Unlock()

	var peerChainTestBalance *big.Int
	switch setup.peerChainType {
	case ChainTypeEthereum:
		peerChainTestBalance = setup.PeerChainHandler.GetBalance(setup.Ctx, setup.TestKeys.EthAddress.Bytes(), params.AbstractTokenIdentifier)
	case ChainTypeSui:
		peerChainTestBalance = setup.PeerChainHandler.GetBalance(setup.Ctx, setup.TestKeys.SuiAddress, params.AbstractTokenIdentifier)
	}
	isTransferDoneFromDharitrI := peerChainTestBalance.String() == expectedReceiver.String()

	expectedDcdtSafe := big.NewInt(0).Add(initialBalanceForSafe, params.DCDTSafeExtraBalance)
	balanceForSafe := setup.DharitriHandler.GetDCDTChainSpecificTokenBalance(setup.Ctx, setup.DharitriHandler.SafeAddress, params.AbstractTokenIdentifier)
	isSafeContractOnCorrectBalance := expectedDcdtSafe.String() == balanceForSafe.String()

	return isTransferDoneFromDharitrI && isSafeContractOnCorrectBalance
}

// CreateBatchOnDharitrI will create deposits that will be gathered in a batch on DharitrI
func (setup *TestSetup) CreateBatchOnDharitrI(tokensParams ...TestTokenParams) {
	for _, params := range tokensParams {
		setup.createBatchOnDharitrIForToken(params)
	}
}

func (setup *TestSetup) createBatchOnDharitrIForToken(params TestTokenParams) {
	token := setup.GetTokenData(params.AbstractTokenIdentifier)
	require.NotNil(setup, token)

	setup.transferTokensToTestKey(params)
	valueToMintOnPeerChain := setup.sendFromDharitriToPeerChainForToken(params)
	setup.PeerChainHandler.Mint(setup.Ctx, params, valueToMintOnPeerChain)
}

func (setup *TestSetup) transferTokensToTestKey(params TestTokenParams) {
	depositValue := big.NewInt(0)
	for _, operation := range params.TestOperations {
		if operation.ValueToSendFromDrT == nil {
			continue
		}

		depositValue.Add(depositValue, operation.ValueToSendFromDrT)
	}

	setup.DharitriHandler.TransferToken(
		setup.Ctx,
		setup.OwnerKeys,
		setup.TestKeys,
		depositValue,
		params,
	)
}

// SendFromDharitriToPeerChain will create the deposits that will be gathered in a batch on DharitrI (without mint on peer chain)
func (setup *TestSetup) SendFromDharitriToPeerChain(tokensParams ...TestTokenParams) {
	for _, params := range tokensParams {
		if params.IsLocked {
			setup.transferTokensToTestKey(params)
		}
		_ = setup.sendFromDharitriToPeerChainForToken(params)
	}
}

func (setup *TestSetup) sendFromDharitriToPeerChainForToken(params TestTokenParams) *big.Int {
	token := setup.GetTokenData(params.AbstractTokenIdentifier)
	require.NotNil(setup, token)

	depositValue := big.NewInt(0)
	for _, operation := range params.TestOperations {
		if operation.ValueToSendFromDrT == nil {
			continue
		}

		depositValue.Add(depositValue, operation.ValueToSendFromDrT)

		var receiverAddress []byte
		switch setup.peerChainType {
		case ChainTypeEthereum:
			receiverAddress = setup.TestKeys.EthAddress.Bytes()
		case ChainTypeSui:
			receiverAddress, _ = hex.DecodeString(strings.TrimPrefix(string(setup.TestKeys.SuiAddress), "0x"))
		}
		setup.DharitriHandler.SendDepositTransactionFromDharitri(setup.Ctx, token, params, operation.ValueToSendFromDrT, receiverAddress)
	}

	return depositValue
}

// TestWithdrawTotalFeesOnPeerChainForTokens will test the withdrawal functionality for the provided test tokens
func (setup *TestSetup) TestWithdrawTotalFeesOnPeerChainForTokens(tokensParams ...TestTokenParams) {
	for _, param := range tokensParams {
		token := setup.GetTokenData(param.AbstractTokenIdentifier)

		expectedAccumulated := big.NewInt(0)
		for _, operation := range param.TestOperations {
			if operation.ValueToSendFromDrT == nil {
				continue
			}
			if operation.ValueToSendFromDrT.Cmp(zeroValueBigInt) == 0 {
				continue
			}

			expectedAccumulated.Add(expectedAccumulated, feeInt)
		}

		setup.DharitriHandler.TestWithdrawFees(setup.Ctx, token.DrtChainSpecificToken, zeroValueBigInt, expectedAccumulated)
	}
}

// Close will close the test subcomponents
func (setup *TestSetup) Close() {
	log.Info(fmt.Sprintf(LogStepMarker, "closing relayers & sc execution module"))

	setup.Bridge.CloseRelayers()
	require.NoError(setup, setup.PeerChainHandler.Close())

	setup.ctxCancel()
	_ = setup.ScCallerModuleInstance.Close()
}
