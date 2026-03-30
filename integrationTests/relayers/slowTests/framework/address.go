package framework

import (
	"encoding/hex"
	"testing"

	sdkCore "github.com/TerraDharitri/drt-go-sdk/core"
	"github.com/TerraDharitri/drt-go-sdk/data"
	"github.com/stretchr/testify/require"
)

// DrtAddress holds the different forms a DharitrI address might have
type DrtAddress struct {
	sdkCore.AddressHandler
	bytes  []byte
	bech32 string
	hex    string
}

// NewDrtAddressFromBytes return a new instance of DrtAddress from bytes
func NewDrtAddressFromBytes(tb testing.TB, bytes []byte) *DrtAddress {
	address := &DrtAddress{
		bytes:          make([]byte, len(bytes)),
		hex:            hex.EncodeToString(bytes),
		AddressHandler: data.NewAddressFromBytes(bytes),
	}

	var err error
	copy(address.bytes, bytes)
	address.bech32, err = addressPubkeyConverter.Encode(bytes)
	require.Nil(tb, err)

	return address
}

// NewDrtAddressFromBech32 return a new instance of DrtAddress from the bech32 string
func NewDrtAddressFromBech32(tb testing.TB, bech32 string) *DrtAddress {
	addressHandler, err := data.NewAddressFromBech32String(bech32)
	require.Nil(tb, err)

	return &DrtAddress{
		bytes:          addressHandler.AddressBytes(),
		hex:            hex.EncodeToString(addressHandler.AddressBytes()),
		bech32:         bech32,
		AddressHandler: addressHandler,
	}
}

// Bytes returns the bytes format address
func (address *DrtAddress) Bytes() []byte {
	return address.bytes
}

// Bech32 returns the bech32 string format address
func (address *DrtAddress) Bech32() string {
	return address.bech32
}

// Hex returns the hex string format address
func (address *DrtAddress) Hex() string {
	return address.hex
}

// String returns the address in bech32 format
func (address *DrtAddress) String() string {
	return address.bech32
}
