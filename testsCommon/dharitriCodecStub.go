package testsCommon

import "github.com/TerraDharitri/drt-go-bridge-eth/parsers"

// DharitriCodecStub -
type DharitriCodecStub struct {
	DecodeProxySCCompleteCallDataCalled  func(buff []byte) (parsers.ProxySCCompleteCallData, error)
	ExtractGasLimitFromRawCallDataCalled func(buff []byte) (uint64, error)
}

// DecodeProxySCCompleteCallData -
func (stub *DharitriCodecStub) DecodeProxySCCompleteCallData(buff []byte) (parsers.ProxySCCompleteCallData, error) {
	if stub.DecodeProxySCCompleteCallDataCalled != nil {
		return stub.DecodeProxySCCompleteCallDataCalled(buff)
	}

	return parsers.ProxySCCompleteCallData{}, nil
}

// ExtractGasLimitFromRawCallData -
func (stub *DharitriCodecStub) ExtractGasLimitFromRawCallData(buff []byte) (uint64, error) {
	if stub.ExtractGasLimitFromRawCallDataCalled != nil {
		return stub.ExtractGasLimitFromRawCallDataCalled(buff)
	}

	return 0, nil
}

// IsInterfaceNil -
func (stub *DharitriCodecStub) IsInterfaceNil() bool {
	return stub == nil
}
