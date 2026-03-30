package steps

import (
	"context"
	"math/big"

	bridgeCore "github.com/TerraDharitri/drt-go-bridge-eth/core"
	"github.com/TerraDharitri/drt-go-bridge-eth/core/batchProcessor"
	logger "github.com/TerraDharitri/drt-go-chain-logger"
)

// Executor defines a generic bridge interface able to handle both halves of the bridge
type Executor interface {
	PrintInfo(logLevel logger.LogLevel, message string, extras ...interface{})
	MyTurnAsLeader() bool

	GetBatchFromDharitrI(ctx context.Context) (*bridgeCore.TransferBatch, error)
	StoreBatchFromDharitrI(batch *bridgeCore.TransferBatch) error
	GetStoredBatch() *bridgeCore.TransferBatch

	GetLastExecutedPeerBatchIDFromDharitrI(ctx context.Context) (uint64, error)
	VerifyLastDepositNonceExecutedOnPeerBatch(ctx context.Context) error

	GetAndStoreActionIDForProposeTransferOnDharitrI(ctx context.Context) (uint64, error)
	GetAndStoreActionIDForProposeSetStatusFromDharitrI(ctx context.Context) (uint64, error)
	GetStoredActionID() uint64

	WasTransferProposedOnDharitrI(ctx context.Context) (bool, error)
	ProposeTransferOnDharitrI(ctx context.Context) error
	ProcessMaxRetriesOnWasTransferProposedOnDharitrI() bool
	ResetRetriesOnWasTransferProposedOnDharitrI()

	WasSetStatusProposedOnDharitrI(ctx context.Context) (bool, error)
	ProposeSetStatusOnDharitrI(ctx context.Context) error

	WasActionSignedOnDharitrI(ctx context.Context) (bool, error)
	SignActionOnDharitrI(ctx context.Context) error

	ProcessQuorumReachedOnDharitrI(ctx context.Context) (bool, error)
	WasActionPerformedOnDharitrI(ctx context.Context) (bool, error)
	PerformActionOnDharitrI(ctx context.Context) error
	ResolveNewDepositsStatuses(numDeposits uint64)

	ProcessMaxQuorumRetriesOnDharitrI() bool
	ResetRetriesCountOnDharitrI()

	GetAndStoreBatchFromPeerChain(ctx context.Context, nonce uint64) error
	WasTransferPerformedOnPeerChain(ctx context.Context) (bool, error)
	SignTransferOnPeerChain() error
	PerformTransferOnPeerChain(ctx context.Context) error
	ProcessQuorumReachedOnPeerChain(ctx context.Context) (bool, error)
	WaitForTransferConfirmation(ctx context.Context)
	WaitAndReturnFinalBatchStatuses(ctx context.Context) []byte
	GetBatchStatusesFromPeerChain(ctx context.Context) ([]byte, error)

	ProcessMaxQuorumRetriesOnPeerChain() bool
	ResetRetriesCountOnPeerChain()
	ClearStoredP2PSignaturesForPeerChain()

	CheckDharitrIClientAvailability(ctx context.Context) error
	CheckPeerClientAvailability(ctx context.Context) error
	CheckAvailableTokens(ctx context.Context, peerTokens [][]byte, drtTokens [][]byte, amounts []*big.Int, direction batchProcessor.Direction) error

	IsInterfaceNil() bool
}
