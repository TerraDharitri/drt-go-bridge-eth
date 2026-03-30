package dharitritoeth

import (
	"context"
	"github.com/TerraDharitri/drt-go-bridge-eth/bridges"
	"github.com/TerraDharitri/drt-go-bridge-eth/bridges/steps"

	"github.com/TerraDharitri/drt-go-bridge-eth/core"
	logger "github.com/TerraDharitri/drt-go-chain-logger"
)

type signProposedSetStatusStep struct {
	bridge steps.Executor
}

// Execute will execute this step returning the next step to be executed
func (step *signProposedSetStatusStep) Execute(ctx context.Context) core.StepIdentifier {
	storedBatch := step.bridge.GetStoredBatch()
	if storedBatch == nil {
		step.bridge.PrintInfo(logger.LogDebug, "nil stored batch")
		return GettingPendingBatchFromDharitrI
	}

	actionID, err := step.bridge.GetAndStoreActionIDForProposeSetStatusFromDharitrI(ctx)
	if err != nil {
		step.bridge.PrintInfo(logger.LogError, "error fetching action ID", "batch ID", storedBatch.ID, "error", err)
		return GettingPendingBatchFromDharitrI
	}
	if actionID == bridges.InvalidActionID {
		step.bridge.PrintInfo(logger.LogError, "contract error, got invalid action ID",
			"batch ID", storedBatch.ID, "error", err, "action ID", actionID)
		return GettingPendingBatchFromDharitrI
	}

	step.bridge.PrintInfo(logger.LogInfo, "fetched action ID", "action ID", actionID, "batch ID", storedBatch.ID)

	wasSigned, err := step.bridge.WasActionSignedOnDharitrI(ctx)
	if err != nil {
		step.bridge.PrintInfo(logger.LogError, "error determining if the proposed transfer was signed or not",
			"batch ID", storedBatch.ID, "error", err)
		return GettingPendingBatchFromDharitrI
	}

	if wasSigned {
		return WaitingForQuorumOnSetStatus
	}

	err = step.bridge.SignActionOnDharitrI(ctx)
	if err != nil {
		step.bridge.PrintInfo(logger.LogError, "error signing the proposed set status",
			"batch ID", storedBatch.ID, "error", err)
		return GettingPendingBatchFromDharitrI
	}

	return WaitingForQuorumOnSetStatus
}

// Identifier returns the step's identifier
func (step *signProposedSetStatusStep) Identifier() core.StepIdentifier {
	return SigningProposedSetStatusOnDharitrI
}

// IsInterfaceNil returns true if there is no value under the interface
func (step *signProposedSetStatusStep) IsInterfaceNil() bool {
	return step == nil
}
