package bridge

import (
	"context"
	"fmt"
	"math/big"
	"runtime"
	"strings"
	"sync"

	bridgeCore "github.com/TerraDharitri/drt-go-bridge-eth/core"
	"github.com/TerraDharitri/drt-go-bridge-eth/core/batchProcessor"
	logger "github.com/TerraDharitri/drt-go-chain-logger"
)

// BridgeExecutorStub -
type BridgeExecutorStub struct {
	functionCalledCounter map[string]int
	mutExecutor           sync.RWMutex
	fullPath              string

	PrintInfoCalled                                            func(logLevel logger.LogLevel, message string, extras ...interface{})
	MyTurnAsLeaderCalled                                       func() bool
	GetBatchFromDharitrICalled                               func(ctx context.Context) (*bridgeCore.TransferBatch, error)
	StoreBatchFromDharitrICalled                             func(batch *bridgeCore.TransferBatch) error
	GetStoredBatchCalled                                       func() *bridgeCore.TransferBatch
	GetLastExecutedPeerBatchIDFromDharitrICalled             func(ctx context.Context) (uint64, error)
	VerifyLastDepositNonceExecutedOnPeerBatchCalled            func(ctx context.Context) error
	GetAndStoreActionIDForProposeTransferOnDharitrICalled    func(ctx context.Context) (uint64, error)
	GetAndStoreActionIDForProposeSetStatusFromDharitrICalled func(ctx context.Context) (uint64, error)
	GetStoredActionIDCalled                                    func() uint64
	WasTransferProposedOnDharitrICalled                      func(ctx context.Context) (bool, error)
	ProposeTransferOnDharitrICalled                          func(ctx context.Context) error
	ProcessMaxRetriesOnWasTransferProposedOnDharitrICalled   func() bool
	ResetRetriesOnWasTransferProposedOnDharitrICalled        func()
	WasSetStatusProposedOnDharitrICalled                     func(ctx context.Context) (bool, error)
	ProposeSetStatusOnDharitrICalled                         func(ctx context.Context) error
	WasActionSignedOnDharitrICalled                          func(ctx context.Context) (bool, error)
	SignActionOnDharitrICalled                               func(ctx context.Context) error
	ProcessQuorumReachedOnDharitrICalled                     func(ctx context.Context) (bool, error)
	WasActionPerformedOnDharitrICalled                       func(ctx context.Context) (bool, error)
	PerformActionOnDharitrICalled                            func(ctx context.Context) error
	ResolveNewDepositsStatusesCalled                           func(numDeposits uint64)
	ProcessMaxQuorumRetriesOnDharitrICalled                  func() bool
	ResetRetriesCountOnDharitrICalled                        func()
	GetAndStoreBatchFromPeerChainCalled                        func(ctx context.Context, nonce uint64) error
	WasTransferPerformedOnPeerChainCalled                      func(ctx context.Context) (bool, error)
	SignTransferOnPeerChainCalled                              func() error
	PerformTransferOnPeerChainCalled                           func(ctx context.Context) error
	ProcessQuorumReachedOnPeerChainCalled                      func(ctx context.Context) (bool, error)
	WaitForTransferConfirmationCalled                          func(ctx context.Context)
	WaitAndReturnFinalBatchStatusesCalled                      func(ctx context.Context) []byte
	GetBatchStatusesFromPeerChainCalled                        func(ctx context.Context) ([]byte, error)
	ProcessMaxQuorumRetriesOnPeerChainCalled                   func() bool
	ResetRetriesCountOnPeerChainCalled                         func()
	ClearStoredP2PSignaturesForPeerChainCalled                 func()
	CheckDharitrIClientAvailabilityCalled                    func(ctx context.Context) error
	CheckPeerClientAvailabilityCalled                          func(ctx context.Context) error
	CheckAvailableTokensCalled                                 func(ctx context.Context, peerTokens [][]byte, drtTokens [][]byte, amounts []*big.Int, direction batchProcessor.Direction) error
}

// NewBridgeExecutorStub creates a new BridgeExecutorStub instance
func NewBridgeExecutorStub() *BridgeExecutorStub {
	return &BridgeExecutorStub{
		functionCalledCounter: make(map[string]int),
		fullPath:              "github.com/TerraDharitri/drt-go-bridge-eth/testsCommon/bridge.(*BridgeExecutorStub).",
	}
}

// PrintInfo -
func (stub *BridgeExecutorStub) PrintInfo(logLevel logger.LogLevel, message string, extras ...interface{}) {
	stub.incrementFunctionCounter()
	if stub.PrintInfoCalled != nil {
		stub.PrintInfoCalled(logLevel, message, extras...)
	}
}

// MyTurnAsLeader -
func (stub *BridgeExecutorStub) MyTurnAsLeader() bool {
	stub.incrementFunctionCounter()
	if stub.MyTurnAsLeaderCalled != nil {
		return stub.MyTurnAsLeaderCalled()
	}
	return false
}

// GetBatchFromDharitrI -
func (stub *BridgeExecutorStub) GetBatchFromDharitrI(ctx context.Context) (*bridgeCore.TransferBatch, error) {
	stub.incrementFunctionCounter()
	if stub.GetBatchFromDharitrICalled != nil {
		return stub.GetBatchFromDharitrICalled(ctx)
	}
	return nil, errNotImplemented
}

// StoreBatchFromDharitrI -
func (stub *BridgeExecutorStub) StoreBatchFromDharitrI(batch *bridgeCore.TransferBatch) error {
	stub.incrementFunctionCounter()
	if stub.StoreBatchFromDharitrICalled != nil {
		return stub.StoreBatchFromDharitrICalled(batch)
	}
	return errNotImplemented
}

// GetStoredBatch -
func (stub *BridgeExecutorStub) GetStoredBatch() *bridgeCore.TransferBatch {
	stub.incrementFunctionCounter()
	if stub.GetStoredBatchCalled != nil {
		return stub.GetStoredBatchCalled()
	}
	return nil
}

// GetLastExecutedPeerBatchIDFromDharitrI -
func (stub *BridgeExecutorStub) GetLastExecutedPeerBatchIDFromDharitrI(ctx context.Context) (uint64, error) {
	stub.incrementFunctionCounter()
	if stub.GetLastExecutedPeerBatchIDFromDharitrICalled != nil {
		return stub.GetLastExecutedPeerBatchIDFromDharitrICalled(ctx)
	}
	return 0, errNotImplemented
}

// VerifyLastDepositNonceExecutedOnPeerBatch -
func (stub *BridgeExecutorStub) VerifyLastDepositNonceExecutedOnPeerBatch(ctx context.Context) error {
	stub.incrementFunctionCounter()
	if stub.VerifyLastDepositNonceExecutedOnPeerBatchCalled != nil {
		return stub.VerifyLastDepositNonceExecutedOnPeerBatchCalled(ctx)
	}
	return errNotImplemented
}

// GetAndStoreActionIDForProposeTransferOnDharitrI -
func (stub *BridgeExecutorStub) GetAndStoreActionIDForProposeTransferOnDharitrI(ctx context.Context) (uint64, error) {
	stub.incrementFunctionCounter()
	if stub.GetAndStoreActionIDForProposeTransferOnDharitrICalled != nil {
		return stub.GetAndStoreActionIDForProposeTransferOnDharitrICalled(ctx)
	}
	return 0, errNotImplemented
}

// GetAndStoreActionIDForProposeSetStatusFromDharitrI -
func (stub *BridgeExecutorStub) GetAndStoreActionIDForProposeSetStatusFromDharitrI(ctx context.Context) (uint64, error) {
	stub.incrementFunctionCounter()
	if stub.GetAndStoreActionIDForProposeSetStatusFromDharitrICalled != nil {
		return stub.GetAndStoreActionIDForProposeSetStatusFromDharitrICalled(ctx)
	}
	return 0, errNotImplemented
}

// GetStoredActionID -
func (stub *BridgeExecutorStub) GetStoredActionID() uint64 {
	stub.incrementFunctionCounter()
	if stub.GetStoredActionIDCalled != nil {
		return stub.GetStoredActionIDCalled()
	}
	return 0
}

// WasTransferProposedOnDharitrI -
func (stub *BridgeExecutorStub) WasTransferProposedOnDharitrI(ctx context.Context) (bool, error) {
	stub.incrementFunctionCounter()
	if stub.WasTransferProposedOnDharitrICalled != nil {
		return stub.WasTransferProposedOnDharitrICalled(ctx)
	}
	return false, errNotImplemented
}

// ProposeTransferOnDharitrI -
func (stub *BridgeExecutorStub) ProposeTransferOnDharitrI(ctx context.Context) error {
	stub.incrementFunctionCounter()
	if stub.ProposeTransferOnDharitrICalled != nil {
		return stub.ProposeTransferOnDharitrICalled(ctx)
	}
	return errNotImplemented
}

// ProcessMaxRetriesOnWasTransferProposedOnDharitrI -
func (stub *BridgeExecutorStub) ProcessMaxRetriesOnWasTransferProposedOnDharitrI() bool {
	stub.incrementFunctionCounter()
	if stub.ProcessMaxRetriesOnWasTransferProposedOnDharitrICalled != nil {
		return stub.ProcessMaxRetriesOnWasTransferProposedOnDharitrICalled()
	}
	return false
}

// ResetRetriesOnWasTransferProposedOnDharitrI -
func (stub *BridgeExecutorStub) ResetRetriesOnWasTransferProposedOnDharitrI() {
	stub.incrementFunctionCounter()
	if stub.ResetRetriesOnWasTransferProposedOnDharitrICalled != nil {
		stub.ResetRetriesOnWasTransferProposedOnDharitrICalled()
	}
}

// WasSetStatusProposedOnDharitrI -
func (stub *BridgeExecutorStub) WasSetStatusProposedOnDharitrI(ctx context.Context) (bool, error) {
	stub.incrementFunctionCounter()
	if stub.WasSetStatusProposedOnDharitrICalled != nil {
		return stub.WasSetStatusProposedOnDharitrICalled(ctx)
	}
	return false, errNotImplemented
}

// ProposeSetStatusOnDharitrI -
func (stub *BridgeExecutorStub) ProposeSetStatusOnDharitrI(ctx context.Context) error {
	stub.incrementFunctionCounter()
	if stub.ProposeSetStatusOnDharitrICalled != nil {
		return stub.ProposeSetStatusOnDharitrICalled(ctx)
	}
	return errNotImplemented
}

// WasActionSignedOnDharitrI -
func (stub *BridgeExecutorStub) WasActionSignedOnDharitrI(ctx context.Context) (bool, error) {
	stub.incrementFunctionCounter()
	if stub.WasActionSignedOnDharitrICalled != nil {
		return stub.WasActionSignedOnDharitrICalled(ctx)
	}
	return false, errNotImplemented
}

// SignActionOnDharitrI -
func (stub *BridgeExecutorStub) SignActionOnDharitrI(ctx context.Context) error {
	stub.incrementFunctionCounter()
	if stub.SignActionOnDharitrICalled != nil {
		return stub.SignActionOnDharitrICalled(ctx)
	}
	return errNotImplemented
}

// ProcessQuorumReachedOnDharitrI -
func (stub *BridgeExecutorStub) ProcessQuorumReachedOnDharitrI(ctx context.Context) (bool, error) {
	stub.incrementFunctionCounter()
	if stub.ProcessQuorumReachedOnDharitrICalled != nil {
		return stub.ProcessQuorumReachedOnDharitrICalled(ctx)
	}
	return false, errNotImplemented
}

// WasActionPerformedOnDharitrI -
func (stub *BridgeExecutorStub) WasActionPerformedOnDharitrI(ctx context.Context) (bool, error) {
	stub.incrementFunctionCounter()
	if stub.WasActionPerformedOnDharitrICalled != nil {
		return stub.WasActionPerformedOnDharitrICalled(ctx)
	}
	return false, errNotImplemented
}

// PerformActionOnDharitrI -
func (stub *BridgeExecutorStub) PerformActionOnDharitrI(ctx context.Context) error {
	stub.incrementFunctionCounter()
	if stub.PerformActionOnDharitrICalled != nil {
		return stub.PerformActionOnDharitrICalled(ctx)
	}
	return errNotImplemented
}

// ResolveNewDepositsStatuses -
func (stub *BridgeExecutorStub) ResolveNewDepositsStatuses(numDeposits uint64) {
	stub.incrementFunctionCounter()
	if stub.ResolveNewDepositsStatusesCalled != nil {
		stub.ResolveNewDepositsStatusesCalled(numDeposits)
	}
}

// ProcessMaxQuorumRetriesOnDharitrI -
func (stub *BridgeExecutorStub) ProcessMaxQuorumRetriesOnDharitrI() bool {
	stub.incrementFunctionCounter()
	if stub.ProcessMaxQuorumRetriesOnDharitrICalled != nil {
		return stub.ProcessMaxQuorumRetriesOnDharitrICalled()
	}
	return false
}

// ResetRetriesCountOnDharitrI -
func (stub *BridgeExecutorStub) ResetRetriesCountOnDharitrI() {
	stub.incrementFunctionCounter()
	if stub.ResetRetriesCountOnDharitrICalled != nil {
		stub.ResetRetriesCountOnDharitrICalled()
	}
}

// GetAndStoreBatchFromPeerChain -
func (stub *BridgeExecutorStub) GetAndStoreBatchFromPeerChain(ctx context.Context, nonce uint64) error {
	stub.incrementFunctionCounter()
	if stub.GetAndStoreBatchFromPeerChainCalled != nil {
		return stub.GetAndStoreBatchFromPeerChainCalled(ctx, nonce)
	}
	return errNotImplemented
}

// WasTransferPerformedOnPeerChain -
func (stub *BridgeExecutorStub) WasTransferPerformedOnPeerChain(ctx context.Context) (bool, error) {
	stub.incrementFunctionCounter()
	if stub.WasTransferPerformedOnPeerChainCalled != nil {
		return stub.WasTransferPerformedOnPeerChainCalled(ctx)
	}
	return false, errNotImplemented
}

// SignTransferOnPeerChain -
func (stub *BridgeExecutorStub) SignTransferOnPeerChain() error {
	stub.incrementFunctionCounter()
	if stub.SignTransferOnPeerChainCalled != nil {
		return stub.SignTransferOnPeerChainCalled()
	}
	return errNotImplemented
}

// PerformTransferOnPeerChain -
func (stub *BridgeExecutorStub) PerformTransferOnPeerChain(ctx context.Context) error {
	stub.incrementFunctionCounter()
	if stub.PerformTransferOnPeerChainCalled != nil {
		return stub.PerformTransferOnPeerChainCalled(ctx)
	}
	return errNotImplemented
}

// ProcessQuorumReachedOnPeerChain -
func (stub *BridgeExecutorStub) ProcessQuorumReachedOnPeerChain(ctx context.Context) (bool, error) {
	stub.incrementFunctionCounter()
	if stub.ProcessQuorumReachedOnPeerChainCalled != nil {
		return stub.ProcessQuorumReachedOnPeerChainCalled(ctx)
	}
	return false, errNotImplemented
}

// WaitForTransferConfirmation -
func (stub *BridgeExecutorStub) WaitForTransferConfirmation(ctx context.Context) {
	stub.incrementFunctionCounter()
	if stub.WaitForTransferConfirmationCalled != nil {
		stub.WaitForTransferConfirmationCalled(ctx)
	}
}

// WaitAndReturnFinalBatchStatuses -
func (stub *BridgeExecutorStub) WaitAndReturnFinalBatchStatuses(ctx context.Context) []byte {
	stub.incrementFunctionCounter()
	if stub.WaitAndReturnFinalBatchStatusesCalled != nil {
		return stub.WaitAndReturnFinalBatchStatusesCalled(ctx)
	}
	return nil
}

// GetBatchStatusesFromPeerChain -
func (stub *BridgeExecutorStub) GetBatchStatusesFromPeerChain(ctx context.Context) ([]byte, error) {
	stub.incrementFunctionCounter()
	if stub.GetBatchStatusesFromPeerChainCalled != nil {
		return stub.GetBatchStatusesFromPeerChainCalled(ctx)
	}
	return nil, errNotImplemented
}

// ProcessMaxQuorumRetriesOnPeerChain -
func (stub *BridgeExecutorStub) ProcessMaxQuorumRetriesOnPeerChain() bool {
	stub.incrementFunctionCounter()
	if stub.ProcessMaxQuorumRetriesOnPeerChainCalled != nil {
		return stub.ProcessMaxQuorumRetriesOnPeerChainCalled()
	}
	return false
}

// ResetRetriesCountOnPeerChain -
func (stub *BridgeExecutorStub) ResetRetriesCountOnPeerChain() {
	stub.incrementFunctionCounter()
	if stub.ResetRetriesCountOnPeerChainCalled != nil {
		stub.ResetRetriesCountOnPeerChainCalled()
	}
}

// ClearStoredP2PSignaturesForPeerChain -
func (stub *BridgeExecutorStub) ClearStoredP2PSignaturesForPeerChain() {
	stub.incrementFunctionCounter()
	if stub.ClearStoredP2PSignaturesForPeerChainCalled != nil {
		stub.ClearStoredP2PSignaturesForPeerChainCalled()
	}
}

// CheckDharitrIClientAvailability -
func (stub *BridgeExecutorStub) CheckDharitrIClientAvailability(ctx context.Context) error {
	if stub.CheckDharitrIClientAvailabilityCalled != nil {
		return stub.CheckDharitrIClientAvailabilityCalled(ctx)
	}
	return errNotImplemented
}

// CheckPeerClientAvailability -
func (stub *BridgeExecutorStub) CheckPeerClientAvailability(ctx context.Context) error {
	if stub.CheckPeerClientAvailabilityCalled != nil {
		return stub.CheckPeerClientAvailabilityCalled(ctx)
	}
	return errNotImplemented
}

// IsInterfaceNil -
func (stub *BridgeExecutorStub) IsInterfaceNil() bool {
	return stub == nil
}

// -------- helper functions

// incrementFunctionCounter increments the counter for the function that called it
func (stub *BridgeExecutorStub) incrementFunctionCounter() {
	stub.mutExecutor.Lock()
	defer stub.mutExecutor.Unlock()

	pc, _, _, _ := runtime.Caller(1)
	fmt.Printf("BridgeExecutorMock: called %s\n", runtime.FuncForPC(pc).Name())
	stub.functionCalledCounter[strings.ReplaceAll(runtime.FuncForPC(pc).Name(), stub.fullPath, "")]++
}

// GetFunctionCounter returns the called counter of a given function
func (stub *BridgeExecutorStub) GetFunctionCounter(function string) int {
	stub.mutExecutor.Lock()
	defer stub.mutExecutor.Unlock()

	return stub.functionCalledCounter[function]
}

// CheckAvailableTokens -
func (stub *BridgeExecutorStub) CheckAvailableTokens(ctx context.Context, peerTokens [][]byte, drtTokens [][]byte, amounts []*big.Int, direction batchProcessor.Direction) error {
	if stub.CheckAvailableTokensCalled != nil {
		return stub.CheckAvailableTokensCalled(ctx, peerTokens, drtTokens, amounts, direction)
	}

	return nil
}
