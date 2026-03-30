package factory

import (
	"github.com/TerraDharitri/drt-go-bridge-eth/clients"
	"github.com/TerraDharitri/drt-go-bridge-eth/clients/gasManagement"
	"github.com/TerraDharitri/drt-go-bridge-eth/clients/gasManagement/disabled"
)

// CreateGasStation generates an implementation of GasHandler
func CreateGasStation(args gasManagement.ArgsGasStation, enabled bool) (clients.GasHandler, error) {
	if enabled {
		return gasManagement.NewGasStation(args)
	}
	return &disabled.DisabledGasStation{}, nil
}
