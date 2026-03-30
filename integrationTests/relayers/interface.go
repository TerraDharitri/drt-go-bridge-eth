package relayers

import (
	sdkCore "github.com/TerraDharitri/drt-go-sdk/core"
)

type bridgeComponents interface {
	DharitrIRelayerAddress() sdkCore.AddressHandler
	PeerChainRelayerAddress() string
	Start() error
	Close() error
}
