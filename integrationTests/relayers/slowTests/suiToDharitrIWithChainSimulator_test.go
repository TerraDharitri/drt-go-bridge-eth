//go:build slow

package slowTests

import (
	"context"
	"errors"
	"math/big"
	"testing"

	"github.com/TerraDharitri/drt-go-bridge-eth/integrationTests/mock"
	"github.com/TerraDharitri/drt-go-bridge-eth/integrationTests/relayers/slowTests/framework"
	logger "github.com/TerraDharitri/drt-go-chain-logger"
	"github.com/stretchr/testify/require"
)

func TestRelayersShouldExecuteTransfersWithCoins(t *testing.T) {
	t.Skip()
	walToken := GenerateTestWALToken()
	walToken.InitialSupplyValue = "1000000000"

	suiUsdcToken := GenerateTestSuiUSDCToken()
	suiUsdcToken.InitialSupplyValue = "1000000000"

	_ = testRelayersWithChainSimulatorAndTokens(
		t,
		make(chan error),
		walToken,
		suiUsdcToken,
	)
}

func TestRelayersShouldExecuteTransfersWithTokens(t *testing.T) {
	t.Skip()
	xmnToken := GenerateTestLKXMNToken()
	xmnToken.InitialSupplyValue = "1000000000"

	_ = testRelayersWithChainSimulatorAndTokens(
		t,
		make(chan error),
		xmnToken,
	)
}

func TestRelayerShouldExecuteTransfersAndNotCatchErrorsSui(t *testing.T) {
	t.Skip()
	errorString := "ERROR"
	mockLogObserver := mock.NewMockLogObserver(errorString)
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

	walToken := GenerateTestWALToken()
	walToken.InitialSupplyValue = "1000000000"

	suiUsdcToken := GenerateTestSuiUSDCToken()
	suiUsdcToken.InitialSupplyValue = "2000000000"

	_ = testRelayersWithChainSimulatorAndTokens(
		t,
		stopChan,
		walToken,
		suiUsdcToken,
	)
}

func createSuiBadToken() framework.TestTokenParams {
	return framework.TestTokenParams{
		IssueTokenParams: framework.IssueTokenParams{
			AbstractTokenIdentifier:          "BAD",
			NumOfDecimalsUniversal:           6,
			NumOfDecimalsChainSpecific:       6,
			DrtUniversalTokenTicker:          "BAD",
			DrtChainSpecificTokenTicker:      "SUIBAD",
			DrtUniversalTokenDisplayName:     "WrappedBAD",
			DrtChainSpecificTokenDisplayName: "SuiWrappedBAD",
			ValueToMintOnDrt:                 "10000000000",
			PeerChainTokenName:               "SUITOKEN",
			PeerChainTokenSymbol:             "SUIT",
			ValueToMintOnPeerChain:           "10000000000",
			PeerChainType:                    framework.ChainTypeSui,
		},
		TestOperations: []framework.TokenOperations{
			{
				ValueToTransferToDrt: big.NewInt(5000),
				ValueToSendFromDrT:   big.NewInt(2500),
			},
			{
				ValueToTransferToDrt: big.NewInt(7000),
				ValueToSendFromDrT:   big.NewInt(300),
			},
		},
		DCDTSafeExtraBalance:          big.NewInt(0),
		PeerChainTestAddrExtraBalance: big.NewInt(0),
	}
}

func TestRelayersShouldNotExecuteTransfersSui(t *testing.T) {
	// t.Skip()
	t.Run("IsNativeOnPeerChain = true, IsMintBurnOnPeerChain = false, isNativeOnDrT = true, isMintBurnOnDrT = false", func(t *testing.T) {
		badToken := createSuiBadToken()
		badToken.IsNativeOnPeerChain = true
		badToken.IsMintBurnOnPeerChain = false
		badToken.IsNativeOnDrT = true
		badToken.IsMintBurnOnDrT = false
		badToken.HasChainSpecificToken = true

		expectedStringInLogs := "error = invalid setup isNativeOnEthereum = true, isNativeOnDharitrI = true"
		testRelayersShouldNotExecuteTransfers(t, expectedStringInLogs, badToken)
	})
	t.Run("IsNativeOnPeerChain = true, IsMintBurnOnPeerChain = false, isNativeOnDrT = true, isMintBurnOnDrT = true", func(t *testing.T) {
		badToken := createSuiBadToken()
		badToken.IsNativeOnPeerChain = true
		badToken.IsMintBurnOnPeerChain = false
		badToken.IsNativeOnDrT = true
		badToken.IsMintBurnOnDrT = true
		badToken.HasChainSpecificToken = false

		expectedStringInLogs := "error = invalid setup isNativeOnEthereum = true, isNativeOnDharitrI = true"
		testRelayersShouldNotExecuteTransfers(t, expectedStringInLogs, badToken)
	})
}
