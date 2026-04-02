package dharitri

import (
	"context"

	"github.com/TerraDharitri/drt-go-bridge-eth/parsers"
	"github.com/TerraDharitri/drt-go-chain-core/data/api"
	"github.com/TerraDharitri/drt-go-chain-core/data/transaction"
	"github.com/TerraDharitri/drt-go-sdk/core"
	"github.com/TerraDharitri/drt-go-sdk/data"
)

// Proxy defines the behavior of a proxy able to serve DharitrI blockchain requests
type Proxy interface {
	GetNetworkConfig(ctx context.Context) (*data.NetworkConfig, error)
	SendTransaction(ctx context.Context, tx *transaction.FrontendTransaction) (string, error)
	SendTransactions(ctx context.Context, txs []*transaction.FrontendTransaction) ([]string, error)
	ExecuteVMQuery(ctx context.Context, vmRequest *data.VmValueRequest) (*data.VmValuesResponseData, error)
	GetAccount(ctx context.Context, address core.AddressHandler) (*data.Account, error)
	GetNetworkStatus(ctx context.Context, shardID uint32) (*data.NetworkStatus, error)
	GetShardOfAddress(ctx context.Context, bech32Address string) (uint32, error)
	GetDCDTTokenData(ctx context.Context, address core.AddressHandler, tokenIdentifier string, queryOptions api.AccountQueryOptions) (*data.DCDTFungibleTokenData, error)
	GetTransactionInfoWithResults(ctx context.Context, hash string) (*data.TransactionInfo, error)
	ProcessTransactionStatus(ctx context.Context, hexTxHash string) (transaction.TxStatus, error)
	IsInterfaceNil() bool
}

// NonceTransactionsHandler represents the interface able to handle the current nonce and the transactions resend mechanism
type NonceTransactionsHandler interface {
	ApplyNonceAndGasPrice(ctx context.Context, address core.AddressHandler, tx *transaction.FrontendTransaction) error
	SendTransaction(ctx context.Context, tx *transaction.FrontendTransaction) (string, error)
	Close() error
	IsInterfaceNil() bool
}

// ScCallsExecuteFilter defines the operations supported by a filter that allows selective executions of batches
type ScCallsExecuteFilter interface {
	ShouldExecute(callData parsers.ProxySCCompleteCallData) bool
	IsInterfaceNil() bool
}

// Codec defines the operations implemented by a DharitrI codec
type Codec interface {
	DecodeProxySCCompleteCallData(buff []byte) (parsers.ProxySCCompleteCallData, error)
	ExtractGasLimitFromRawCallData(buff []byte) (uint64, error)
	IsInterfaceNil() bool
}
