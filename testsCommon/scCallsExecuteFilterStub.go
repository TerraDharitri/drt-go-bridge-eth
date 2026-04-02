package testsCommon

import "github.com/TerraDharitri/drt-go-bridge-eth/parsers"

// ScCallsExecuteFilterStub -
type ScCallsExecuteFilterStub struct {
	ShouldExecuteCalled func(callData parsers.ProxySCCompleteCallData) bool
}

// ShouldExecute -
func (stub *ScCallsExecuteFilterStub) ShouldExecute(callData parsers.ProxySCCompleteCallData) bool {
	if stub.ShouldExecuteCalled != nil {
		return stub.ShouldExecuteCalled(callData)
	}

	return true
}

// IsInterfaceNil -
func (stub *ScCallsExecuteFilterStub) IsInterfaceNil() bool {
	return stub == nil
}
