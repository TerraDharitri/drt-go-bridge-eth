package roleproviders

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"strings"
	"testing"

	"github.com/TerraDharitri/drt-go-bridge-eth/clients"
	bridgeTests "github.com/TerraDharitri/drt-go-bridge-eth/testsCommon/bridge"
	"github.com/TerraDharitri/drt-go-chain-core/core/check"
	logger "github.com/TerraDharitri/drt-go-chain-logger"
	"github.com/TerraDharitri/drt-go-sdk/data"
	"github.com/stretchr/testify/assert"
)

func createMockArgs() ArgsDharitrIRoleProvider {
	return ArgsDharitrIRoleProvider{
		Log:        logger.GetOrCreate("test"),
		DataGetter: &bridgeTests.DataGetterStub{},
	}
}

func TestNewDharitrIRoleProvider(t *testing.T) {
	t.Parallel()

	t.Run("nil data getter should error", func(t *testing.T) {
		t.Parallel()

		args := createMockArgs()
		args.DataGetter = nil

		erp, err := NewDharitrIRoleProvider(args)
		assert.True(t, check.IfNil(erp))
		assert.Equal(t, clients.ErrNilDataGetter, err)
	})
	t.Run("nil logger should error", func(t *testing.T) {
		t.Parallel()

		args := createMockArgs()
		args.Log = nil

		erp, err := NewDharitrIRoleProvider(args)
		assert.True(t, check.IfNil(erp))
		assert.Equal(t, clients.ErrNilLogger, err)
	})
	t.Run("should work", func(t *testing.T) {
		t.Parallel()

		args := createMockArgs()

		erp, err := NewDharitrIRoleProvider(args)
		assert.False(t, check.IfNil(erp))
		assert.Nil(t, err)
	})
}

func TestDharitrIRoleProvider_ExecuteErrors(t *testing.T) {
	t.Parallel()

	expectedErr := errors.New("expected error")
	args := createMockArgs()
	args.DataGetter = &bridgeTests.DataGetterStub{
		GetAllStakedRelayersCalled: func(ctx context.Context) ([][]byte, error) {
			return nil, expectedErr
		},
	}

	erp, _ := NewDharitrIRoleProvider(args)
	err := erp.Execute(context.TODO())
	assert.Equal(t, expectedErr, err)
}

func TestDharitrIRoleProvider_ExecuteShouldWork(t *testing.T) {
	t.Parallel()

	whitelistedAddresses := [][]byte{
		bytes.Repeat([]byte("1"), 32),
		bytes.Repeat([]byte("3"), 32),
		bytes.Repeat([]byte("2"), 32),
	}
	expectedSortedPublicKeys := [][]byte{
		bytes.Repeat([]byte("1"), 32),
		bytes.Repeat([]byte("2"), 32),
		bytes.Repeat([]byte("3"), 32),
	}

	t.Run("nil whitelisted", testDharitrIExecuteShouldWork(nil, make([][]byte, 0)))
	t.Run("empty whitelisted", testDharitrIExecuteShouldWork(make([][]byte, 0), make([][]byte, 0)))
	t.Run("with whitelisted", testDharitrIExecuteShouldWork(whitelistedAddresses, expectedSortedPublicKeys))
}

func testDharitrIExecuteShouldWork(whitelistedAddresses [][]byte, expectedSortedPublicKeys [][]byte) func(t *testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		args := createMockArgs()
		args.DataGetter = &bridgeTests.DataGetterStub{
			GetAllStakedRelayersCalled: func(ctx context.Context) ([][]byte, error) {
				return whitelistedAddresses, nil
			},
		}

		erp, _ := NewDharitrIRoleProvider(args)
		err := erp.Execute(context.TODO())
		assert.Nil(t, err)

		for _, addr := range whitelistedAddresses {
			addressHandler := data.NewAddressFromBytes(addr)
			assert.True(t, erp.IsWhitelisted(addressHandler))
		}

		randomAddress := data.NewAddressFromBytes([]byte("random address"))
		assert.False(t, erp.IsWhitelisted(randomAddress))
		assert.False(t, erp.IsWhitelisted(nil))
		erp.mut.RLock()
		assert.Equal(t, len(whitelistedAddresses), len(erp.whitelistedAddresses))
		erp.mut.RUnlock()
		sortedPublicKeys := erp.SortedPublicKeys()
		assert.Equal(t, expectedSortedPublicKeys, sortedPublicKeys)
	}
}

func TestDharitrIRoleProvider_MisconfiguredAddressesShouldError(t *testing.T) {
	t.Parallel()

	misconfiguredAddresses := [][]byte{
		bytes.Repeat([]byte("1"), 32),
		bytes.Repeat([]byte("2"), 32),
		[]byte("bad address"),
	}

	args := createMockArgs()
	args.DataGetter = &bridgeTests.DataGetterStub{
		GetAllStakedRelayersCalled: func(ctx context.Context) ([][]byte, error) {
			return misconfiguredAddresses, nil
		},
	}

	erp, _ := NewDharitrIRoleProvider(args)
	err := erp.Execute(context.TODO())
	assert.True(t, errors.Is(err, ErrInvalidAddressBytes))
	assert.True(t, strings.Contains(err.Error(), hex.EncodeToString(misconfiguredAddresses[2])))
	assert.Zero(t, len(erp.whitelistedAddresses))
}
