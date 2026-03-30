package testsCommon

import (
	"github.com/TerraDharitri/drt-go-bridge-eth/config"
)

// CreateTestDharitrIGasMap will create a testing gas map for DharitrI client
func CreateTestDharitrIGasMap() config.DharitrIGasMapConfig {
	return config.DharitrIGasMapConfig{
		Sign:                   101,
		ProposeTransferBase:    102,
		ProposeTransferForEach: 103,
		ProposeStatusBase:      104,
		ProposeStatusForEach:   105,
		PerformActionBase:      106,
		PerformActionForEach:   107,
		ScCallPerByte:          108,
		ScCallPerformForEach:   109,
	}
}
