package module

import (
	"context"

	"github.com/TerraDharitri/drt-go-chain-core/data/transaction"
	"github.com/TerraDharitri/drt-go-sdk/core"
)

type nonceTransactionsHandler interface {
	ApplyNonceAndGasPrice(ctx context.Context, address core.AddressHandler, tx *transaction.FrontendTransaction) error
	SendTransaction(ctx context.Context, tx *transaction.FrontendTransaction) (string, error)
	Close() error
	IsInterfaceNil() bool
}

type pollingHandler interface {
	StartProcessingLoop() error
	Close() error
	IsInterfaceNil() bool
}

type executor interface {
	Execute(ctx context.Context) error
	GetNumSentTransaction() uint32
	IsInterfaceNil() bool
}
