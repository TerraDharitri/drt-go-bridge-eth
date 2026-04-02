package balanceValidator

import (
	"context"
	"math/big"

	bridgeCore "github.com/TerraDharitri/drt-go-bridge-eth/core"
)

// DharitrIClient defines the behavior of the DharitrI client able to communicate with the DharitrI chain
type DharitrIClient interface {
	GetPendingBatch(ctx context.Context) (*bridgeCore.TransferBatch, error)
	GetBatch(ctx context.Context, batchID uint64) (*bridgeCore.TransferBatch, error)
	GetLastExecutedEthBatchID(ctx context.Context) (uint64, error)
	GetLastDrtBatchID(ctx context.Context) (uint64, error)
	IsMintBurnToken(ctx context.Context, token []byte) (bool, error)
	IsNativeToken(ctx context.Context, token []byte) (bool, error)
	TotalBalances(ctx context.Context, token []byte) (*big.Int, error)
	MintBalances(ctx context.Context, token []byte) (*big.Int, error)
	BurnBalances(ctx context.Context, token []byte) (*big.Int, error)
	CheckRequiredBalance(ctx context.Context, token []byte, value *big.Int) error
	IsInterfaceNil() bool
}

// PeerChainClient defines the behavior of the peer client able to communicate with the peer chain
type PeerChainClient interface {
	GetBatch(ctx context.Context, nonce uint64) (*bridgeCore.TransferBatch, bool, error)
	TotalBalances(ctx context.Context, token []byte) (*big.Int, error)
	MintBalances(ctx context.Context, token []byte) (*big.Int, error)
	BurnBalances(ctx context.Context, token []byte) (*big.Int, error)
	MintBurnTokens(ctx context.Context, token []byte) (bool, error)
	NativeTokens(ctx context.Context, token []byte) (bool, error)
	CheckRequiredBalance(ctx context.Context, token []byte, value *big.Int) error
	WasExecuted(ctx context.Context, drtBatchID uint64) (bool, error)
	IsInterfaceNil() bool
}
