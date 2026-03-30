package ethtodharitri

import (
	"context"
	"github.com/TerraDharitri/drt-go-bridge-eth/bridges/steps"

	"github.com/TerraDharitri/drt-go-bridge-eth/core"
	logger "github.com/TerraDharitri/drt-go-chain-logger"
)

type proposeTransferStep struct {
	bridge steps.Executor
}

// Execute will execute this step returning the next step to be executed
func (step *proposeTransferStep) Execute(ctx context.Context) core.StepIdentifier {
	batch := step.bridge.GetStoredBatch()
	if batch == nil {
		step.bridge.PrintInfo(logger.LogDebug, "no batch found")
		return GettingPendingBatchFromPeerChain
	}

	wasTransferProposed, err := step.bridge.WasTransferProposedOnDharitrI(ctx)
	if err != nil {
		step.bridge.PrintInfo(logger.LogError, "error determining if the batch was proposed or not on DharitrI",
			"batch ID", batch.ID, "error", err)
		return GettingPendingBatchFromPeerChain
	}

	if wasTransferProposed {
		return SigningProposedTransferOnDharitrI
	}

	if !step.bridge.MyTurnAsLeader() {
		step.bridge.PrintInfo(logger.LogDebug, "not my turn as leader in this round")
		return step.Identifier()
	}

	err = step.bridge.ProposeTransferOnDharitrI(ctx)
	if err != nil {
		step.bridge.PrintInfo(logger.LogError, "error proposing transfer on DharitrI",
			"batch ID", batch.ID, "error", err)
		return GettingPendingBatchFromPeerChain
	}

	return SigningProposedTransferOnDharitrI
}

// Identifier returns the step's identifier
func (step *proposeTransferStep) Identifier() core.StepIdentifier {
	return ProposingTransferOnDharitrI
}

// IsInterfaceNil returns true if there is no value under the interface
func (step *proposeTransferStep) IsInterfaceNil() bool {
	return step == nil
}
