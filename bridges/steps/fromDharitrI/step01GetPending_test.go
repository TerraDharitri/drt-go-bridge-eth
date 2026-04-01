package dharitritoeth

import (
	"context"
	"errors"
	"math/big"
	"testing"

	bridgeCore "github.com/TerraDharitri/drt-go-bridge-eth/core"
	"github.com/TerraDharitri/drt-go-bridge-eth/core/batchProcessor"
	bridgeTests "github.com/TerraDharitri/drt-go-bridge-eth/testsCommon/bridge"
	"github.com/stretchr/testify/assert"
)

var errExpected = errors.New("expected error")//nolint:staticcheck
var testBatch = &bridgeCore.TransferBatch{
	ID:       112233,
	Deposits: nil,
	Statuses: nil,
}

func TestExecute_GetPending(t *testing.T) {
	t.Parallel()

	t.Run("error on GetBatchFromDharitrI", func(t *testing.T) {
		t.Parallel()
		bridgeStub := createStubExecutorGetPending()
		bridgeStub.GetBatchFromDharitrICalled = func(ctx context.Context) (*bridgeCore.TransferBatch, error) {
			return nil, errExpected
		}

		step := getPendingStep{
			bridge: bridgeStub,
		}

		expectedStepIdentifier := step.Identifier()
		stepIdentifier := step.Execute(context.Background())
		assert.Equal(t, expectedStepIdentifier, stepIdentifier)
	})
	t.Run("nil batch on GetBatchFromDharitrI", func(t *testing.T) {
		t.Parallel()
		bridgeStub := createStubExecutorGetPending()
		bridgeStub.GetBatchFromDharitrICalled = func(ctx context.Context) (*bridgeCore.TransferBatch, error) {
			return nil, nil
		}

		step := getPendingStep{
			bridge: bridgeStub,
		}

		expectedStepIdentifier := step.Identifier()
		stepIdentifier := step.Execute(context.Background())
		assert.Equal(t, expectedStepIdentifier, stepIdentifier)
	})
	t.Run("error on StoreBatchFromDharitrI", func(t *testing.T) {
		t.Parallel()
		bridgeStub := createStubExecutorGetPending()
		bridgeStub.StoreBatchFromDharitrICalled = func(batch *bridgeCore.TransferBatch) error {
			return errExpected
		}

		step := getPendingStep{
			bridge: bridgeStub,
		}

		expectedStepIdentifier := step.Identifier()
		stepIdentifier := step.Execute(context.Background())
		assert.Equal(t, expectedStepIdentifier, stepIdentifier)
	})
	t.Run("error on WasTransferPerformedOnPeerChain", func(t *testing.T) {
		t.Parallel()
		bridgeStub := createStubExecutorGetPending()
		bridgeStub.WasTransferPerformedOnPeerChainCalled = func(ctx context.Context) (bool, error) {
			return false, errExpected
		}

		step := getPendingStep{
			bridge: bridgeStub,
		}

		expectedStepIdentifier := step.Identifier()
		stepIdentifier := step.Execute(context.Background())
		assert.Equal(t, expectedStepIdentifier, stepIdentifier)
	})
	t.Run("error on WasTransferPerformedOnPeerChain", func(t *testing.T) {
		t.Parallel()
		bridgeStub := createStubExecutorGetPending()
		bridgeStub.WasTransferPerformedOnPeerChainCalled = func(ctx context.Context) (bool, error) {
			return false, nil
		}
		bridgeStub.CheckAvailableTokensCalled = func(ctx context.Context, peerTokens [][]byte, drtTokens [][]byte, amounts []*big.Int, direction batchProcessor.Direction) error {
			return errExpected
		}

		step := getPendingStep{
			bridge: bridgeStub,
		}

		expectedStepIdentifier := step.Identifier()
		stepIdentifier := step.Execute(context.Background())
		assert.Equal(t, expectedStepIdentifier, stepIdentifier)
	})
	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		t.Run("if transfer already performed next step should be ResolvingSetStatusOnDharitrI", func(t *testing.T) {
			t.Parallel()
			bridgeStub := createStubExecutorGetPending()
			bridgeStub.WasTransferPerformedOnPeerChainCalled = func(ctx context.Context) (bool, error) {
				return true, nil
			}
			checkAvailableTokensCalled := false
			bridgeStub.CheckAvailableTokensCalled = func(ctx context.Context, ethTokens [][]byte, drtTokens [][]byte, amounts []*big.Int, direction batchProcessor.Direction) error {
				checkAvailableTokensCalled = true
				return nil
			}

			step := getPendingStep{
				bridge: bridgeStub,
			}

			assert.False(t, step.IsInterfaceNil())

			expectedStepIdentifier := bridgeCore.StepIdentifier(ResolvingSetStatusOnDharitrI)
			stepIdentifier := step.Execute(context.Background())
			assert.Equal(t, expectedStepIdentifier, stepIdentifier)
			assert.False(t, checkAvailableTokensCalled)
		})
		t.Run("if transfer was not performed next step should be SigningProposedTransferOnPeerChain", func(t *testing.T) {
			t.Parallel()
			bridgeStub := createStubExecutorGetPending()
			bridgeStub.WasTransferPerformedOnPeerChainCalled = func(ctx context.Context) (bool, error) {
				return false, nil
			}
			checkAvailableTokensCalled := false
			bridgeStub.CheckAvailableTokensCalled = func(ctx context.Context, ethTokens [][]byte, drtTokens [][]byte, amounts []*big.Int, direction batchProcessor.Direction) error {
				checkAvailableTokensCalled = true
				return nil
			}

			step := getPendingStep{
				bridge: bridgeStub,
			}

			expectedStepIdentifier := bridgeCore.StepIdentifier(SigningProposedTransferOnPeerChain)
			stepIdentifier := step.Execute(context.Background())
			assert.Equal(t, expectedStepIdentifier, stepIdentifier)
			assert.True(t, checkAvailableTokensCalled)
		})
	})
}

func createStubExecutorGetPending() *bridgeTests.BridgeExecutorStub {
	stub := bridgeTests.NewBridgeExecutorStub()
	stub.GetBatchFromDharitrICalled = func(ctx context.Context) (*bridgeCore.TransferBatch, error) {
		return testBatch, nil
	}
	stub.StoreBatchFromDharitrICalled = func(batch *bridgeCore.TransferBatch) error {
		return nil
	}
	return stub
}
