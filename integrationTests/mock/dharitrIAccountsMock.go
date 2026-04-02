package mock

import (
	"github.com/TerraDharitri/drt-go-sdk/core"
	"github.com/TerraDharitri/drt-go-sdk/data"
)

type dharitrIAccountsMock struct {
	accounts map[string]*data.Account
}

func newDharitrIAccountsMock() *dharitrIAccountsMock {
	return &dharitrIAccountsMock{
		accounts: make(map[string]*data.Account),
	}
}

func (mock *dharitrIAccountsMock) getOrCreate(address core.AddressHandler) *data.Account {
	addrAsString := string(address.AddressBytes())
	acc, found := mock.accounts[addrAsString]
	if !found {
		acc = &data.Account{}
		mock.accounts[addrAsString] = acc
	}

	return acc
}

func (mock *dharitrIAccountsMock) updateNonce(address core.AddressHandler, nonce uint64) {
	acc := mock.getOrCreate(address)
	acc.Nonce = nonce
}
