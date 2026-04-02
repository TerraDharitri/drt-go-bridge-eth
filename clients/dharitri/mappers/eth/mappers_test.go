package eth

import (
	"context"
	"errors"
	"testing"

	"github.com/TerraDharitri/drt-go-bridge-eth/clients"
	bridgeTests "github.com/TerraDharitri/drt-go-bridge-eth/testsCommon/bridge"
	"github.com/TerraDharitri/drt-go-chain-core/core/check"
	"github.com/stretchr/testify/assert"
)

func TestNewMapper(t *testing.T) {
	t.Parallel()
	{
		t.Run("Erc20ToDharitrI: nil dataGetter", func(t *testing.T) {
			mapper, err := NewErc20ToDharitrIMapper(nil)
			assert.Equal(t, clients.ErrNilDataGetter, err)
			assert.True(t, check.IfNil(mapper))
		})
		t.Run("Erc20ToDharitrI: should work", func(t *testing.T) {
			mapper, err := NewErc20ToDharitrIMapper(&bridgeTests.DataGetterStub{})
			assert.Nil(t, err)
			assert.False(t, check.IfNil(mapper))
		})
	}
	{
		t.Run("DharitrIToErc20: nil dataGetter", func(t *testing.T) {
			mapper, err := NewDharitrIToErc20Mapper(nil)
			assert.Equal(t, clients.ErrNilDataGetter, err)
			assert.True(t, check.IfNil(mapper))
		})
		t.Run("DharitrIToErc20: should work", func(t *testing.T) {
			mapper, err := NewDharitrIToErc20Mapper(&bridgeTests.DataGetterStub{})
			assert.Nil(t, err)
			assert.False(t, check.IfNil(mapper))
		})
	}
}

func TestConvertToken(t *testing.T) {
	t.Parallel()

	{
		t.Run("DharitrIToErc20: dataGetter returns error", func(t *testing.T) {
			expectedError := errors.New("expected error")
			dg := &bridgeTests.DataGetterStub{
				GetERC20AddressForTokenIdCalled: func(ctx context.Context, tokenId []byte) ([][]byte, error) {
					return nil, expectedError
				}}
			mapper, err := NewDharitrIToErc20Mapper(dg)
			assert.Nil(t, err)
			assert.False(t, check.IfNil(mapper))

			_, err = mapper.ConvertToken(context.Background(), []byte("drtAddress"))
			assert.Equal(t, expectedError, err)
		})
		t.Run("DharitrIToErc20: should work", func(t *testing.T) {
			expectedErc20Address := []byte("erc20Address")
			dg := &bridgeTests.DataGetterStub{
				GetERC20AddressForTokenIdCalled: func(ctx context.Context, tokenId []byte) ([][]byte, error) {
					return [][]byte{expectedErc20Address}, nil
				}}
			mapper, err := NewDharitrIToErc20Mapper(dg)
			assert.Nil(t, err)
			assert.False(t, check.IfNil(mapper))
			erc20AddressReturned, err := mapper.ConvertToken(context.Background(), []byte("drtAddress"))
			assert.Nil(t, err)
			assert.Equal(t, expectedErc20Address, erc20AddressReturned)
		})
	}
	{
		t.Run("Erc20ToDharitrI: dataGetter returns error", func(t *testing.T) {
			expectedError := errors.New("expected error")
			dg := &bridgeTests.DataGetterStub{
				GetTokenIdForErc20AddressCalled: func(ctx context.Context, erc20Address []byte) ([][]byte, error) {
					return nil, expectedError
				}}
			mapper, err := NewErc20ToDharitrIMapper(dg)
			assert.Nil(t, err)
			assert.False(t, check.IfNil(mapper))

			_, err = mapper.ConvertToken(context.Background(), []byte("erc20Address"))
			assert.Equal(t, expectedError, err)
		})
		t.Run("Erc20ToDharitrI: should work", func(t *testing.T) {
			expectedDrtAddress := []byte("drtAddress")
			dg := &bridgeTests.DataGetterStub{
				GetTokenIdForErc20AddressCalled: func(ctx context.Context, erc20Address []byte) ([][]byte, error) {
					return [][]byte{expectedDrtAddress}, nil
				}}
			mapper, err := NewErc20ToDharitrIMapper(dg)
			assert.Nil(t, err)
			assert.False(t, check.IfNil(mapper))
			drtAddressReturned, err := mapper.ConvertToken(context.Background(), []byte("erc20Address"))
			assert.Nil(t, err)
			assert.Equal(t, expectedDrtAddress, drtAddressReturned)
		})
	}
}
