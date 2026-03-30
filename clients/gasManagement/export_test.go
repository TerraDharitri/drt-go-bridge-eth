package gasManagement

import "github.com/TerraDharitri/drt-go-bridge-eth/core"

// GetLatestGasPrice -
func (gs *gasStation) GetLatestGasPrice() float64 {
	gs.mut.RLock()
	defer gs.mut.RUnlock()

	return gs.latestGasPrice
}

// SetSelector -
func (gs *gasStation) SetSelector(gasPriceSelector core.EthGasPriceSelector) {
	gs.mut.Lock()
	defer gs.mut.Unlock()

	gs.gasPriceSelector = gasPriceSelector
}
