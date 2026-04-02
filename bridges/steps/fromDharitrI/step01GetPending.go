package dharitritoeth

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
		step.bridge.PrintInfo(logger.LogDebug, "PeerChain client unavailable", "message", err)
	}
	step.bridge.ResetRetriesCountOnPeerChain()
	step.resetCountersOnDharitrI()

	batch, err := step.bridge.GetBatchFromDharitrI(ctx)
	if err != nil {
		step.bridge.PrintInfo(logger.LogDebug, "cannot fetch DharitrI batch", "message", err)
		return step.Identifier()
	}
	if batch == nil {
		step.bridge.PrintInfo(logger.LogDebug, "no new batch found on DharitrI")
		return step.Identifier()
	}

	err = step.bridge.StoreBatchFromDharitrI(batch)
	if err != nil {
		step.bridge.PrintInfo(logger.LogError, "error storing DharitrI batch", "error", err)
		return step.Identifier()
	}

	step.bridge.PrintInfo(logger.LogInfo, "fetched new batch from DharitrI "+batch.String())

	wasPerformed, err := step.bridge.WasTransferPerformedOnPeerChain(ctx)
	if err != nil {
		step.bridge.PrintInfo(logger.LogError, "error determining if transfer was performed or not", "error", err)
		return step.Identifier()
	}
	if wasPerformed {
		step.bridge.PrintInfo(logger.LogInfo, "transfer performed")
		return ResolvingSetStatusOnDharitrI
	}

	argLists := batchProcessor.ExtractListFromDrt(batch)
	err = step.bridge.CheckAvailableTokens(ctx, argLists.PeerTokens, argLists.DrtTokenBytes, argLists.Amounts, argLists.Direction)
	if err != nil {
		step.bridge.PrintInfo(logger.LogError, "error checking available tokens", "error", err, "batch", batch.String())
		return step.Identifier()
	}

	return SigningProposedTransferOnPeerChain
}

// Identifier returns the step's identifier
func (step *getPendingStep) Identifier() core.StepIdentifier {
	return GettingPendingBatchFromDharitrI
}

// IsInterfaceNil returns true if there is no value under the interface
func (step *getPendingStep) IsInterfaceNil() bool {
	return step == nil
}

func (step *getPendingStep) resetCountersOnDharitrI() {
	step.bridge.ResetRetriesCountOnDharitrI()
	step.bridge.ResetRetriesOnWasTransferProposedOnDharitrI()
}
