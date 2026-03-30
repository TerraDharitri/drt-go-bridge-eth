package dharitritoeth

import (
	"context"
	"github.com/TerraDharitri/drt-go-bridge-eth/bridges/steps"

	"github.com/TerraDharitri/drt-go-bridge-eth/core"
	logger "github.com/TerraDharitri/drt-go-chain-logger"
)

type proposeSetStatusStep struct {
	bridge steps.Executor
}

// Execute will execute this step returning the next step to be executed
func (step *proposeSetStatusStep) Execute(ctx context.Context) core.StepIdentifier {
	batch := step.bridge.GetStoredBatch()
	if batch == nil {
		step.bridge.PrintInfo(logger.LogDebug, "nil batch stored")
		return GettingPendingBatchFromDharitrI
	}

	if step.bridge.ProcessMaxRetriesOnWasTransferProposedOnDharitrI() {
		step.bridge.PrintInfo(logger.LogDebug, "max number of retries reached, resetting counter")
		return GettingPendingBatchFromDharitrI
	}

	wasSetStatusProposed, err := step.bridge.WasSetStatusProposedOnDharitrI(ctx)
	if err != nil {
		step.bridge.PrintInfo(logger.LogError, "error determining if the set status action was proposed or not on DharitrI",
			"batch ID", batch.ID, "error", err)
		return GettingPendingBatchFromDharitrI
	}

	if wasSetStatusProposed {
		return SigningProposedSetStatusOnDharitrI
	}

	if !step.bridge.MyTurnAsLeader() {
		step.bridge.PrintInfo(logger.LogDebug, "not my turn as leader in this round")
		return step.Identifier()
	}

	err = step.bridge.ProposeSetStatusOnDharitrI(ctx)
	if err != nil {
		step.bridge.PrintInfo(logger.LogError, "error proposing transfer on DharitrI",
			"batch ID", batch.ID, "error", err)
		return GettingPendingBatchFromDharitrI
	}

	return SigningProposedSetStatusOnDharitrI
}

// Identifier returns the step's identifier
func (step *proposeSetStatusStep) Identifier() core.StepIdentifier {
	return ProposingSetStatusOnDharitrI
}

// IsInterfaceNil returns true if there is no value under the interface
func (step *proposeSetStatusStep) IsInterfaceNil() bool {
	return step == nil
}
