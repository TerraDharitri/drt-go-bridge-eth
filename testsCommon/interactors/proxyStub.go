package interactors

import (
	"context"
	"fmt"

	"github.com/TerraDharitri/drt-go-chain-core/data/api"
	"github.com/TerraDharitri/drt-go-chain-core/data/transaction"
	"github.com/TerraDharitri/drt-go-sdk/core"
	"github.com/TerraDharitri/drt-go-sdk/data"
)

// ProxyStub -
type ProxyStub struct {
	GetNetworkConfigCalled              func(ctx context.Context) (*data.NetworkConfig, error)
	SendTransactionCalled               func(ctx context.Context, transaction *transaction.FrontendTransaction) (string, error)
	SendTransactionsCalled              func(ctx context.Context, txs []*transaction.FrontendTransaction) ([]string, error)
	ExecuteVMQueryCalled                func(ctx context.Context, vmRequest *data.VmValueRequest) (*data.VmValuesResponseData, error)
	GetAccountCalled                    func(ctx context.Context, address core.AddressHandler) (*data.Account, error)
	GetNetworkStatusCalled              func(ctx context.Context, shardID uint32) (*data.NetworkStatus, error)
	GetShardOfAddressCalled             func(ctx context.Context, bech32Address string) (uint32, error)
	GetDCDTTokenDataCalled              func(ctx context.Context, address core.AddressHandler, tokenIdentifier string, queryOptions api.AccountQueryOptions) (*data.DCDTFungibleTokenData, error)
	GetTransactionInfoWithResultsCalled func(_ context.Context, _ string) (*data.TransactionInfo, error)
	ProcessTransactionStatusCalled      func(ctx context.Context, hexTxHash string) (transaction.TxStatus, error)
}

// GetNetworkConfig -
func (eps *ProxyStub) GetNetworkConfig(ctx context.Context) (*data.NetworkConfig, error) {
	if eps.GetNetworkConfigCalled != nil {
		return eps.GetNetworkConfigCalled(ctx)
	}

	return &data.NetworkConfig{}, nil
}

// SendTransaction -
func (eps *ProxyStub) SendTransaction(ctx context.Context, transaction *transaction.FrontendTransaction) (string, error) {
	if eps.SendTransactionCalled != nil {
		return eps.SendTransactionCalled(ctx, transaction)
	}

	return "", nil
}

// SendTransactions -
func (eps *ProxyStub) SendTransactions(ctx context.Context, txs []*transaction.FrontendTransaction) ([]string, error) {
	if eps.SendTransactionCalled != nil {
		return eps.SendTransactionsCalled(ctx, txs)
	}

	return make([]string, 0), nil
}

// ExecuteVMQuery -
func (eps *ProxyStub) ExecuteVMQuery(ctx context.Context, vmRequest *data.VmValueRequest) (*data.VmValuesResponseData, error) {
	if eps.ExecuteVMQueryCalled != nil {
		return eps.ExecuteVMQueryCalled(ctx, vmRequest)
	}

	return &data.VmValuesResponseData{}, nil
}

// GetAccount -
func (eps *ProxyStub) GetAccount(ctx context.Context, address core.AddressHandler) (*data.Account, error) {
	if eps.GetAccountCalled != nil {
		return eps.GetAccountCalled(ctx, address)
	}

	return &data.Account{}, nil
}

// GetNetworkStatus -
func (eps *ProxyStub) GetNetworkStatus(ctx context.Context, shardID uint32) (*data.NetworkStatus, error) {
	if eps.GetNetworkStatusCalled != nil {
		return eps.GetNetworkStatusCalled(ctx, shardID)
	}

	return nil, fmt.Errorf("not implemented")
}

// GetShardOfAddress -
func (eps *ProxyStub) GetShardOfAddress(ctx context.Context, bech32Address string) (uint32, error) {
	if eps.GetShardOfAddressCalled != nil {
		return eps.GetShardOfAddressCalled(ctx, bech32Address)
	}

	return 0, fmt.Errorf("not implemented")
}

// GetDCDTTokenData -
func (eps *ProxyStub) GetDCDTTokenData(ctx context.Context, address core.AddressHandler, tokenIdentifier string, queryOptions api.AccountQueryOptions) (*data.DCDTFungibleTokenData, error) {
	if eps.GetDCDTTokenDataCalled != nil {
		return eps.GetDCDTTokenDataCalled(ctx, address, tokenIdentifier, queryOptions)
	}

	return nil, fmt.Errorf("not implemented")
}

// GetTransactionInfoWithResults -
func (eps *ProxyStub) GetTransactionInfoWithResults(ctx context.Context, hash string) (*data.TransactionInfo, error) {
	if eps.GetTransactionInfoWithResultsCalled != nil {
		return eps.GetTransactionInfoWithResultsCalled(ctx, hash)
	}

	return nil, fmt.Errorf("not implemented")
}

// ProcessTransactionStatus -
func (eps *ProxyStub) ProcessTransactionStatus(ctx context.Context, hexTxHash string) (transaction.TxStatus, error) {
	if eps.ProcessTransactionStatusCalled != nil {
		return eps.ProcessTransactionStatusCalled(ctx, hexTxHash)
	}

	return "", nil
}

// IsInterfaceNil -
func (eps *ProxyStub) IsInterfaceNil() bool {
	return eps == nil
}
