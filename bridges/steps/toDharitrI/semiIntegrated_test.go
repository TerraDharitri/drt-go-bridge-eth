package ethtodharitri

import (
	"context"
	"errors"
	"fmt"
	"github.com/TerraDharitri/drt-go-bridge-eth/bridges/steps"
	"testing"

	bridgeCore "github.com/TerraDharitri/drt-go-bridge-eth/core"
	bridgeTests "github.com/TerraDharitri/drt-go-bridge-eth/testsCommon/bridge"
	"github.com/TerraDharitri/drt-go-bridge-eth/testsCommon/stateMachine"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	getAndStoreBatchFromPeerChain                     = "GetAndStoreBatchFromPeerChain"
	getLastExecutedPeerBatchIDFromDharitrI          = "GetLastExecutedPeerBatchIDFromDharitrI"
	verifyLastDepositNonceExecutedOnPeerBatch         = "VerifyLastDepositNonceExecutedOnPeerBatch"
	wasTransferProposedOnDharitrI                   = "WasTransferProposedOnDharitrI"
	wasActionSignedOnDharitrI                       = "WasActionSignedOnDharitrI"
	signActionOnDharitrI                            = "SignActionOnDharitrI"
	getAndStoreActionIDForProposeTransferOnDharitrI = "GetAndStoreActionIDForProposeTransferOnDharitrI"
	ProcessMaxQuorumRetriesOnDharitrI               = "ProcessMaxQuorumRetriesOnDharitrI"
	resetRetriesCountOnDharitrI                     = "ResetRetriesCountOnDharitrI"
	processQuorumReachedOnDharitrI                  = "ProcessQuorumReachedOnDharitrI"
	wasActionPerformedOnDharitrI                    = "WasActionPerformedOnDharitrI"
	proposeTransferOnDharitrI                       = "ProposeTransferOnDharitrI"
	performActionOnDharitrI                         = "PerformActionOnDharitrI"
)

var trueHandler = func() bool { return true }
var falseHandler = func() bool { return false }

type errorHandler struct {
	lastError error
}

func (eh *errorHandler) storeAndReturnError(err error) error {
	eh.lastError = err
	return err
}

type argsBridgeStub struct {
	failingStep                      string
	myTurnHandler                    func() bool
	wasTransferProposedHandler       func() bool
	wasProposedTransferSignedHandler func() bool
	wasActionSigned                  func() bool
	isQuorumReachedHandler           func() bool
	wasActionIDPerformedHandler      func() bool
	maxRetriesReachedHandler         func() bool
	validateBatchHandler             func() bool
}

func createMockBridge(args argsBridgeStub) (*bridgeTests.BridgeExecutorStub, *errorHandler) {
	errHandler := &errorHandler{}
	stub := bridgeTests.NewBridgeExecutorStub()
	expectedErr := errors.New("expected error")
	stub.MyTurnAsLeaderCalled = func() bool {
		return args.myTurnHandler()
	}
	stub.GetAndStoreActionIDForProposeTransferOnDharitrICalled = func(ctx context.Context) (uint64, error) {
		if args.failingStep == getAndStoreActionIDForProposeTransferOnDharitrI {
			return 0, errHandler.storeAndReturnError(expectedErr)
		}

		return 2, errHandler.storeAndReturnError(nil)
	}
	stub.GetStoredActionIDCalled = func() uint64 {
		return 2
	}
	stub.GetAndStoreBatchFromPeerChainCalled = func(ctx context.Context, nonce uint64) error {
		if args.failingStep == getAndStoreBatchFromPeerChain {
			return errHandler.storeAndReturnError(expectedErr)
		}

		return errHandler.storeAndReturnError(nil)
	}
	stub.GetStoredBatchCalled = func() *bridgeCore.TransferBatch {
		return &bridgeCore.TransferBatch{}
	}
	stub.GetLastExecutedPeerBatchIDFromDharitrICalled = func(ctx context.Context) (uint64, error) {
		if args.failingStep == getLastExecutedPeerBatchIDFromDharitrI {
			return 0, errHandler.storeAndReturnError(expectedErr)
		}

		return 3, errHandler.storeAndReturnError(nil)
	}
	stub.VerifyLastDepositNonceExecutedOnPeerBatchCalled = func(ctx context.Context) error {
		if args.failingStep == verifyLastDepositNonceExecutedOnPeerBatch {
			return errHandler.storeAndReturnError(expectedErr)
		}

		return errHandler.storeAndReturnError(nil)
	}
	stub.WasTransferProposedOnDharitrICalled = func(ctx context.Context) (bool, error) {
		if args.failingStep == wasTransferProposedOnDharitrI {
			return false, errHandler.storeAndReturnError(expectedErr)
		}

		return args.wasTransferProposedHandler(), errHandler.storeAndReturnError(nil)
	}
	stub.ProposeTransferOnDharitrICalled = func(ctx context.Context) error {
		if args.failingStep == proposeTransferOnDharitrI {
			return errHandler.storeAndReturnError(expectedErr)
		}

		return errHandler.storeAndReturnError(nil)
	}
	stub.WasActionSignedOnDharitrICalled = func(ctx context.Context) (bool, error) {
		if args.failingStep == wasActionSignedOnDharitrI {
			return false, errHandler.storeAndReturnError(expectedErr)
		}

		return args.wasActionSigned(), errHandler.storeAndReturnError(nil)
	}
	stub.SignActionOnDharitrICalled = func(ctx context.Context) error {
		if args.failingStep == signActionOnDharitrI {
			return errHandler.storeAndReturnError(expectedErr)
		}

		return errHandler.storeAndReturnError(nil)
	}
	stub.ProcessQuorumReachedOnDharitrICalled = func(ctx context.Context) (bool, error) {
		if args.failingStep == processQuorumReachedOnDharitrI {
			return false, errHandler.storeAndReturnError(expectedErr)
		}

		return args.isQuorumReachedHandler(), errHandler.storeAndReturnError(nil)
	}
	stub.WasActionPerformedOnDharitrICalled = func(ctx context.Context) (bool, error) {
		if args.failingStep == wasActionPerformedOnDharitrI {
			return false, errHandler.storeAndReturnError(expectedErr)
		}

		return args.wasActionIDPerformedHandler(), errHandler.storeAndReturnError(nil)
	}
	stub.PerformActionOnDharitrICalled = func(ctx context.Context) error {
		if args.failingStep == performActionOnDharitrI {
			return errHandler.storeAndReturnError(expectedErr)
		}

		return errHandler.storeAndReturnError(nil)
	}
	stub.ProcessMaxQuorumRetriesOnDharitrICalled = func() bool {
		return args.maxRetriesReachedHandler()
	}

	return stub, errHandler
}

func createStateMachine(t *testing.T, executor steps.Executor, initialStep bridgeCore.StepIdentifier) *stateMachine.StateMachineMock {
	stepsSlice, err := CreateSteps(executor)
	require.Nil(t, err)

	sm := stateMachine.NewStateMachineMock(stepsSlice, initialStep)
	err = sm.Initialize()
	require.Nil(t, err)

	return sm
}

func TestHappyCaseWhenLeader(t *testing.T) {
	t.Parallel()

	args := argsBridgeStub{
		myTurnHandler:                    trueHandler,
		isQuorumReachedHandler:           trueHandler,
		wasActionIDPerformedHandler:      trueHandler,
		validateBatchHandler:             trueHandler,
		maxRetriesReachedHandler:         falseHandler,
		wasProposedTransferSignedHandler: falseHandler,
		wasTransferProposedHandler:       falseHandler,
		wasActionSigned:                  falseHandler,
	}
	executor, eh := createMockBridge(args)
	sm := createStateMachine(t, executor, GettingPendingBatchFromPeerChain)
	numSteps := 20
	for i := 0; i < numSteps; i++ {
		err := sm.Execute(context.Background())
		require.Nil(t, err)
	}

	assert.Equal(t, 4, executor.GetFunctionCounter(resetRetriesCountOnDharitrI))
	assert.Equal(t, 4, executor.GetFunctionCounter(getLastExecutedPeerBatchIDFromDharitrI))
	assert.Equal(t, 4, executor.GetFunctionCounter(getAndStoreBatchFromPeerChain))
	assert.Equal(t, 4, executor.GetFunctionCounter(verifyLastDepositNonceExecutedOnPeerBatch))

	assert.Equal(t, 4, executor.GetFunctionCounter(wasTransferProposedOnDharitrI))
	assert.Equal(t, 4, executor.GetFunctionCounter(proposeTransferOnDharitrI))

	assert.Equal(t, 4, executor.GetFunctionCounter(getAndStoreActionIDForProposeTransferOnDharitrI))
	assert.Equal(t, 4, executor.GetFunctionCounter(wasActionSignedOnDharitrI))
	assert.Equal(t, 4, executor.GetFunctionCounter(signActionOnDharitrI))

	assert.Equal(t, 4, executor.GetFunctionCounter(ProcessMaxQuorumRetriesOnDharitrI))
	assert.Equal(t, 4, executor.GetFunctionCounter(processQuorumReachedOnDharitrI))

	assert.Equal(t, 4, executor.GetFunctionCounter(wasActionPerformedOnDharitrI))
	assert.Equal(t, 0, executor.GetFunctionCounter(performActionOnDharitrI))

	assert.Nil(t, eh.lastError)
}

func TestHappyCaseWhenLeaderAndActionIdNotPerformed(t *testing.T) {
	t.Parallel()

	numCalled := 0
	args := argsBridgeStub{
		myTurnHandler:          trueHandler,
		isQuorumReachedHandler: trueHandler,
		validateBatchHandler:   trueHandler,
		wasActionIDPerformedHandler: func() bool {
			numCalled++
			return numCalled > 1
		},
		maxRetriesReachedHandler:         falseHandler,
		wasProposedTransferSignedHandler: falseHandler,
		wasTransferProposedHandler:       falseHandler,
		wasActionSigned:                  falseHandler,
	}
	executor, eh := createMockBridge(args)
	sm := createStateMachine(t, executor, GettingPendingBatchFromPeerChain)
	numSteps := 20
	for i := 0; i < numSteps; i++ {
		err := sm.Execute(context.Background())
		require.Nil(t, err)
	}

	assert.Equal(t, 4, executor.GetFunctionCounter(resetRetriesCountOnDharitrI))
	assert.Equal(t, 4, executor.GetFunctionCounter(getLastExecutedPeerBatchIDFromDharitrI))
	assert.Equal(t, 4, executor.GetFunctionCounter(getAndStoreBatchFromPeerChain))
	assert.Equal(t, 4, executor.GetFunctionCounter(verifyLastDepositNonceExecutedOnPeerBatch))

	assert.Equal(t, 4, executor.GetFunctionCounter(wasTransferProposedOnDharitrI))
	assert.Equal(t, 4, executor.GetFunctionCounter(proposeTransferOnDharitrI))

	assert.Equal(t, 4, executor.GetFunctionCounter(getAndStoreActionIDForProposeTransferOnDharitrI))
	assert.Equal(t, 4, executor.GetFunctionCounter(wasActionSignedOnDharitrI))
	assert.Equal(t, 4, executor.GetFunctionCounter(signActionOnDharitrI))

	assert.Equal(t, 4, executor.GetFunctionCounter(ProcessMaxQuorumRetriesOnDharitrI))
	assert.Equal(t, 4, executor.GetFunctionCounter(processQuorumReachedOnDharitrI))

	assert.Equal(t, 4, executor.GetFunctionCounter(wasActionPerformedOnDharitrI))
	assert.Equal(t, 1, executor.GetFunctionCounter(performActionOnDharitrI))

	assert.Nil(t, eh.lastError)
}

func TestOneStepErrors_ShouldReturnToPendingBatch(t *testing.T) {
	stepsThatCanError := []bridgeCore.StepIdentifier{
		getAndStoreActionIDForProposeTransferOnDharitrI,
		getAndStoreBatchFromPeerChain,
		getLastExecutedPeerBatchIDFromDharitrI,
		verifyLastDepositNonceExecutedOnPeerBatch,
		wasTransferProposedOnDharitrI,
		proposeTransferOnDharitrI,
		wasTransferProposedOnDharitrI,
		signActionOnDharitrI,
		processQuorumReachedOnDharitrI,
		wasActionPerformedOnDharitrI,
		performActionOnDharitrI,
	}

	for _, stepThatError := range stepsThatCanError {
		testErrorFlow(t, stepThatError)
	}
}

func testErrorFlow(t *testing.T, stepThatErrors bridgeCore.StepIdentifier) {
	numCalled := 0
	args := argsBridgeStub{
		failingStep:            string(stepThatErrors),
		myTurnHandler:          trueHandler,
		isQuorumReachedHandler: trueHandler,
		validateBatchHandler:   trueHandler,
		wasActionIDPerformedHandler: func() bool {
			numCalled++
			return numCalled > 1
		},
		maxRetriesReachedHandler:         falseHandler,
		wasProposedTransferSignedHandler: falseHandler,
		wasTransferProposedHandler:       falseHandler,
		wasActionSigned:                  falseHandler,
	}

	executor, eh := createMockBridge(args)
	sm := createStateMachine(t, executor, GettingPendingBatchFromPeerChain)

	maxNumSteps := 10
	for i := 0; i < maxNumSteps; i++ {
		err := sm.Execute(context.Background())
		assert.Nil(t, err)

		if eh.lastError != nil {
			if sm.CurrentStep.Identifier() == GettingPendingBatchFromPeerChain {
				return
			}

			require.Fail(t, fmt.Sprintf("should have jumped to initial step, got next step %s, stepThatErrors %s",
				sm.CurrentStep.Identifier(), stepThatErrors))
		}
	}

	require.Fail(t, fmt.Sprintf("max number of steps reached but not jumped to initial step, stepThatErrors %s", stepThatErrors))
}
