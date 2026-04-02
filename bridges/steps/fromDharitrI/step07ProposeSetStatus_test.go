package dharitritoeth

import (
	"context"
	"testing"

	bridgeCore "github.com/TerraDharitri/drt-go-bridge-eth/core"
	bridgeTests "github.com/TerraDharitri/drt-go-bridge-eth/testsCommon/bridge"
	"github.com/stretchr/testify/assert"
)

func TestExecute_ProposeSetStatus(t *testing.T) {
	t.Parallel()
	t.Run("nil batch on GetStoredBatch", func(t *testing.T) {
		t.Parallel()
		bridgeStub := createStubExecutorProposeSetStatus()
		bridgeStub.GetStoredBatchCalled = func() *bridgeCore.TransferBatch {
			return nil
		}

		step := proposeSetStatusStep{
			bridge: bridgeStub,
		}

		stepIdentifier := step.Execute(context.Background())
		assert.Equal(t, initialStep, stepIdentifier)
	})

	t.Run("max retries reached", func(t *testing.T) {
		t.Parallel()
		bridgeStub := createStubExecutorProposeSetStatus()
		bridgeStub.ProcessMaxRetriesOnWasTransferProposedOnDharitrICalled = func() bool {
			return true
		}

		step := proposeSetStatusStep{
			bridge: bridgeStub,
		}

		stepIdentifier := step.Execute(context.Background())
		assert.Equal(t, initialStep, stepIdentifier)
	})

	t.Run("error on WasSetStatusProposedOnDharitrI", func(t *testing.T) {
		t.Parallel()
		bridgeStub := createStubExecutorProposeSetStatus()
		bridgeStub.WasSetStatusProposedOnDharitrICalled = func(ctx context.Context) (bool, error) {
			return false, errExpected
		}

		step := proposeSetStatusStep{
			bridge: bridgeStub,
		}

		stepIdentifier := step.Execute(context.Background())
		assert.Equal(t, initialStep, stepIdentifier)
	})

	t.Run("error on ProposeSetStatusOnDharitrI", func(t *testing.T) {
		t.Parallel()
		bridgeStub := createStubExecutorProposeSetStatus()
		bridgeStub.ProposeSetStatusOnDharitrICalled = func(ctx context.Context) error {
			return errExpected
		}

		step := proposeSetStatusStep{
			bridge: bridgeStub,
		}

		stepIdentifier := step.Execute(context.Background())
		assert.Equal(t, initialStep, stepIdentifier)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		t.Run("if SetStatus was proposed it should go to SigningProposedSetStatusOnDharitrI", func(t *testing.T) {
			t.Parallel()
			bridgeStub := createStubExecutorProposeSetStatus()
			bridgeStub.WasSetStatusProposedOnDharitrICalled = func(ctx context.Context) (bool, error) {
				return true, nil
			}

			step := proposeSetStatusStep{
				bridge: bridgeStub,
			}

			assert.False(t, step.IsInterfaceNil())
			expectedStep := bridgeCore.StepIdentifier(SigningProposedSetStatusOnDharitrI)
			stepIdentifier := step.Execute(context.Background())
			assert.Equal(t, expectedStep, stepIdentifier)

		})
		t.Run("if SetStatus was not proposed", func(t *testing.T) {
			t.Parallel()
			t.Run("if not leader, should stay in current step", func(t *testing.T) {
				t.Parallel()
				bridgeStub := createStubExecutorProposeSetStatus()
				bridgeStub.MyTurnAsLeaderCalled = func() bool {
					return false
				}
				step := proposeSetStatusStep{
					bridge: bridgeStub,
				}

				stepIdentifier := step.Execute(context.Background())
				assert.Equal(t, step.Identifier(), stepIdentifier)

			})
			t.Run("if leader, should go to SigningProposedTransferOnDharitrI", func(t *testing.T) {
				t.Parallel()
				bridgeStub := createStubExecutorProposeSetStatus()

				step := proposeSetStatusStep{
					bridge: bridgeStub,
				}

				expectedStep := bridgeCore.StepIdentifier(SigningProposedSetStatusOnDharitrI)
				stepIdentifier := step.Execute(context.Background())
				assert.Equal(t, expectedStep, stepIdentifier)

			})
		})

	})
}

func createStubExecutorProposeSetStatus() *bridgeTests.BridgeExecutorStub {
	stub := bridgeTests.NewBridgeExecutorStub()
	stub.GetStoredBatchCalled = func() *bridgeCore.TransferBatch {
		return testBatch
	}
	stub.WasSetStatusProposedOnDharitrICalled = func(ctx context.Context) (bool, error) {
		return false, nil
	}
	stub.MyTurnAsLeaderCalled = func() bool {
		return true
	}
	stub.ProposeSetStatusOnDharitrICalled = func(ctx context.Context) error {
		return nil
	}
	return stub
}
