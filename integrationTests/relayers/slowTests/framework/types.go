package framework

import (
	"math/big"
)

// IssueTokenParams the parameters when issuing a new token
type IssueTokenParams struct {
	InitialSupplyParams
	AbstractTokenIdentifier string

	// DharitrI
	NumOfDecimalsUniversal           int
	NumOfDecimalsChainSpecific       byte
	DrtUniversalTokenTicker          string
	DrtChainSpecificTokenTicker      string
	DrtUniversalTokenDisplayName     string
	DrtChainSpecificTokenDisplayName string
	ValueToMintOnDrt                 string
	IsMintBurnOnDrT                  bool
	IsNativeOnDrT                    bool
	HasChainSpecificToken            bool

	// Peer chain
	PeerChainTokenName     string
	PeerChainTokenSymbol   string
	ValueToMintOnPeerChain string
	IsMintBurnOnPeerChain  bool
	IsNativeOnPeerChain    bool
	PeerChainType          ChainType
	IsLocked               bool
}

// InitialSupplyParams represents the initial supply parameters
type InitialSupplyParams struct {
	InitialSupplyValue string
}

// TokenOperations defines a token operation in a test. Usually this can define one or to deposits in a batch
type TokenOperations struct {
	ValueToTransferToDrt *big.Int
	ValueToSendFromDrT   *big.Int
	DrtSCCallData        []byte
	DrtFaultySCCall      bool
	DrtForceSCCall       bool
}

// TestTokenParams defines a token collection of operations in one or 2 batches
type TestTokenParams struct {
	IssueTokenParams
	TestOperations                []TokenOperations
	DCDTSafeExtraBalance          *big.Int
	PeerChainTestAddrExtraBalance *big.Int
}

// TokenData represents a test token data
type TokenData struct {
	AbstractTokenIdentifier string

	DrtUniversalTokenTicker     string
	DrtChainSpecificTokenTicker string
	PeerChainTokenName          string
	PeerChainTokenSymbol        string

	DrtUniversalToken     string
	DrtChainSpecificToken string
	PeerChainTokenAddress []byte
	PeerChainTokenInfo    interface{}
}

type EthTokenInfo struct {
	Contract ERC20Contract
}

type SuiTokenInfo struct {
	CoinPackageId  string
	TreasuryId     string
	CoinMetadataId string
	IsLocked       bool
}

type ChainType string

const (
	ChainTypeEthereum ChainType = "ethereum"
	ChainTypeSui      ChainType = "sui"
)
