//go:build slow

package slowTests

import (
	"math/big"

	bridgeCore "github.com/TerraDharitri/drt-go-bridge-eth/core"
	"github.com/TerraDharitri/drt-go-bridge-eth/integrationTests/relayers/slowTests/framework"
	"github.com/TerraDharitri/drt-go-bridge-eth/parsers"
	"github.com/TerraDharitri/drt-go-bridge-eth/testsCommon"
	logger "github.com/TerraDharitri/drt-go-chain-logger"
)

var (
	log = logger.GetOrCreate("integrationTests/relayers/slowTests")
)

// GenerateTestUSDCToken will generate a test USDC token
func GenerateTestUSDCToken() framework.TestTokenParams {
	// USDC is peerChainNative = true, peerChainMintBurn = false, drtNative = false, drtMintBurn = true
	return framework.TestTokenParams{
		IssueTokenParams: framework.IssueTokenParams{
			AbstractTokenIdentifier:          "USDC",
			NumOfDecimalsUniversal:           6,
			NumOfDecimalsChainSpecific:       6,
			DrtUniversalTokenTicker:          "USDC",
			DrtChainSpecificTokenTicker:      "ETHUSDC",
			DrtUniversalTokenDisplayName:     "WrappedUSDC",
			DrtChainSpecificTokenDisplayName: "EthereumWrappedUSDC",
			ValueToMintOnDrt:                 "10000000000",
			IsMintBurnOnDrT:                  true,
			IsNativeOnDrT:                    false,
			HasChainSpecificToken:            true,
			PeerChainTokenName:               "EthUSDC",
			PeerChainTokenSymbol:             "USDC",
			ValueToMintOnPeerChain:           "10000000000",
			IsMintBurnOnPeerChain:            false,
			IsNativeOnPeerChain:              true,
			PeerChainType:                    framework.ChainTypeEthereum,
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
			{
				ValueToTransferToDrt: big.NewInt(1000),
				ValueToSendFromDrT:   nil,
				DrtSCCallData:        createScCallData("callPayable", 50000000),
			},
		},
		DCDTSafeExtraBalance:          big.NewInt(100),                                        // extra is just for the fees for the 2 transfers drt->peerChain
		PeerChainTestAddrExtraBalance: big.NewInt(-5000 + 2500 - 50 - 7000 + 300 - 50 - 1000), // -(peerChain->drt) + (drt->peerChain) - fees
	}
}

// GenerateTestMEMEToken will generate a test MEME token
func GenerateTestMEMEToken() framework.TestTokenParams {
	//MEME is peerChainNative = false, peerChainMintBurn = true, drtNative = true, drtMintBurn = false
	return framework.TestTokenParams{
		IssueTokenParams: framework.IssueTokenParams{
			AbstractTokenIdentifier:          "MEME",
			NumOfDecimalsUniversal:           1,
			NumOfDecimalsChainSpecific:       1,
			DrtUniversalTokenTicker:          "MEME",
			DrtChainSpecificTokenTicker:      "ETHMEME",
			DrtUniversalTokenDisplayName:     "WrappedMEME",
			DrtChainSpecificTokenDisplayName: "EthereumWrappedMEME",
			ValueToMintOnDrt:                 "10000000000",
			IsMintBurnOnDrT:                  false,
			IsNativeOnDrT:                    true,
			HasChainSpecificToken:            true,
			PeerChainTokenName:               "EthMEME",
			PeerChainTokenSymbol:             "MEME",
			ValueToMintOnPeerChain:           "10000000000",
			IsMintBurnOnPeerChain:            true,
			IsNativeOnPeerChain:              false,
			PeerChainType:                    framework.ChainTypeEthereum,
		},
		TestOperations: []framework.TokenOperations{
			{
				ValueToTransferToDrt: big.NewInt(2400),
				ValueToSendFromDrT:   big.NewInt(4000),
			},
			{
				ValueToTransferToDrt: big.NewInt(200),
				ValueToSendFromDrT:   big.NewInt(6000),
			},
			{
				ValueToTransferToDrt: big.NewInt(1000),
				ValueToSendFromDrT:   big.NewInt(2000),
				DrtSCCallData:        createScCallData("callPayable", 50000000),
			},
		},
		DCDTSafeExtraBalance:          big.NewInt(4000 + 6000 + 2000), // everything is locked in the safe dcdt contract
		PeerChainTestAddrExtraBalance: big.NewInt(4000 - 50 + 6000 - 50 + 2000 - 50),
	}
}

// GenerateTestEUROCToken will generate a test EUROC token
func GenerateTestEUROCToken() framework.TestTokenParams {
	//EUROC is peerChainNative = true, peerChainMintBurn = true, drtNative = false, drtMintBurn = true
	return framework.TestTokenParams{
		IssueTokenParams: framework.IssueTokenParams{
			AbstractTokenIdentifier:          "EUROC",
			NumOfDecimalsUniversal:           6,
			NumOfDecimalsChainSpecific:       6,
			DrtUniversalTokenTicker:          "EUROC",
			DrtChainSpecificTokenTicker:      "EUROC",
			DrtUniversalTokenDisplayName:     "TestEUROC",
			DrtChainSpecificTokenDisplayName: "TestEUROC",
			ValueToMintOnDrt:                 "10000000000",
			IsMintBurnOnDrT:                  true,
			IsNativeOnDrT:                    false,
			HasChainSpecificToken:            false,
			PeerChainTokenName:               "EthEuroC",
			PeerChainTokenSymbol:             "EUROC",
			ValueToMintOnPeerChain:           "10000000000",
			IsMintBurnOnPeerChain:            true,
			IsNativeOnPeerChain:              true,
			PeerChainType:                    framework.ChainTypeEthereum,
		},
		TestOperations: []framework.TokenOperations{
			{
				ValueToTransferToDrt: big.NewInt(5010),
				ValueToSendFromDrT:   big.NewInt(2510),
			},
			{
				ValueToTransferToDrt: big.NewInt(7010),
				ValueToSendFromDrT:   big.NewInt(310),
			},
			{
				ValueToTransferToDrt: big.NewInt(1010),
				ValueToSendFromDrT:   nil,
				DrtSCCallData:        createScCallData("callPayable", 50000000),
			},
		},
		DCDTSafeExtraBalance:          big.NewInt(100),                                        // extra is just for the fees for the 2 transfers drt->peerChain
		PeerChainTestAddrExtraBalance: big.NewInt(-5010 + 2510 - 50 - 7010 + 310 - 50 - 1010), // -(peerChain->drt) + (drt->peerChain) - fees
	}
}

// GenerateTestMOAToken will generate a test EUROC token
func GenerateTestMOAToken() framework.TestTokenParams {
	//MOA is peerChainNative = false, peerChainMintBurn = true, drtNative = true, drtMintBurn = true
	return framework.TestTokenParams{
		IssueTokenParams: framework.IssueTokenParams{
			AbstractTokenIdentifier:          "MOA",
			NumOfDecimalsUniversal:           2,
			NumOfDecimalsChainSpecific:       2,
			DrtUniversalTokenTicker:          "MOA",
			DrtChainSpecificTokenTicker:      "MOA",
			DrtUniversalTokenDisplayName:     "TestMOA",
			DrtChainSpecificTokenDisplayName: "TestMOA",
			ValueToMintOnDrt:                 "10000000000",
			IsMintBurnOnDrT:                  true,
			IsNativeOnDrT:                    true,
			HasChainSpecificToken:            false,
			PeerChainTokenName:               "EthMex",
			PeerChainTokenSymbol:             "MOA",
			ValueToMintOnPeerChain:           "10000000000",
			IsMintBurnOnPeerChain:            true,
			IsNativeOnPeerChain:              false,
			PeerChainType:                    framework.ChainTypeEthereum,
		},
		TestOperations: []framework.TokenOperations{
			{
				ValueToTransferToDrt: big.NewInt(2410),
				ValueToSendFromDrT:   big.NewInt(4010),
			},
			{
				ValueToTransferToDrt: big.NewInt(210),
				ValueToSendFromDrT:   big.NewInt(6010),
			},
			{
				ValueToTransferToDrt: big.NewInt(1010),
				ValueToSendFromDrT:   big.NewInt(2010),
				DrtSCCallData:        createScCallData("callPayable", 50000000),
			},
		},
		DCDTSafeExtraBalance:          big.NewInt(150), // just the fees should be collected in DCDT safe
		PeerChainTestAddrExtraBalance: big.NewInt(4010 - 50 + 6010 - 50 + 2010 - 50),
	}
}

func GenerateTestSuiUSDCToken() framework.TestTokenParams {
	// SuiUSDC Sui is peerChainNative = true, peerChainMintBurn = false, drtNative = false, drtMintBurn = true
	return framework.TestTokenParams{
		IssueTokenParams: framework.IssueTokenParams{
			AbstractTokenIdentifier:          "USDC",
			NumOfDecimalsUniversal:           6,
			NumOfDecimalsChainSpecific:       6,
			DrtUniversalTokenTicker:          "USDC",
			DrtChainSpecificTokenTicker:      "SUIUSDC",
			DrtUniversalTokenDisplayName:     "WrappedUSDC",
			DrtChainSpecificTokenDisplayName: "SuiWrappedUSDC",
			ValueToMintOnDrt:                 "10000000000",
			IsMintBurnOnDrT:                  true,
			IsNativeOnDrT:                    false,
			HasChainSpecificToken:            true,
			PeerChainTokenName:               "SuiUSDC",
			PeerChainTokenSymbol:             "USDC",
			ValueToMintOnPeerChain:           "10000000000",
			IsMintBurnOnPeerChain:            false,
			IsNativeOnPeerChain:              true,
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
		DCDTSafeExtraBalance:          big.NewInt(100),                                 // extra is just for the fees for the 2 transfers drt->peerChain
		PeerChainTestAddrExtraBalance: big.NewInt(-5000 + 2500 - 50 - 7000 + 300 - 50), // -(peerChain->drt) + (drt->peerChain) - fees
	}
}

func GenerateTestWALToken() framework.TestTokenParams {
	// WAL Sui is peerChainNative = true, peerChainMintBurn = false, drtNative = false, drtMintBurn = true
	return framework.TestTokenParams{
		IssueTokenParams: framework.IssueTokenParams{
			AbstractTokenIdentifier:          "WAL",
			NumOfDecimalsUniversal:           6,
			NumOfDecimalsChainSpecific:       6,
			DrtUniversalTokenTicker:          "WAL",
			DrtChainSpecificTokenTicker:      "SUIWAL",
			DrtUniversalTokenDisplayName:     "WrappedWAL",
			DrtChainSpecificTokenDisplayName: "SuiWrappedWAL",
			ValueToMintOnDrt:                 "10000000000",
			IsMintBurnOnDrT:                  true,
			IsNativeOnDrT:                    false,
			HasChainSpecificToken:            false,
			PeerChainTokenName:               "Walrus",
			PeerChainTokenSymbol:             "WAL",
			ValueToMintOnPeerChain:           "10000000000",
			IsMintBurnOnPeerChain:            false,
			IsNativeOnPeerChain:              true,
			PeerChainType:                    framework.ChainTypeSui,
		},
		TestOperations: []framework.TokenOperations{
			{
				ValueToTransferToDrt: big.NewInt(7300),
				ValueToSendFromDrT:   big.NewInt(6150),
			},
			{
				ValueToTransferToDrt: big.NewInt(1900),
				ValueToSendFromDrT:   big.NewInt(1280),
			},
		},
		DCDTSafeExtraBalance:          big.NewInt(100),                                  // extra is just for the fees for the 2 transfers drt->peerChain
		PeerChainTestAddrExtraBalance: big.NewInt(-7300 + 6150 - 50 - 1900 + 1280 - 50), // -(peerChain->drt) + (drt->peerChain) - fees
	}
}

func GenerateTestLKXMNToken() framework.TestTokenParams {
	// LKXMN is peerChainNative = true, peerChainMintBurn = true, drtNative = false, drtMintBurn = true
	return framework.TestTokenParams{
		IssueTokenParams: framework.IssueTokenParams{
			AbstractTokenIdentifier:          "LKXMN",
			NumOfDecimalsUniversal:           6,
			NumOfDecimalsChainSpecific:       6,
			DrtUniversalTokenTicker:          "LKXMN",
			DrtChainSpecificTokenTicker:      "SUILKXMN",
			DrtUniversalTokenDisplayName:     "WrappedLKXMN",
			DrtChainSpecificTokenDisplayName: "SuiWrappedLKXMN",
			ValueToMintOnDrt:                 "10000000000",
			IsMintBurnOnDrT:                  true,
			IsNativeOnDrT:                    false,
			HasChainSpecificToken:            false,
			PeerChainTokenName:               "xMoney",
			PeerChainTokenSymbol:             "LKXMN",
			ValueToMintOnPeerChain:           "0",
			IsMintBurnOnPeerChain:            true,
			IsNativeOnPeerChain:              true,
			PeerChainType:                    framework.ChainTypeSui,
			IsLocked:                         true,
		},
		TestOperations: []framework.TokenOperations{
			{
				ValueToTransferToDrt: nil,
				ValueToSendFromDrT:   big.NewInt(1550),
			},
			{
				ValueToTransferToDrt: nil,
				ValueToSendFromDrT:   big.NewInt(4650),
			},
		},
		DCDTSafeExtraBalance:          big.NewInt(100),                   // extra is just for the fees for the 2 transfers drt->peerChain
		PeerChainTestAddrExtraBalance: big.NewInt(1550 - 50 + 4650 - 50), // -(peerChain->drt) + (drt->peerChain) - fees
	}
}

func createScCallData(function string, gasLimit uint64, args ...string) []byte {
	codec := testsCommon.TestDharitrICodec{}
	callData := parsers.CallData{
		Type:      bridgeCore.DataPresentProtocolMarker,
		Function:  function,
		GasLimit:  gasLimit,
		Arguments: args,
	}

	return codec.EncodeCallDataStrict(callData)
}
