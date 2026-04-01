package dharitritoeth

import (
	"context"
	"testing"

	bridgeCore "github.com/TerraDharitri/drt-go-bridge-eth/core"
	bridgeTests "github.com/TerraDharitri/drt-go-bridge-eth/testsCommon/bridge"
	"github.com/stretchr/testify/assert"
)

func TestExecute_ResolveSetStatus(t *testing.T) {
	t.Parallel()

	t.Run("nil batch on GetStoredBatch", func(t *testing.T) {
		t.Parallel()
		bridgeStub := createStubExecutorResolveSetStatus()
		bridgeStub.GetStoredBatchCalled = func() *bridgeCore.TransferBatch {
			return nil
		}
		clearWasCalled := false
		bridgeStub.ClearStoredP2PSignaturesForPeerChainCalled = func() {
			clearWasCalled = true
		}

		step := resolveSetStatusStep{
			bridge: bridgeStub,
		}

		stepIdentifier := step.Execute(context.Background())
		assert.Equal(t, initialStep, stepIdentifier)
		assert.True(t, clearWasCalled)
	})
	t.Run("error on GetStoredBatch", func(t *testing.T) {
		t.Parallel()
		bridgeStub := createStubExecutorResolveSetStatus()
		bridgeStub.GetBatchFromDharitrICalled = func(ctx context.Context) (*bridgeCore.TransferBatch, error) {
			return nil, errExpected
		}
		clearWasCalled := false
		bridgeStub.ClearStoredP2PSignaturesForPeerChainCalled = func() {
			clearWasCalled = true
		}

		step := resolveSetStatusStep{
			bridge: bridgeStub,
		}

		stepIdentifier := step.Execute(context.Background())
		assert.Equal(t, initialStep, stepIdentifier)
		assert.True(t, clearWasCalled)
	})
	t.Run("nil batch on GetBatchFromDharitrI", func(t *testing.T) {
		t.Parallel()
		bridgeStub := createStubExecutorResolveSetStatus()
		bridgeStub.GetBatchFromDharitrICalled = func(ctx context.Context) (*bridgeCore.TransferBatch, error) {
			return nil, nil
		}
		clearWasCalled := false
		bridgeStub.ClearStoredP2PSignaturesForPeerChainCalled = func() {
			clearWasCalled = true
		}

		step := resolveSetStatusStep{
			bridge: bridgeStub,
		}

		stepIdentifier := step.Execute(context.Background())
		assert.Equal(t, initialStep, stepIdentifier)
		assert.True(t, clearWasCalled)
	})
	t.Run("WaitAndReturnFinalBatchStatusesCalled returns nil, should go to GettingPendingBatchFromDharitrI", func(t *testing.T) {
		t.Parallel()

		bridgeStub := createStubExecutorResolveSetStatus()

		step := resolveSetStatusStep{
			bridge: bridgeStub,
		}

		stepIdentifier := step.Execute(context.Background())
		assert.Equal(t, initialStep, stepIdentifier)
	})
	t.Run("WaitAndReturnFinalBatchStatusesCalled returns empty slice, should go to GettingPendingBatchFromDharitrI", func(t *testing.T) {
		t.Parallel()

		bridgeStub := createStubExecutorResolveSetStatus()
		bridgeStub.WaitAndReturnFinalBatchStatusesCalled = func(ctx context.Context) []byte {
			return make([]byte, 0)
		}

		step := resolveSetStatusStep{
			bridge: bridgeStub,
		}

		stepIdentifier := step.Execute(context.Background())
		assert.Equal(t, initialStep, stepIdentifier)
	})
	t.Run("WaitAndReturnFinalBatchStatusesCalled should finish with success and go to ProposingSetStatusOnDharitrI", func(t *testing.T) {
		t.Parallel()
		bridgeStub := createStubExecutorResolveSetStatus()
		bridgeStub.WaitAndReturnFinalBatchStatusesCalled = func(ctx context.Context) []byte {
			return []byte{bridgeCore.Executed, bridgeCore.Rejected}
		}

		wasCalled := false
		bridgeStub.ResolveNewDepositsStatusesCalled = func(numDeposits uint64) {
			wasCalled = true
		}
		clearWasCalled := false
		bridgeStub.ClearStoredP2PSignaturesForPeerChainCalled = func() {
			clearWasCalled = true
		}

		step := resolveSetStatusStep{
			bridge: bridgeStub,
		}

		assert.False(t, step.IsInterfaceNil())

		expectedStep := bridgeCore.StepIdentifier(ProposingSetStatusOnDharitrI)
		stepIdentifier := step.Execute(context.Background())
		assert.True(t, wasCalled)
		assert.NotEqual(t, step.Identifier(), stepIdentifier)
		assert.Equal(t, expectedStep, stepIdentifier)
		assert.True(t, clearWasCalled)
	})
}

func createStubExecutorResolveSetStatus() *bridgeTests.BridgeExecutorStub {
	stub := bridgeTests.NewBridgeExecutorStub()
	stub.GetStoredBatchCalled = func() *bridgeCore.TransferBatch {
		return testBatch
	}
	stub.GetBatchFromDharitrICalled = func(ctx context.Context) (*bridgeCore.TransferBatch, error) {
		return testBatch, nil
	}
	return stub
}
