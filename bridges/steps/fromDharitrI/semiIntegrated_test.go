package dharitritoeth

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
	getBatchFromDharitrI                               = "GetBatchFromDharitrI"
	storeBatchFromDharitrI                             = "StoreBatchFromDharitrI"
	wasTransferPerformedOnPeerChain                      = "WasTransferPerformedOnPeerChain"
	signTransferOnPeerChain                              = "SignTransferOnPeerChain"
	ProcessMaxQuorumRetriesOnPeerChain                   = "ProcessMaxQuorumRetriesOnPeerChain"
	processQuorumReachedOnPeerChain                      = "ProcessQuorumReachedOnPeerChain"
	performTransferOnPeerChain                           = "PerformTransferOnPeerChain"
	getBatchStatusesFromPeerChain                        = "GetBatchStatusesFromPeerChain"
	wasSetStatusProposedOnDharitrI                     = "WasSetStatusProposedOnDharitrI"
	proposeSetStatusOnDharitrI                         = "ProposeSetStatusOnDharitrI"
	getAndStoreActionIDForProposeSetStatusFromDharitrI = "GetAndStoreActionIDForProposeSetStatusFromDharitrI"
	wasActionSignedOnDharitrI                          = "WasActionSignedOnDharitrI"
	signActionOnDharitrI                               = "SignActionOnDharitrI"
	ProcessMaxQuorumRetriesOnDharitrI                  = "ProcessMaxQuorumRetriesOnDharitrI"
	processQuorumReachedOnDharitrI                     = "ProcessQuorumReachedOnDharitrI"
	wasActionPerformedOnDharitrI                       = "WasActionPerformedOnDharitrI"
	performActionOnDharitrI                            = "PerformActionOnDharitrI"
	resetRetriesCountOnPeerChain                         = "ResetRetriesCountOnPeerChain"
	resetRetriesCountOnDharitrI                        = "ResetRetriesCountOnDharitrI"
	getStoredBatch                                       = "GetStoredBatch"
	myTurnAsLeader                                       = "MyTurnAsLeader"
	waitForTransferConfirmation                          = "WaitForTransferConfirmation"
	WaitAndReturnFinalBatchStatuses                      = "WaitAndReturnFinalBatchStatuses"
	resolveNewDepositsStatuses                           = "ResolveNewDepositsStatuses"
	getStoredActionID                                    = "GetStoredActionID"
)

type argsBridgeStub struct {
	failingStep                             string
	wasTransferPerformedOnPeerChainHandler  func() bool
	processQuorumReachedOnPeerChainHandler  func() bool
	processQuorumReachedOnDharitrIHandler func() bool
	myTurnHandler                           func() bool
	wasSetStatusProposedOnDharitrIHandler func() bool
	wasActionSignedOnDharitrIHandler      func() bool
	wasActionPerformedOnDharitrIHandler   func() bool
	maxRetriesReachedPeerChainHandler       func() bool
	maxRetriesReachedDharitrIHandler      func() bool
}

var trueHandler = func() bool { return true }
var falseHandler = func() bool { return false }

type errorHandler struct {
	lastError error
}

func (eh *errorHandler) storeAndReturnError(err error) error {
	eh.lastError = err
	return err
}

func createStateMachine(t *testing.T, executor steps.Executor, initialStep bridgeCore.StepIdentifier) *stateMachine.StateMachineMock {
	stepsSlice, err := CreateSteps(executor)
	require.Nil(t, err)

	sm := stateMachine.NewStateMachineMock(stepsSlice, initialStep)
	err = sm.Initialize()
	require.Nil(t, err)

	return sm
}

func createMockBridge(args argsBridgeStub) (*bridgeTests.BridgeExecutorStub, *errorHandler) {
	errHandler := &errorHandler{}
	stub := bridgeTests.NewBridgeExecutorStub()
	expectedErr := errors.New("expected error")
	stub.MyTurnAsLeaderCalled = func() bool {
		return args.myTurnHandler()
	}
	stub.GetAndStoreActionIDForProposeSetStatusFromDharitrICalled = func(ctx context.Context) (uint64, error) {
		if args.failingStep == getAndStoreActionIDForProposeSetStatusFromDharitrI {
			return 0, errHandler.storeAndReturnError(expectedErr)
		}

		return 2, errHandler.storeAndReturnError(nil)
	}
	stub.GetStoredActionIDCalled = func() uint64 {
		return 2
	}
	stub.GetBatchFromDharitrICalled = func(ctx context.Context) (*bridgeCore.TransferBatch, error) {
		if args.failingStep == getBatchFromDharitrI {
			return &bridgeCore.TransferBatch{}, errHandler.storeAndReturnError(expectedErr)
		}
		return &bridgeCore.TransferBatch{}, errHandler.storeAndReturnError(nil)
	}
	stub.StoreBatchFromDharitrICalled = func(batch *bridgeCore.TransferBatch) error {
		return nil
	}
	stub.GetStoredBatchCalled = func() *bridgeCore.TransferBatch {
		return &bridgeCore.TransferBatch{}
	}
	stub.WasTransferPerformedOnPeerChainCalled = func(ctx context.Context) (bool, error) {
		if args.failingStep == wasTransferPerformedOnPeerChain {
			return false, errHandler.storeAndReturnError(expectedErr)
		}

		return args.wasTransferPerformedOnPeerChainHandler(), errHandler.storeAndReturnError(nil)
	}
	stub.SignTransferOnPeerChainCalled = func() error {
		if args.failingStep == signTransferOnPeerChain {
			return errHandler.storeAndReturnError(expectedErr)
		}

		return errHandler.storeAndReturnError(nil)
	}
	stub.ProcessQuorumReachedOnPeerChainCalled = func(ctx context.Context) (bool, error) {
		if args.failingStep == processQuorumReachedOnPeerChain {
			return false, errHandler.storeAndReturnError(expectedErr)
		}

		return args.processQuorumReachedOnPeerChainHandler(), errHandler.storeAndReturnError(nil)
	}
	stub.PerformTransferOnPeerChainCalled = func(ctx context.Context) error {
		if args.failingStep == performTransferOnPeerChain {
			return errHandler.storeAndReturnError(expectedErr)
		}
		return errHandler.storeAndReturnError(nil)
	}
	stub.WaitForTransferConfirmationCalled = func(ctx context.Context) {
		stub.WasTransferPerformedOnPeerChainCalled = func(ctx context.Context) (bool, error) {
			return true, errHandler.storeAndReturnError(nil)
		}
	}
	stub.WaitAndReturnFinalBatchStatusesCalled = func(ctx context.Context) []byte {
		if args.failingStep == getBatchStatusesFromPeerChain {
			return nil
		}
		return []byte{0x3}
	}
	stub.GetBatchStatusesFromPeerChainCalled = func(ctx context.Context) ([]byte, error) {
		if args.failingStep == getBatchStatusesFromPeerChain {
			return nil, errHandler.storeAndReturnError(expectedErr)
		}
		return []byte{}, errHandler.storeAndReturnError(nil)
	}
	stub.ResolveNewDepositsStatusesCalled = func(numDeposits uint64) {

	}
	stub.WasSetStatusProposedOnDharitrICalled = func(ctx context.Context) (bool, error) {
		if args.failingStep == wasSetStatusProposedOnDharitrI {
			return false, errHandler.storeAndReturnError(expectedErr)
		}
		return args.wasSetStatusProposedOnDharitrIHandler(), errHandler.storeAndReturnError(nil)
	}
	stub.ProposeSetStatusOnDharitrICalled = func(ctx context.Context) error {
		if args.failingStep == proposeSetStatusOnDharitrI {
			return errHandler.storeAndReturnError(expectedErr)
		}

		return errHandler.storeAndReturnError(nil)
	}
	stub.WasActionSignedOnDharitrICalled = func(ctx context.Context) (bool, error) {
		if args.failingStep == wasActionSignedOnDharitrI {
			return false, errHandler.storeAndReturnError(expectedErr)
		}

		return args.wasActionSignedOnDharitrIHandler(), errHandler.storeAndReturnError(nil)
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

		return args.processQuorumReachedOnDharitrIHandler(), errHandler.storeAndReturnError(nil)
	}
	stub.WasActionPerformedOnDharitrICalled = func(ctx context.Context) (bool, error) {
		if args.failingStep == wasActionPerformedOnDharitrI {
			return false, errHandler.storeAndReturnError(expectedErr)
		}

		return args.wasActionPerformedOnDharitrIHandler(), errHandler.storeAndReturnError(nil)
	}
	stub.PerformActionOnDharitrICalled = func(ctx context.Context) error {
		if args.failingStep == performActionOnDharitrI {
			return errHandler.storeAndReturnError(expectedErr)
		}

		return errHandler.storeAndReturnError(nil)
	}
	stub.ProcessMaxQuorumRetriesOnDharitrICalled = func() bool {
		return args.maxRetriesReachedPeerChainHandler()
	}
	stub.ProcessMaxQuorumRetriesOnPeerChainCalled = func() bool {
		return args.maxRetriesReachedDharitrIHandler()
	}

	return stub, errHandler
}

func TestHappyCaseWhenLeaderSetStatusAlreadySigned(t *testing.T) {
	t.Parallel()

	numCalled := 0
	args := argsBridgeStub{
		myTurnHandler:                           trueHandler,
		processQuorumReachedOnPeerChainHandler:  trueHandler,
		processQuorumReachedOnDharitrIHandler: trueHandler,
		wasActionSignedOnDharitrIHandler:      trueHandler,
		wasActionPerformedOnDharitrIHandler: func() bool {
			numCalled++
			return numCalled > 1
		},
		wasTransferPerformedOnPeerChainHandler:  falseHandler,
		maxRetriesReachedPeerChainHandler:       falseHandler,
		maxRetriesReachedDharitrIHandler:      falseHandler,
		wasSetStatusProposedOnDharitrIHandler: falseHandler,
	}
	executor, eh := createMockBridge(args)
	sm := createStateMachine(t, executor, GettingPendingBatchFromDharitrI)
	numSteps := 12
	for i := 0; i < numSteps; i++ {
		err := sm.Execute(context.Background())
		require.Nil(t, err)
	}

	assert.Equal(t, 1, executor.GetFunctionCounter(resetRetriesCountOnPeerChain))
	assert.Equal(t, 1, executor.GetFunctionCounter(resetRetriesCountOnDharitrI))
	assert.Equal(t, 2, executor.GetFunctionCounter(getBatchFromDharitrI))
	assert.Equal(t, 1, executor.GetFunctionCounter(storeBatchFromDharitrI))
	assert.Equal(t, 3, executor.GetFunctionCounter(wasTransferPerformedOnPeerChain))
	assert.Equal(t, 4, executor.GetFunctionCounter(getStoredBatch))
	assert.Equal(t, 1, executor.GetFunctionCounter(signTransferOnPeerChain))
	assert.Equal(t, 3, executor.GetFunctionCounter(wasTransferPerformedOnPeerChain))
	assert.Equal(t, 1, executor.GetFunctionCounter(ProcessMaxQuorumRetriesOnPeerChain))
	assert.Equal(t, 1, executor.GetFunctionCounter(processQuorumReachedOnPeerChain))
	assert.Equal(t, 3, executor.GetFunctionCounter(myTurnAsLeader))
	assert.Equal(t, 1, executor.GetFunctionCounter(ProcessMaxQuorumRetriesOnDharitrI))
	assert.Equal(t, 1, executor.GetFunctionCounter(processQuorumReachedOnDharitrI))
	assert.Equal(t, 1, executor.GetFunctionCounter(waitForTransferConfirmation))
	assert.Equal(t, 1, executor.GetFunctionCounter(resolveNewDepositsStatuses))
	assert.Equal(t, 1, executor.GetFunctionCounter(wasSetStatusProposedOnDharitrI))
	assert.Equal(t, 1, executor.GetFunctionCounter(performTransferOnPeerChain))
	assert.Equal(t, 1, executor.GetFunctionCounter(WaitAndReturnFinalBatchStatuses))
	assert.Equal(t, 1, executor.GetFunctionCounter(proposeSetStatusOnDharitrI))
	assert.Equal(t, 1, executor.GetFunctionCounter(getAndStoreActionIDForProposeSetStatusFromDharitrI))
	assert.Equal(t, 2, executor.GetFunctionCounter(wasActionPerformedOnDharitrI))
	assert.Equal(t, 1, executor.GetFunctionCounter(performActionOnDharitrI))

	assert.Equal(t, 1, executor.GetFunctionCounter(wasActionSignedOnDharitrI))
	assert.Equal(t, 1, executor.GetFunctionCounter(getStoredActionID))

	assert.Nil(t, eh.lastError)
}

func TestOneStepErrors_ShouldReturnToPendingBatch(t *testing.T) {
	stepsThatCanError := []bridgeCore.StepIdentifier{
		getBatchFromDharitrI,
		wasTransferPerformedOnPeerChain,
		signTransferOnPeerChain,
		processQuorumReachedOnPeerChain,
		performTransferOnPeerChain,
		wasSetStatusProposedOnDharitrI,
		proposeSetStatusOnDharitrI,
		getAndStoreActionIDForProposeSetStatusFromDharitrI,
		wasActionSignedOnDharitrI,
		processQuorumReachedOnDharitrI,
		wasActionPerformedOnDharitrI,
		performActionOnDharitrI,
		signActionOnDharitrI,
	}

	for _, stepThatError := range stepsThatCanError {
		testErrorFlow(t, stepThatError)
	}
}

func testErrorFlow(t *testing.T, stepThatErrors bridgeCore.StepIdentifier) {
	t.Logf("\n\n\nnew test for stepThatError: %s", stepThatErrors)
	numCalled := 0
	args := argsBridgeStub{
		failingStep:                             string(stepThatErrors),
		myTurnHandler:                           trueHandler,
		processQuorumReachedOnPeerChainHandler:  trueHandler,
		processQuorumReachedOnDharitrIHandler: trueHandler,
		wasActionSignedOnDharitrIHandler:      trueHandler,
		wasActionPerformedOnDharitrIHandler: func() bool {
			numCalled++
			return numCalled > 1
		},
		wasTransferPerformedOnPeerChainHandler:  falseHandler,
		maxRetriesReachedPeerChainHandler:       falseHandler,
		maxRetriesReachedDharitrIHandler:      falseHandler,
		wasSetStatusProposedOnDharitrIHandler: falseHandler,
	}

	if stepThatErrors == "SignActionOnDharitrI" {
		args.wasActionSignedOnDharitrIHandler = falseHandler
	}

	executor, eh := createMockBridge(args)
	sm := createStateMachine(t, executor, GettingPendingBatchFromDharitrI)

	maxNumSteps := 12
	for i := 0; i < maxNumSteps; i++ {
		err := sm.Execute(context.Background())
		assert.Nil(t, err)

		if eh.lastError != nil {
			if sm.CurrentStep.Identifier() == GettingPendingBatchFromDharitrI {
				return
			}

			require.Fail(t, fmt.Sprintf("should have jumped to initial step, got next step %s, stepThatErrors %s",
				sm.CurrentStep.Identifier(), stepThatErrors))
		}
	}

	require.Fail(t, fmt.Sprintf("max number of steps reached but not jumped to initial step, stepThatErrors %s", stepThatErrors))
}
