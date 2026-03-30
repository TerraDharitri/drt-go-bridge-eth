//go:build slow

package slowTests

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"testing"

	"github.com/TerraDharitri/drt-go-bridge-eth/integrationTests/mock"
	"github.com/TerraDharitri/drt-go-bridge-eth/integrationTests/relayers/slowTests/framework"
	logger "github.com/TerraDharitri/drt-go-chain-logger"
	"github.com/stretchr/testify/require"
)

func TestRelayerShouldExecuteSimultaneousSwapsAndNotCatchErrors(t *testing.T) {
	t.Skip()
	errorString := "ERROR"
	mockLogObserver := mock.NewMockLogObserver(errorString, "got invalid action ID")
	err := logger.AddLogObserver(mockLogObserver, &logger.PlainFormatter{})
	require.NoError(t, err)
	defer func() {
		require.NoError(t, logger.RemoveLogObserver(mockLogObserver))
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stopChan := make(chan error, 1000) // ensure sufficient error buffer

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-mockLogObserver.LogFoundChan():
				stopChan <- errors.New("logger should have not caught errors")
			}
		}
	}()

	usdcToken := GenerateTestUSDCToken()
	usdcToken.TestOperations = []framework.TokenOperations{
		{
			ValueToTransferToDrt: big.NewInt(5000),
			ValueToSendFromDrT:   big.NewInt(200),
			DrtSCCallData:        nil,
			DrtFaultySCCall:      false,
			DrtForceSCCall:       false,
		},
	}
	usdcToken.DCDTSafeExtraBalance = big.NewInt(50)
	usdcToken.PeerChainTestAddrExtraBalance = big.NewInt(-5000 - 5000 + 200 - 50)

	_ = testRelayersWithChainSimulatorAndTokensForSimultaneousSwaps(
		t,
		stopChan,
		usdcToken,
	)
}

func testRelayersWithChainSimulatorAndTokensForSimultaneousSwaps(tb testing.TB, manualStopChan chan error, tokens ...framework.TestTokenParams) *framework.TestSetup {
	startsFromEthFlow := &startsFromEthereumEdgecaseFlow{
		TB:     tb,
		tokens: tokens,
	}

	setupFunc := func(tb testing.TB, setup *framework.TestSetup) {
		startsFromEthFlow.setup = setup

		setup.IssueAndConfigureTokens(tokens...)
		setup.DharitriHandler.CheckForZeroBalanceOnReceivers(setup.Ctx, tokens...)
		setup.PeerChainHandler.CreateBatchOnPeerChain(setup.Ctx, setup.DharitriHandler.TestCallerAddress, startsFromEthFlow.tokens...)
	}

	processFunc := func(tb testing.TB, setup *framework.TestSetup) bool {
		if startsFromEthFlow.process() {
			setup.TestWithdrawTotalFeesOnPeerChainForTokens(startsFromEthFlow.tokens...)

			return true
		}

		switch handler := setup.PeerChainHandler.(type) {
		case *framework.EthereumHandler:
			handler.SimulatedChain.Commit()
		case *framework.SuiHandler:
			panic("sui chain simulator not yet implemented")
		default:
			panic(fmt.Sprintf("unsupported peer chain handler type: %T", handler))
		}
		setup.ChainSimulator.GenerateBlocks(setup.Ctx, 1)
		require.LessOrEqual(tb, setup.ScCallerModuleInstance.GetNumSentTransaction(), setup.GetNumScCallsOperations())

		return false
	}

	chainType := tokens[0].PeerChainType

	return testRelayersWithChainSimulator(tb,
		setupFunc,
		processFunc,
		manualStopChan,
		chainType,
	)
}
