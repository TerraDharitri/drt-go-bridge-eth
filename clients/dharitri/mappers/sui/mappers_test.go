package sui

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
		t.Run("SuiToDharitrI: nil dataGetter", func(t *testing.T) {
			mapper, err := NewSuiToDharitrIMapper(nil)
			assert.Equal(t, clients.ErrNilDataGetter, err)
			assert.True(t, check.IfNil(mapper))
		})
		t.Run("SuiToDharitrI: should work", func(t *testing.T) {
			mapper, err := NewSuiToDharitrIMapper(&bridgeTests.DataGetterStub{})
			assert.Nil(t, err)
			assert.False(t, check.IfNil(mapper))
		})
	}
	{
		t.Run("DharitrIToSui: nil dataGetter", func(t *testing.T) {
			mapper, err := NewDharitrIToSuiMapper(nil)
			assert.Equal(t, clients.ErrNilDataGetter, err)
			assert.True(t, check.IfNil(mapper))
		})
		t.Run("DharitrIToSui: should work", func(t *testing.T) {
			mapper, err := NewDharitrIToSuiMapper(&bridgeTests.DataGetterStub{})
			assert.Nil(t, err)
			assert.False(t, check.IfNil(mapper))
		})
	}
}

func TestConvertToken(t *testing.T) {
	t.Parallel()

	{
		t.Run("DharitrIToSui: dataGetter returns error", func(t *testing.T) {
			expectedError := errors.New("expected error")
			dg := &bridgeTests.DataGetterStub{
				GetSuiCoinForTokenIdCalled: func(ctx context.Context, tokenId []byte) ([][]byte, error) {
					return nil, expectedError
				}}
			mapper, err := NewDharitrIToSuiMapper(dg)
			assert.Nil(t, err)
			assert.False(t, check.IfNil(mapper))

			_, err = mapper.ConvertToken(context.Background(), []byte("drtAddress"))
			assert.Equal(t, expectedError, err)
		})
		t.Run("DharitrIToSui: should work", func(t *testing.T) {
			expectedSuiAddress := []byte("erc20Address")
			dg := &bridgeTests.DataGetterStub{
				GetSuiCoinForTokenIdCalled: func(ctx context.Context, tokenId []byte) ([][]byte, error) {
					return [][]byte{expectedSuiAddress}, nil
				}}
			mapper, err := NewDharitrIToSuiMapper(dg)
			assert.Nil(t, err)
			assert.False(t, check.IfNil(mapper))
			erc20AddressReturned, err := mapper.ConvertToken(context.Background(), []byte("drtAddress"))
			assert.Nil(t, err)
			assert.Equal(t, expectedSuiAddress, erc20AddressReturned)
		})
	}
	{
		t.Run("SuiToDharitrI: dataGetter returns error", func(t *testing.T) {
			expectedError := errors.New("expected error")
			dg := &bridgeTests.DataGetterStub{
				GetTokenIdForSuiCoinCalled: func(ctx context.Context, erc20Address []byte) ([][]byte, error) {
					return nil, expectedError
				}}
			mapper, err := NewSuiToDharitrIMapper(dg)
			assert.Nil(t, err)
			assert.False(t, check.IfNil(mapper))

			_, err = mapper.ConvertToken(context.Background(), []byte("erc20Address"))
			assert.Equal(t, expectedError, err)
		})
		t.Run("SuiToDharitrI: should work", func(t *testing.T) {
			expectedDrtAddress := []byte("drtAddress")
			dg := &bridgeTests.DataGetterStub{
				GetTokenIdForSuiCoinCalled: func(ctx context.Context, erc20Address []byte) ([][]byte, error) {
					return [][]byte{expectedDrtAddress}, nil
				}}
			mapper, err := NewSuiToDharitrIMapper(dg)
			assert.Nil(t, err)
			assert.False(t, check.IfNil(mapper))
			drtAddressReturned, err := mapper.ConvertToken(context.Background(), []byte("erc20Address"))
			assert.Nil(t, err)
			assert.Equal(t, expectedDrtAddress, drtAddressReturned)
		})
	}
}
