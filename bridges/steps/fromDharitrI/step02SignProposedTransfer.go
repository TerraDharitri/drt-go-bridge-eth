package dharitritoeth

import (
	"context"

	"github.com/TerraDharitri/drt-go-bridge-eth/bridges/steps"
	"github.com/TerraDharitri/drt-go-bridge-eth/core"
	logger "github.com/TerraDharitri/drt-go-chain-logger"
)

type signProposedTransferStep struct {
	bridge steps.Executor
}

// Execute will execute this step returning the next step to be executed
func (step *signProposedTransferStep) Execute(_ context.Context) core.StepIdentifier {
	storedBatch := step.bridge.GetStoredBatch()
	if storedBatch == nil {
		step.bridge.PrintInfo(logger.LogDebug, "nil batch stored")
		return GettingPendingBatchFromDharitrI
	}

	err := step.bridge.SignTransferOnPeerChain()
	if err != nil {
		step.bridge.PrintInfo(logger.LogError, "error signing", "batch ID", storedBatch.ID, "error", err)
		return GettingPendingBatchFromDharitrI
	}

	return WaitingForQuorumOnTransfer
}

// Identifier returns the step's identifier
func (step *signProposedTransferStep) Identifier() core.StepIdentifier {
	return SigningProposedTransferOnPeerChain
}

// IsInterfaceNil returns true if there is no value under the interface
func (step *signProposedTransferStep) IsInterfaceNil() bool {
	return step == nil
}
