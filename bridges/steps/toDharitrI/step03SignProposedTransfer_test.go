package ethtodharitri

import (
	"context"
	"errors"
	"testing"

	"github.com/TerraDharitri/drt-go-bridge-eth/bridges"
	"github.com/TerraDharitri/drt-go-bridge-eth/core"
	"github.com/stretchr/testify/assert"
)

func TestExecuteSignProposedTransferStep(t *testing.T) {
	t.Parallel()

	t.Run("nil batch", func(t *testing.T) {
		t.Parallel()
		bridgeStub := createStubExecutor()
		bridgeStub.GetStoredBatchCalled = func() *core.TransferBatch {
			return nil
		}

		step := signProposedTransferStep{
			bridge: bridgeStub,
		}

		expectedStepIdentifier := core.StepIdentifier(GettingPendingBatchFromPeerChain)
		stepIdentifier := step.Execute(context.Background())
		assert.Equal(t, expectedStepIdentifier, stepIdentifier)
	})

	t.Run("error on WasProposedTransferSigned", func(t *testing.T) {
		t.Parallel()
		bridgeStub := createStubExecutor()
		bridgeStub.GetStoredBatchCalled = func() *core.TransferBatch {
			return testBatch
		}
		bridgeStub.WasActionSignedOnDharitrICalled = func(ctx context.Context) (bool, error) {
			return false, expectedError
		}

		step := signProposedTransferStep{
			bridge: bridgeStub,
		}

		expectedStepIdentifier := core.StepIdentifier(GettingPendingBatchFromPeerChain)
		stepIdentifier := step.Execute(context.Background())
		assert.Equal(t, expectedStepIdentifier, stepIdentifier)
	})

	t.Run("error on SignProposedTransfer", func(t *testing.T) {
		t.Parallel()
		bridgeStub := createStubExecutor()
		bridgeStub.GetStoredBatchCalled = func() *core.TransferBatch {
			return testBatch
		}
		bridgeStub.WasActionSignedOnDharitrICalled = func(ctx context.Context) (bool, error) {
			return false, nil
		}
		bridgeStub.SignActionOnDharitrICalled = func(ctx context.Context) error {
			return expectedError
		}

		step := signProposedTransferStep{
			bridge: bridgeStub,
		}

		expectedStepIdentifier := core.StepIdentifier(GettingPendingBatchFromPeerChain)
		stepIdentifier := step.Execute(context.Background())
		assert.Equal(t, expectedStepIdentifier, stepIdentifier)
	})

	t.Run("get action ID errors", func(t *testing.T) {
		t.Parallel()
		expectedErr := errors.New("expected error")
		bridgeStub := createStubExecutor()
		bridgeStub.GetStoredBatchCalled = func() *core.TransferBatch {
			return testBatch
		}
		bridgeStub.WasActionSignedOnDharitrICalled = func(ctx context.Context) (bool, error) {
			return true, nil
		}
		bridgeStub.GetAndStoreActionIDForProposeTransferOnDharitrICalled = func(ctx context.Context) (uint64, error) {
			return 0, expectedErr
		}

		step := signProposedTransferStep{
			bridge: bridgeStub,
		}

		expectedStepIdentifier := core.StepIdentifier(GettingPendingBatchFromPeerChain)
		stepIdentifier := step.Execute(context.Background())
		assert.Equal(t, expectedStepIdentifier, stepIdentifier)
	})

	t.Run("invalid action ID", func(t *testing.T) {
		t.Parallel()
		bridgeStub := createStubExecutor()
		bridgeStub.GetStoredBatchCalled = func() *core.TransferBatch {
			return testBatch
		}
		bridgeStub.WasActionSignedOnDharitrICalled = func(ctx context.Context) (bool, error) {
			return true, nil
		}
		bridgeStub.GetAndStoreActionIDForProposeTransferOnDharitrICalled = func(ctx context.Context) (uint64, error) {
			return bridges.InvalidActionID, nil
		}

		step := signProposedTransferStep{
			bridge: bridgeStub,
		}

		expectedStepIdentifier := core.StepIdentifier(GettingPendingBatchFromPeerChain)
		stepIdentifier := step.Execute(context.Background())
		assert.Equal(t, expectedStepIdentifier, stepIdentifier)
	})

	t.Run("error on WasActionSignedOnDharitrI", func(t *testing.T) {
		t.Parallel()
		bridgeStub := createStubExecutor()
		bridgeStub.GetStoredBatchCalled = func() *core.TransferBatch {
			return testBatch
		}
		bridgeStub.WasActionSignedOnDharitrICalled = func(ctx context.Context) (bool, error) {
			return false, expectedError
		}
		bridgeStub.GetAndStoreActionIDForProposeTransferOnDharitrICalled = func(ctx context.Context) (uint64, error) {
			return bridges.InvalidActionID, nil
		}

		step := signProposedTransferStep{
			bridge: bridgeStub,
		}

		expectedStepIdentifier := core.StepIdentifier(GettingPendingBatchFromPeerChain)
		stepIdentifier := step.Execute(context.Background())
		assert.Equal(t, expectedStepIdentifier, stepIdentifier)
	})

	t.Run("should work - transfer was already signed", func(t *testing.T) {
		t.Parallel()
		bridgeStub := createStubExecutor()
		bridgeStub.GetStoredBatchCalled = func() *core.TransferBatch {
			return testBatch
		}
		bridgeStub.WasActionSignedOnDharitrICalled = func(ctx context.Context) (bool, error) {
			return true, nil
		}
		bridgeStub.GetAndStoreActionIDForProposeTransferOnDharitrICalled = func(ctx context.Context) (uint64, error) {
			return 2, nil
		}

		step := signProposedTransferStep{
			bridge: bridgeStub,
		}

		expectedStepIdentifier := core.StepIdentifier(WaitingForQuorum)
		stepIdentifier := step.Execute(context.Background())
		assert.Equal(t, expectedStepIdentifier, stepIdentifier)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		bridgeStub := createStubExecutor()
		bridgeStub.GetStoredBatchCalled = func() *core.TransferBatch {
			return testBatch
		}
		bridgeStub.WasActionSignedOnDharitrICalled = func(ctx context.Context) (bool, error) {
			return false, nil
		}
		bridgeStub.SignActionOnDharitrICalled = func(ctx context.Context) error {
			return nil
		}
		bridgeStub.GetAndStoreActionIDForProposeTransferOnDharitrICalled = func(ctx context.Context) (uint64, error) {
			return 2, nil
		}

		step := signProposedTransferStep{
			bridge: bridgeStub,
		}
		// Test Identifier()
		expectedStepIdentifier := core.StepIdentifier(SigningProposedTransferOnDharitrI)
		assert.Equal(t, expectedStepIdentifier, step.Identifier())
		// Test IsInterfaceNil
		assert.NotNil(t, step.IsInterfaceNil())

		expectedStepIdentifier = WaitingForQuorum
		stepIdentifier := step.Execute(context.Background())
		assert.Equal(t, expectedStepIdentifier, stepIdentifier)
	})
}
