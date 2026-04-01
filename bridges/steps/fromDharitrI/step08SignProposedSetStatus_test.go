package dharitritoeth

import (
	"context"
	"github.com/TerraDharitri/drt-go-bridge-eth/bridges"
	"testing"

	bridgeCore "github.com/TerraDharitri/drt-go-bridge-eth/core"
	bridgeTests "github.com/TerraDharitri/drt-go-bridge-eth/testsCommon/bridge"
	"github.com/stretchr/testify/assert"
)

var actionID = uint64(662528)

func TestExecute_SignProposedSetStatus(t *testing.T) {
	t.Parallel()
	t.Run("nil batch on GetStoredBatch", func(t *testing.T) {
		t.Parallel()
		bridgeStub := createStubExecutorSignProposedSetStatus()
		bridgeStub.GetStoredBatchCalled = func() *bridgeCore.TransferBatch {
			return nil
		}

		step := signProposedSetStatusStep{
			bridge: bridgeStub,
		}

		stepIdentifier := step.Execute(context.Background())
		assert.Equal(t, initialStep, stepIdentifier)
	})
	t.Run("error on GetAndStoreActionIDForProposeSetStatusFromDharitrI", func(t *testing.T) {
		t.Parallel()
		bridgeStub := createStubExecutorSignProposedSetStatus()
		bridgeStub.GetAndStoreActionIDForProposeSetStatusFromDharitrICalled = func(ctx context.Context) (uint64, error) {
			return bridges.InvalidActionID, errExpected
		}

		step := signProposedSetStatusStep{
			bridge: bridgeStub,
		}

		stepIdentifier := step.Execute(context.Background())
		assert.Equal(t, initialStep, stepIdentifier)
	})
	t.Run("invalid actionID on GetAndStoreActionIDForProposeSetStatusFromDharitrI", func(t *testing.T) {
		t.Parallel()
		bridgeStub := createStubExecutorSignProposedSetStatus()
		bridgeStub.GetAndStoreActionIDForProposeSetStatusFromDharitrICalled = func(ctx context.Context) (uint64, error) {
			return bridges.InvalidActionID, nil
		}

		step := signProposedSetStatusStep{
			bridge: bridgeStub,
		}

		stepIdentifier := step.Execute(context.Background())
		assert.Equal(t, initialStep, stepIdentifier)
	})
	t.Run("error on WasActionSignedOnDharitrI", func(t *testing.T) {
		t.Parallel()
		bridgeStub := createStubExecutorSignProposedSetStatus()
		bridgeStub.WasActionSignedOnDharitrICalled = func(ctx context.Context) (bool, error) {
			return false, errExpected
		}

		step := signProposedSetStatusStep{
			bridge: bridgeStub,
		}

		stepIdentifier := step.Execute(context.Background())
		assert.Equal(t, initialStep, stepIdentifier)
	})
	t.Run("error on SignActionOnDharitrI", func(t *testing.T) {
		t.Parallel()
		bridgeStub := createStubExecutorSignProposedSetStatus()
		bridgeStub.SignActionOnDharitrICalled = func(ctx context.Context) error {
			return errExpected
		}

		step := signProposedSetStatusStep{
			bridge: bridgeStub,
		}

		stepIdentifier := step.Execute(context.Background())
		assert.Equal(t, initialStep, stepIdentifier)
	})
	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		t.Run("if proposed set status was signed, go to WaitingForQuorumOnSetStatus", func(t *testing.T) {
			t.Parallel()
			bridgeStub := createStubExecutorSignProposedSetStatus()
			bridgeStub.WasActionSignedOnDharitrICalled = func(ctx context.Context) (bool, error) {
				return true, nil
			}

			wasCalled := false
			bridgeStub.SignActionOnDharitrICalled = func(ctx context.Context) error {
				wasCalled = true
				return nil
			}

			step := signProposedSetStatusStep{
				bridge: bridgeStub,
			}

			expectedStep := bridgeCore.StepIdentifier(WaitingForQuorumOnSetStatus)
			stepIdentifier := step.Execute(context.Background())
			assert.False(t, wasCalled)
			assert.Equal(t, expectedStep, stepIdentifier)
		})
		t.Run("if proposed set status was not signed, sign and go to WaitingForQuorumOnSetStatus", func(t *testing.T) {
			t.Parallel()
			bridgeStub := createStubExecutorSignProposedSetStatus()
			wasCalled := false
			bridgeStub.SignActionOnDharitrICalled = func(ctx context.Context) error {
				wasCalled = true
				return nil
			}

			step := signProposedSetStatusStep{
				bridge: bridgeStub,
			}

			assert.False(t, step.IsInterfaceNil())
			expectedStep := bridgeCore.StepIdentifier(WaitingForQuorumOnSetStatus)
			stepIdentifier := step.Execute(context.Background())
			assert.True(t, wasCalled)
			assert.NotEqual(t, step.Identifier(), stepIdentifier)
			assert.Equal(t, expectedStep, stepIdentifier)
		})
	})

}

func createStubExecutorSignProposedSetStatus() *bridgeTests.BridgeExecutorStub {
	stub := bridgeTests.NewBridgeExecutorStub()
	stub.GetStoredBatchCalled = func() *bridgeCore.TransferBatch {
		return testBatch
	}
	stub.GetAndStoreActionIDForProposeSetStatusFromDharitrICalled = func(ctx context.Context) (uint64, error) {
		return actionID, nil
	}
	stub.WasActionSignedOnDharitrICalled = func(ctx context.Context) (bool, error) {
		return false, nil
	}
	stub.SignActionOnDharitrICalled = func(ctx context.Context) error {
		return nil
	}
	return stub
}
