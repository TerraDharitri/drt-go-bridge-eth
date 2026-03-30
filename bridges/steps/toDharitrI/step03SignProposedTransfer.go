package ethtodharitri

import (
	"context"
	"github.com/TerraDharitri/drt-go-bridge-eth/bridges"
	"github.com/TerraDharitri/drt-go-bridge-eth/bridges/steps"

	"github.com/TerraDharitri/drt-go-bridge-eth/core"
	logger "github.com/TerraDharitri/drt-go-chain-logger"
)

type signProposedTransferStep struct {
	bridge steps.Executor
}

// Execute will execute this step returning the next step to be executed
func (step *signProposedTransferStep) Execute(ctx context.Context) core.StepIdentifier {
	batch := step.bridge.GetStoredBatch()
	if batch == nil {
		step.bridge.PrintInfo(logger.LogDebug, "no batch found")
		return GettingPendingBatchFromPeerChain
	}

	actionID, err := step.bridge.GetAndStoreActionIDForProposeTransferOnDharitrI(ctx)
	if err != nil {
		step.bridge.PrintInfo(logger.LogError, "error fetching action ID", "batch ID", batch.ID, "error", err)
		return GettingPendingBatchFromPeerChain
	}
	if actionID == bridges.InvalidActionID {
		step.bridge.PrintInfo(logger.LogError, "contract error, got invalid action ID",
			"batch ID", batch.ID, "error", err, "action ID", actionID)
		return GettingPendingBatchFromPeerChain
	}

	step.bridge.PrintInfo(logger.LogInfo, "fetched action ID", "action ID", actionID, "batch ID", batch.ID)

	wasSigned, err := step.bridge.WasActionSignedOnDharitrI(ctx)
	if err != nil {
		step.bridge.PrintInfo(logger.LogError, "error determining if the proposed transfer was signed or not",
			"batch ID", batch.ID, "error", err)
		return GettingPendingBatchFromPeerChain
	}

	if wasSigned {
		return WaitingForQuorum
	}

	err = step.bridge.SignActionOnDharitrI(ctx)
	if err != nil {
		step.bridge.PrintInfo(logger.LogError, "error signing the proposed transfer",
			"batch ID", batch.ID, "error", err)
		return GettingPendingBatchFromPeerChain
	}

	return WaitingForQuorum
}

// Identifier returns the step's identifier
func (step *signProposedTransferStep) Identifier() core.StepIdentifier {
	return SigningProposedTransferOnDharitrI
}

// IsInterfaceNil returns true if there is no value under the interface
func (step *signProposedTransferStep) IsInterfaceNil() bool {
	return step == nil
}
