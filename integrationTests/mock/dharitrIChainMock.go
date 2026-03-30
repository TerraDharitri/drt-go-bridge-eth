package mock

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/TerraDharitri/drt-go-bridge-eth/integrationTests"
	"github.com/TerraDharitri/drt-go-chain-core/core"
	"github.com/TerraDharitri/drt-go-chain-core/data/api"
	"github.com/TerraDharitri/drt-go-chain-core/data/transaction"
	logger "github.com/TerraDharitri/drt-go-chain-logger"
	sdkCore "github.com/TerraDharitri/drt-go-sdk/core"
	"github.com/TerraDharitri/drt-go-sdk/data"
)

var log = logger.GetOrCreate("integrationTests/mock")

// DharitrIChainMock -
type DharitrIChainMock struct {
	*dharitrIContractStateMock
	mutState         sync.RWMutex
	sentTransactions map[string]*transaction.FrontendTransaction
	accounts         *dharitrIAccountsMock
}

// NewDharitrIChainMock -
func NewDharitrIChainMock() *DharitrIChainMock {
	return &DharitrIChainMock{
		dharitrIContractStateMock: newDharitrIContractStateMock(),
		sentTransactions:            make(map[string]*transaction.FrontendTransaction),
		accounts:                    newDharitrIAccountsMock(),
	}
}

// GetNetworkConfig -
func (mock *DharitrIChainMock) GetNetworkConfig(_ context.Context) (*data.NetworkConfig, error) {
	return &data.NetworkConfig{
		ChainID:                  "t",
		LatestTagSoftwareVersion: "",
		MinGasPrice:              1000000000,
		MinTransactionVersion:    1,
	}, nil
}

// GetNetworkStatus -
func (mock *DharitrIChainMock) GetNetworkStatus(_ context.Context, _ uint32) (*data.NetworkStatus, error) {
	return &data.NetworkStatus{}, nil
}

// GetShardOfAddress -
func (mock *DharitrIChainMock) GetShardOfAddress(_ context.Context, _ string) (uint32, error) {
	return 0, nil
}

// SendTransaction -
func (mock *DharitrIChainMock) SendTransaction(_ context.Context, transaction *transaction.FrontendTransaction) (string, error) {
	if transaction == nil {
		panic("nil transaction")
	}

	addrAsBech32 := transaction.Sender
	addressHandler, err := data.NewAddressFromBech32String(addrAsBech32)
	if err != nil {
		panic(fmt.Sprintf("%v while creating address handler for string %s", err, addrAsBech32))
	}

	hash, err := core.CalculateHash(integrationTests.TestMarshalizer, integrationTests.TestHasher, transaction)
	if err != nil {
		panic(err)
	}

	log.Info("sent DharitrI transaction", "sender", addrAsBech32, "data", string(transaction.Data))

	mock.mutState.Lock()
	defer mock.mutState.Unlock()

	mock.sentTransactions[string(hash)] = transaction
	mock.accounts.updateNonce(addressHandler, transaction.Nonce)

	mock.processTransaction(transaction)

	return hex.EncodeToString(hash), nil
}

// SendTransactions -
func (mock *DharitrIChainMock) SendTransactions(ctx context.Context, txs []*transaction.FrontendTransaction) ([]string, error) {
	hashes := make([]string, 0, len(txs))
	for _, tx := range txs {
		hash, _ := mock.SendTransaction(ctx, tx)
		hashes = append(hashes, hash)
	}

	return hashes, nil
}

// GetAllSentTransactions -
func (mock *DharitrIChainMock) GetAllSentTransactions(_ context.Context) map[string]*transaction.FrontendTransaction {
	mock.mutState.RLock()
	defer mock.mutState.RUnlock()

	txs := make(map[string]*transaction.FrontendTransaction)
	for hash, tx := range mock.sentTransactions {
		txs[hash] = tx
	}

	return txs
}

// ExecuteVMQuery -
func (mock *DharitrIChainMock) ExecuteVMQuery(_ context.Context, vmRequest *data.VmValueRequest) (*data.VmValuesResponseData, error) {
	mock.mutState.Lock()
	defer mock.mutState.Unlock()

	return mock.processVmRequests(vmRequest)
}

// GetAccount -
func (mock *DharitrIChainMock) GetAccount(_ context.Context, address sdkCore.AddressHandler) (*data.Account, error) {
	mock.mutState.Lock()
	defer mock.mutState.Unlock()

	return mock.accounts.getOrCreate(address), nil
}

// GetTransactionInfoWithResults -
func (mock *DharitrIChainMock) GetTransactionInfoWithResults(_ context.Context, _ string) (*data.TransactionInfo, error) {
	return &data.TransactionInfo{}, nil
}

// ProcessTransactionStatus -
func (mock *DharitrIChainMock) ProcessTransactionStatus(_ context.Context, _ string) (transaction.TxStatus, error) {
	return "", nil
}

// AddRelayer -
func (mock *DharitrIChainMock) AddRelayer(address sdkCore.AddressHandler) {
	mock.mutState.Lock()
	defer mock.mutState.Unlock()

	mock.relayers = append(mock.relayers, address.AddressBytes())
}

// SetLastExecutedEthBatchID -
func (mock *DharitrIChainMock) SetLastExecutedEthBatchID(lastExecutedEthBatchId uint64) {
	mock.mutState.Lock()
	defer mock.mutState.Unlock()

	mock.lastExecutedEthBatchId = lastExecutedEthBatchId
}

// SetLastExecutedEthTxId -
func (mock *DharitrIChainMock) SetLastExecutedEthTxId(lastExecutedEthTxId uint64) {
	mock.mutState.Lock()
	defer mock.mutState.Unlock()

	mock.lastExecutedEthTxId = lastExecutedEthTxId
}

// AddTokensPair -
func (mock *DharitrIChainMock) AddTokensPair(erc20 common.Address, ticker string, isNativeToken, isMintBurnToken bool, totalBalance, mintBalances, burnBalances *big.Int) {
	mock.mutState.Lock()
	defer mock.mutState.Unlock()

	mock.addTokensPair(erc20, ticker, isNativeToken, isMintBurnToken, totalBalance, mintBalances, burnBalances)
}

// SetQuorum -
func (mock *DharitrIChainMock) SetQuorum(quorum int) {
	mock.mutState.Lock()
	defer mock.mutState.Unlock()

	mock.quorum = quorum
}

// PerformedActionID returns the performed action ID
func (mock *DharitrIChainMock) PerformedActionID() *big.Int {
	mock.mutState.RLock()
	defer mock.mutState.RUnlock()

	return mock.performedAction
}

// ProposedTransfer returns the proposed transfer that matches the performed action ID
func (mock *DharitrIChainMock) ProposedTransfer() *dharitrIProposedTransfer {
	mock.mutState.RLock()
	defer mock.mutState.RUnlock()

	if mock.performedAction == nil {
		return nil
	}

	for hash, transfer := range mock.proposedTransfers {
		if HashToActionID(hash).String() == mock.performedAction.String() {
			return transfer
		}
	}

	return nil
}

// SetPendingBatch -
func (mock *DharitrIChainMock) SetPendingBatch(pendingBatch *DharitrIPendingBatch) {
	mock.mutState.Lock()
	mock.setPendingBatch(pendingBatch)
	mock.mutState.Unlock()
}

// AddDepositToCurrentBatch -
func (mock *DharitrIChainMock) AddDepositToCurrentBatch(deposit DharitrIDeposit) {
	mock.mutState.Lock()
	mock.pendingBatch.DharitrIDeposits = append(mock.pendingBatch.DharitrIDeposits, deposit)
	mock.mutState.Unlock()
}

// GetDCDTTokenData -
func (mock *DharitrIChainMock) GetDCDTTokenData(_ context.Context, _ sdkCore.AddressHandler, tokenIdentifier string, _ api.AccountQueryOptions) (*data.DCDTFungibleTokenData, error) {
	mock.mutState.RLock()
	defer mock.mutState.RUnlock()

	isMintBurn, found := mock.mintBurnTokens[tokenIdentifier]
	balance := mock.totalBalances[tokenIdentifier]
	if found && isMintBurn {
		balance = big.NewInt(0)
	}

	return &data.DCDTFungibleTokenData{
		TokenIdentifier: tokenIdentifier,
		Balance:         balance.String(),
	}, nil
}

// IsInterfaceNil -
func (mock *DharitrIChainMock) IsInterfaceNil() bool {
	return mock == nil
}
