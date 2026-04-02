package factory

import (
	"context"

	"github.com/block-vision/sui-go-sdk/models"
	"github.com/TerraDharitri/drt-go-bridge-eth/core"
	sdkCore "github.com/TerraDharitri/drt-go-sdk/core"
)

type dataGetter interface {
	GetTokenIdForErc20Address(ctx context.Context, erc20Address []byte) ([][]byte, error)
	GetERC20AddressForTokenId(ctx context.Context, tokenId []byte) ([][]byte, error)
	GetTokenIdForSuiCoin(ctx context.Context, tokenId []byte) ([][]byte, error)
	GetSuiCoinForTokenId(ctx context.Context, tokenId []byte) ([][]byte, error)
	GetAllStakedRelayers(ctx context.Context) ([][]byte, error)
	IsInterfaceNil() bool
}

type suiDataGetter interface {
	GetRelayers(ctx context.Context) ([]models.SuiAddress, error)
	IsInterfaceNil() bool
}

// BridgeComponents defines the operations for the bridge components
type BridgeComponents interface {
	Start() error
	Close() error
	DharitrIRelayerAddress() sdkCore.AddressHandler
	PeerChainRelayerAddress() string
}

// DharitrIRoleProvider defines the operations for the DharitrI role provider
type DharitrIRoleProvider interface {
	Execute(ctx context.Context) error
	IsWhitelisted(address sdkCore.AddressHandler) bool
	SortedPublicKeys() [][]byte
	IsInterfaceNil() bool
}

// PeerChainRoleProvider defines the operations for the peer chain role provider
type PeerChainRoleProvider interface {
	Execute(ctx context.Context) error
	VerifySignature(signature []byte, messageHash []byte) error
	IsInterfaceNil() bool
}

// Broadcaster defines a component able to communicate with other such instances and manage signatures and other state related data
type Broadcaster interface {
	BroadcastSignature(signature []byte, messageHash []byte)
	BroadcastJoinTopic()
	SortedPublicKeys() [][]byte
	RegisterOnTopics() error
	AddBroadcastClient(client core.BroadcastClient) error
	Close() error
	IsInterfaceNil() bool
}

// StateMachine defines a state machine component
type StateMachine interface {
	Execute(ctx context.Context) error
	IsInterfaceNil() bool
}

// PollingHandler defines a polling handler component
type PollingHandler interface {
	StartProcessingLoop() error
	IsInterfaceNil() bool
}
