package ethtodharitri

import (
	"context"
	"github.com/TerraDharitri/drt-go-bridge-eth/bridges/steps"

	"github.com/TerraDharitri/drt-go-bridge-eth/core"
	"github.com/TerraDharitri/drt-go-bridge-eth/core/batchProcessor"
	logger "github.com/TerraDharitri/drt-go-chain-logger"
)

type getPendingStep struct {
	bridge steps.Executor
}

// Execute will execute this step returning the next step to be executed
func (step *getPendingStep) Execute(ctx context.Context) core.StepIdentifier {
	err := step.bridge.CheckDharitrIClientAvailability(ctx)
	if err != nil {
		step.bridge.PrintInfo(logger.LogDebug, "DharitrI client unavailable", "message", err)
	}
	err = step.bridge.CheckPeerClientAvailability(ctx)
	if err != nil {
		step.bridge.PrintInfo(logger.LogDebug, "Peer client unavailable", "message", err)
	}
	step.bridge.ResetRetriesCountOnDharitrI()
	lastEthBatchExecuted, err := step.bridge.GetLastExecutedPeerBatchIDFromDharitrI(ctx)
	if err != nil {
		step.bridge.PrintInfo(logger.LogError, "error fetching last executed eth batch ID", "error", err)
		return step.Identifier()
	}

	err = step.bridge.GetAndStoreBatchFromPeerChain(ctx, lastEthBatchExecuted+1)
	if err != nil {
		step.bridge.PrintInfo(logger.LogDebug, "cannot fetch peer chain batch", "batch ID", lastEthBatchExecuted+1, "message", err)
		return step.Identifier()
	}

	batch := step.bridge.GetStoredBatch()
	if batch == nil {
		step.bridge.PrintInfo(logger.LogDebug, "no new batch found on peer chain", "last executed on DharitrI", lastEthBatchExecuted)
		return step.Identifier()
	}

	step.bridge.PrintInfo(logger.LogInfo, "fetched new batch from peer chain "+batch.String())

	err = step.bridge.VerifyLastDepositNonceExecutedOnPeerBatch(ctx)
	if err != nil {
		step.bridge.PrintInfo(logger.LogError, "verification failed on the new batch from peer chain", "batch ID", lastEthBatchExecuted+1, "error", err)
		return step.Identifier()
	}

	argLists := batchProcessor.ExtractListToDrt(batch)
	err = step.bridge.CheckAvailableTokens(ctx, argLists.PeerTokens, argLists.DrtTokenBytes, argLists.Amounts, argLists.Direction)
	if err != nil {
		step.bridge.PrintInfo(logger.LogError, "error checking available tokens", "error", err, "batch", batch.String())
		return step.Identifier()
	}

	return ProposingTransferOnDharitrI
}

// Identifier returns the step's identifier
func (step *getPendingStep) Identifier() core.StepIdentifier {
	return GettingPendingBatchFromPeerChain
}

// IsInterfaceNil returns true if there is no value under the interface
func (step *getPendingStep) IsInterfaceNil() bool {
	return step == nil
}
