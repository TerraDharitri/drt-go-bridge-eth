package factory

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/TerraDharitri/drt-go-bridge-eth/config"
	"github.com/TerraDharitri/drt-go-bridge-eth/core"
	"github.com/TerraDharitri/drt-go-bridge-eth/status"
	"github.com/TerraDharitri/drt-go-bridge-eth/testsCommon"
	p2pMocks "github.com/TerraDharitri/drt-go-bridge-eth/testsCommon/p2p"
	"github.com/TerraDharitri/drt-go-chain/testscommon/statusHandler"
	"github.com/TerraDharitri/drt-go-sdk/blockchain"
	sdkCore "github.com/TerraDharitri/drt-go-sdk/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createMockBridgeCommonArgs() ArgsBridgeCommon {
	stateMachineConfig := config.ConfigStateMachine{
		StepDurationInMillis:       1000,
		IntervalForLeaderInSeconds: 60,
	}

	cfg := config.Config{
		DharitrI: config.DharitrIConfig{
			PrivateKeyFile:                  "testdata/grace.pem",
			IntervalToResendTxsInSeconds:    60,
			NetworkAddress:                  "http://127.0.0.1:8079",
			MultisigContractAddress:         "drt1qqqqqqqqqqqqqpgqgftcwj09u0nhmskrw7xxqcqh8qmzwyexd8ssr4um9c",
			SafeContractAddress:             "drt1qqqqqqqqqqqqqpgqgftcwj09u0nhmskrw7xxqcqh8qmzwyexd8ssr4um9c",
			GasMap:                          testsCommon.CreateTestDharitrIGasMap(),
			MaxRetriesOnQuorumReached:       1,
			MaxRetriesOnWasTransferProposed: 1,
			ClientAvailabilityAllowDelta:    10,
			Proxy: config.ProxyConfig{
				CacherExpirationSeconds: 600,
				RestAPIEntityType:       "observer",
				MaxNoncesDelta:          10,
				FinalityCheck:           true,
			},
		},
		Relayer: config.ConfigRelayer{
			RoleProvider: config.RoleProviderConfig{
				PollingIntervalInMillis: 1000,
			},
		},
		StateMachine: map[string]config.ConfigStateMachine{
			"toDharitrI":   stateMachineConfig,
			"fromDharitrI": stateMachineConfig,
		},
	}
	configs := config.Configs{
		GeneralConfig:   cfg,
		ApiRoutesConfig: config.ApiRoutesConfig{},
		FlagsConfig: config.ContextFlagsConfig{
			RestApiInterface: core.WebServerOffString,
		},
	}

	argsProxy := blockchain.ArgsProxy{
		ProxyURL:            cfg.DharitrI.NetworkAddress,
		CacheExpirationTime: time.Minute,
		EntityType:          sdkCore.ObserverNode,
	}
	proxy, _ := blockchain.NewProxy(argsProxy)
	return ArgsBridgeCommon{
		Configs:                       configs,
		Messenger:                     &p2pMocks.MessengerStub{},
		StatusStorer:                  testsCommon.NewStorerMock(),
		Proxy:                         proxy,
		DharitrIClientStatusHandler: &testsCommon.StatusHandlerStub{},
		TimeForBootstrap:              minTimeForBootstrap,
		TimeBeforeRepeatJoin:          minTimeBeforeRepeatJoin,
		MetricsHolder:                 status.NewMetricsHolder(),
		AppStatusHandler:              &statusHandler.AppStatusHandlerStub{},
	}
}

func TestNewBaseComponents(t *testing.T) {
	t.Parallel()

	t.Run("nil Proxy", func(t *testing.T) {
		t.Parallel()
		args := createMockBridgeCommonArgs()
		args.Proxy = nil

		components, err := NewBaseComponents(args)
		assert.True(t, errors.Is(err, errNilProxy))
		assert.Nil(t, components)
	})
	t.Run("nil Messenger", func(t *testing.T) {
		t.Parallel()
		args := createMockBridgeCommonArgs()
		args.Messenger = nil

		components, err := NewBaseComponents(args)
		assert.Equal(t, errNilMessenger, err)
		assert.Nil(t, components)
	})
	t.Run("nil StatusStorer", func(t *testing.T) {
		t.Parallel()
		args := createMockBridgeCommonArgs()
		args.StatusStorer = nil

		components, err := NewBaseComponents(args)
		assert.Equal(t, errNilStatusStorer, err)
		assert.Nil(t, components)
	})
	t.Run("err on createDharitrIKeysAndAddresses, empty pk file", func(t *testing.T) {
		t.Parallel()
		args := createMockBridgeCommonArgs()
		args.Configs.GeneralConfig.DharitrI.PrivateKeyFile = ""

		components, err := NewBaseComponents(args)
		assert.NotNil(t, err)
		assert.Nil(t, components)
	})
	t.Run("err on createDharitrIKeysAndAddresses, empty multisig address", func(t *testing.T) {
		t.Parallel()
		args := createMockBridgeCommonArgs()
		args.Configs.GeneralConfig.DharitrI.MultisigContractAddress = ""

		components, err := NewBaseComponents(args)
		assert.NotNil(t, err)
		assert.Nil(t, components)
	})
	t.Run("invalid time for bootstrap", func(t *testing.T) {
		t.Parallel()
		args := createMockBridgeCommonArgs()
		args.TimeForBootstrap = minTimeForBootstrap - 1

		components, err := NewBaseComponents(args)
		assert.True(t, errors.Is(err, errInvalidValue))
		assert.True(t, strings.Contains(err.Error(), "for TimeForBootstrap"))
		assert.Nil(t, components)
	})
	t.Run("invalid time before retry", func(t *testing.T) {
		t.Parallel()
		args := createMockBridgeCommonArgs()
		args.TimeBeforeRepeatJoin = minTimeBeforeRepeatJoin - 1

		components, err := NewBaseComponents(args)
		assert.True(t, errors.Is(err, errInvalidValue))
		assert.True(t, strings.Contains(err.Error(), "for TimeBeforeRepeatJoin"))
		assert.Nil(t, components)
	})
	t.Run("nil MetricsHolder", func(t *testing.T) {
		t.Parallel()
		args := createMockBridgeCommonArgs()
		args.MetricsHolder = nil

		components, err := NewBaseComponents(args)
		assert.Equal(t, errNilMetricsHolder, err)
		assert.Nil(t, components)
	})
	t.Run("should work", func(t *testing.T) {
		t.Parallel()
		args := createMockBridgeCommonArgs()

		components, err := NewBaseComponents(args)
		require.Nil(t, err)
		require.NotNil(t, components)
		require.Equal(t, 1, len(components.closableHandlers))
	})
}

func TestBaseBridgeComponents_DrtRelayerAddresses(t *testing.T) {
	t.Parallel()

	args := createMockBridgeCommonArgs()
	components, _ := NewBaseComponents(args)

	bech32Address, _ := components.DharitrIRelayerAddress().AddressAsBech32String()
	assert.Equal(t, "drt1r69gk66fmedhhcg24g2c5kn2f2a5k4kvpr6jfw67dn2lyydd8cfsncd6w8", bech32Address)
}
