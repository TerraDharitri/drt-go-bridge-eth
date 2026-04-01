package bridge

import (
	"context"
	"math/big"

	bridgeCore "github.com/TerraDharitri/drt-go-bridge-eth/core"
)

// DharitrIClientStub -
type DharitrIClientStub struct {
	GetPendingBatchCalled                          func(ctx context.Context) (*bridgeCore.TransferBatch, error)
	GetBatchCalled                                 func(ctx context.Context, batchID uint64) (*bridgeCore.TransferBatch, error)
	GetCurrentBatchAsDataBytesCalled               func(ctx context.Context) ([][]byte, error)
	WasProposedTransferCalled                      func(ctx context.Context, batch *bridgeCore.TransferBatch) (bool, error)
	QuorumReachedCalled                            func(ctx context.Context, actionID uint64) (bool, error)
	WasExecutedCalled                              func(ctx context.Context, actionID uint64) (bool, error)
	GetActionIDForProposeTransferCalled            func(ctx context.Context, batch *bridgeCore.TransferBatch) (uint64, error)
	WasProposedSetStatusCalled                     func(ctx context.Context, batch *bridgeCore.TransferBatch) (bool, error)
	GetTransactionsStatusesCalled                  func(ctx context.Context, batchID uint64) ([]byte, error)
	GetActionIDForSetStatusOnPendingTransferCalled func(ctx context.Context, batch *bridgeCore.TransferBatch) (uint64, error)
	GetLastExecutedEthBatchIDCalled                func(ctx context.Context) (uint64, error)
	GetLastExecutedEthTxIDCalled                   func(ctx context.Context) (uint64, error)
	GetCurrentNonceCalled                          func(ctx context.Context) (uint64, error)
	ProposeSetStatusCalled                         func(ctx context.Context, batch *bridgeCore.TransferBatch) (string, error)
	ResolveNewDepositsCalled                       func(ctx context.Context, batch *bridgeCore.TransferBatch) error
	ProposeTransferCalled                          func(ctx context.Context, batch *bridgeCore.TransferBatch) (string, error)
	SignCalled                                     func(ctx context.Context, actionID uint64) (string, error)
	WasSignedCalled                                func(ctx context.Context, actionID uint64) (bool, error)
	PerformActionCalled                            func(ctx context.Context, actionID uint64, batch *bridgeCore.TransferBatch) (string, error)
	CheckClientAvailabilityCalled                  func(ctx context.Context) error
	IsMintBurnTokenCalled                          func(ctx context.Context, token []byte) (bool, error)
	IsNativeTokenCalled                            func(ctx context.Context, token []byte) (bool, error)
	TotalBalancesCalled                            func(ctx context.Context, token []byte) (*big.Int, error)
	MintBalancesCalled                             func(ctx context.Context, token []byte) (*big.Int, error)
	BurnBalancesCalled                             func(ctx context.Context, token []byte) (*big.Int, error)
	CheckRequiredBalanceCalled                     func(ctx context.Context, token []byte, value *big.Int) error
	GetLastDrtBatchIDCalled                        func(ctx context.Context) (uint64, error)
	CloseCalled                                    func() error
}

// GetPendingBatch -
func (stub *DharitrIClientStub) GetPendingBatch(ctx context.Context) (*bridgeCore.TransferBatch, error) {
	if stub.GetPendingBatchCalled != nil {
		return stub.GetPendingBatchCalled(ctx)
	}

	return nil, errNotImplemented
}

// GetBatch -
func (stub *DharitrIClientStub) GetBatch(ctx context.Context, batchID uint64) (*bridgeCore.TransferBatch, error) {
	if stub.GetBatchCalled != nil {
		return stub.GetBatchCalled(ctx, batchID)
	}

	return nil, errNotImplemented
}

// GetCurrentBatchAsDataBytes -
func (stub *DharitrIClientStub) GetCurrentBatchAsDataBytes(ctx context.Context) ([][]byte, error) {
	if stub.GetCurrentBatchAsDataBytesCalled != nil {
		return stub.GetCurrentBatchAsDataBytesCalled(ctx)
	}

	return make([][]byte, 0), nil
}

// WasProposedTransfer -
func (stub *DharitrIClientStub) WasProposedTransfer(ctx context.Context, batch *bridgeCore.TransferBatch) (bool, error) {
	if stub.WasProposedTransferCalled != nil {
		return stub.WasProposedTransferCalled(ctx, batch)
	}

	return false, nil
}

// QuorumReached -
func (stub *DharitrIClientStub) QuorumReached(ctx context.Context, actionID uint64) (bool, error) {
	if stub.QuorumReachedCalled != nil {
		return stub.QuorumReachedCalled(ctx, actionID)
	}

	return false, nil
}

// WasExecuted -
func (stub *DharitrIClientStub) WasExecuted(ctx context.Context, actionID uint64) (bool, error) {
	if stub.WasExecutedCalled != nil {
		return stub.WasExecutedCalled(ctx, actionID)
	}

	return false, nil
}

// GetActionIDForProposeTransfer -
func (stub *DharitrIClientStub) GetActionIDForProposeTransfer(ctx context.Context, batch *bridgeCore.TransferBatch) (uint64, error) {
	if stub.GetActionIDForProposeTransferCalled != nil {
		return stub.GetActionIDForProposeTransferCalled(ctx, batch)
	}

	return 0, nil
}

// WasProposedSetStatus -
func (stub *DharitrIClientStub) WasProposedSetStatus(ctx context.Context, batch *bridgeCore.TransferBatch) (bool, error) {
	if stub.WasProposedSetStatusCalled != nil {
		return stub.WasProposedSetStatusCalled(ctx, batch)
	}

	return false, nil
}

// GetTransactionsStatuses -
func (stub *DharitrIClientStub) GetTransactionsStatuses(ctx context.Context, batchID uint64) ([]byte, error) {
	if stub.GetTransactionsStatusesCalled != nil {
		return stub.GetTransactionsStatusesCalled(ctx, batchID)
	}

	return make([]byte, 0), nil
}

// GetActionIDForSetStatusOnPendingTransfer -
func (stub *DharitrIClientStub) GetActionIDForSetStatusOnPendingTransfer(ctx context.Context, batch *bridgeCore.TransferBatch) (uint64, error) {
	if stub.GetActionIDForSetStatusOnPendingTransferCalled != nil {
		return stub.GetActionIDForSetStatusOnPendingTransferCalled(ctx, batch)
	}

	return 0, nil
}

// GetLastExecutedEthBatchID -
func (stub *DharitrIClientStub) GetLastExecutedEthBatchID(ctx context.Context) (uint64, error) {
	if stub.GetLastExecutedEthBatchIDCalled != nil {
		return stub.GetLastExecutedEthBatchIDCalled(ctx)
	}

	return 0, nil
}

// GetLastExecutedEthTxID -
func (stub *DharitrIClientStub) GetLastExecutedEthTxID(ctx context.Context) (uint64, error) {
	if stub.GetLastExecutedEthTxIDCalled != nil {
		return stub.GetLastExecutedEthTxIDCalled(ctx)
	}

	return 0, nil
}

// GetCurrentNonce -
func (stub *DharitrIClientStub) GetCurrentNonce(ctx context.Context) (uint64, error) {
	if stub.GetCurrentNonceCalled != nil {
		return stub.GetCurrentNonceCalled(ctx)
	}

	return 0, nil
}

// ProposeSetStatus -
func (stub *DharitrIClientStub) ProposeSetStatus(ctx context.Context, batch *bridgeCore.TransferBatch) (string, error) {
	if stub.ProposeSetStatusCalled != nil {
		return stub.ProposeSetStatusCalled(ctx, batch)
	}

	return "", nil
}

// ProposeTransfer -
func (stub *DharitrIClientStub) ProposeTransfer(ctx context.Context, batch *bridgeCore.TransferBatch) (string, error) {
	if stub.ProposeTransferCalled != nil {
		return stub.ProposeTransferCalled(ctx, batch)
	}

	return "", nil
}

// Sign -
func (stub *DharitrIClientStub) Sign(ctx context.Context, actionID uint64) (string, error) {
	if stub.SignCalled != nil {
		return stub.SignCalled(ctx, actionID)
	}

	return "", nil
}

// WasSigned -
func (stub *DharitrIClientStub) WasSigned(ctx context.Context, actionID uint64) (bool, error) {
	if stub.WasSignedCalled != nil {
		return stub.WasSignedCalled(ctx, actionID)
	}

	return false, nil
}

// PerformAction -
func (stub *DharitrIClientStub) PerformAction(ctx context.Context, actionID uint64, batch *bridgeCore.TransferBatch) (string, error) {
	if stub.PerformActionCalled != nil {
		return stub.PerformActionCalled(ctx, actionID, batch)
	}

	return "", nil
}

// CheckClientAvailability -
func (stub *DharitrIClientStub) CheckClientAvailability(ctx context.Context) error {
	if stub.CheckClientAvailabilityCalled != nil {
		return stub.CheckClientAvailabilityCalled(ctx)
	}

	return nil
}

// IsMintBurnToken -
func (stub *DharitrIClientStub) IsMintBurnToken(ctx context.Context, token []byte) (bool, error) {
	if stub.IsMintBurnTokenCalled != nil {
		return stub.IsMintBurnTokenCalled(ctx, token)
	}
	return false, errNotImplemented
}

// IsNativeToken -
func (stub *DharitrIClientStub) IsNativeToken(ctx context.Context, token []byte) (bool, error) {
	if stub.IsNativeTokenCalled != nil {
		return stub.IsNativeTokenCalled(ctx, token)
	}
	return false, errNotImplemented
}

// TotalBalances -
func (stub *DharitrIClientStub) TotalBalances(ctx context.Context, token []byte) (*big.Int, error) {
	if stub.TotalBalancesCalled != nil {
		return stub.TotalBalancesCalled(ctx, token)
	}
	return nil, errNotImplemented
}

// MintBalances -
func (stub *DharitrIClientStub) MintBalances(ctx context.Context, token []byte) (*big.Int, error) {
	if stub.MintBalancesCalled != nil {
		return stub.MintBalancesCalled(ctx, token)
	}
	return nil, errNotImplemented
}

// BurnBalances -
func (stub *DharitrIClientStub) BurnBalances(ctx context.Context, token []byte) (*big.Int, error) {
	if stub.BurnBalancesCalled != nil {
		return stub.BurnBalancesCalled(ctx, token)
	}
	return nil, errNotImplemented
}

// CheckRequiredBalance -
func (stub *DharitrIClientStub) CheckRequiredBalance(ctx context.Context, token []byte, value *big.Int) error {
	if stub.CheckRequiredBalanceCalled != nil {
		return stub.CheckRequiredBalanceCalled(ctx, token, value)
	}
	return nil
}

// GetLastDrtBatchID -
func (stub *DharitrIClientStub) GetLastDrtBatchID(ctx context.Context) (uint64, error) {
	if stub.GetLastDrtBatchIDCalled != nil {
		return stub.GetLastDrtBatchIDCalled(ctx)
	}

	return 0, nil
}

// Close -
func (stub *DharitrIClientStub) Close() error {
	if stub.CloseCalled != nil {
		return stub.CloseCalled()
	}

	return nil
}

// IsInterfaceNil -
func (stub *DharitrIClientStub) IsInterfaceNil() bool {
	return stub == nil
}
