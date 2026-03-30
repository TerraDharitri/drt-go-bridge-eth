package framework

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/TerraDharitri/drt-go-bridge-eth/clients/ethereum"
	"github.com/TerraDharitri/drt-go-bridge-eth/config"
	bridgeCore "github.com/TerraDharitri/drt-go-bridge-eth/core"
	"github.com/TerraDharitri/drt-go-bridge-eth/factory"
	"github.com/TerraDharitri/drt-go-bridge-eth/integrationTests"
	testsRelayers "github.com/TerraDharitri/drt-go-bridge-eth/integrationTests/relayers"
	"github.com/TerraDharitri/drt-go-bridge-eth/status"
	"github.com/TerraDharitri/drt-go-bridge-eth/testsCommon"
	"github.com/TerraDharitri/drt-go-chain/testscommon/statusHandler"
	"github.com/stretchr/testify/require"
)

const (
	relayerETHKeyPathFormat  = "../testdata/ethereum%d.sk"
	relayerSuiSeedPathFormat = "../testdata/sui%d.sk"
)

// BridgeComponents holds and manages the relayers components
type BridgeComponents struct {
	testing.TB
	RelayerInstances   []Relayer
	gasStationInstance *gasStation
}

// NewEthereumBridgeComponents will create the bridge components (relayers) for Ethereum
func NewEthereumBridgeComponents(
	tb testing.TB,
	workingDir string,
	chainSimulator ChainSimulatorWrapper,
	ethereumChain ethereum.ClientWrapper,
	erc20ContractsHolder ethereum.Erc20ContractsHolder,
	ethBackend *simulated.Backend,
	numRelayers int,
	ethSafeContractAddress string,
	drtSafeAddress *DrtAddress,
	drtMultisigAddress *DrtAddress,
) *BridgeComponents {
	bridge := &BridgeComponents{
		TB:                 tb,
		RelayerInstances:   make([]Relayer, 0, numRelayers),
		gasStationInstance: NewGasStation(ethBackend),
	}

	messengers := integrationTests.CreateLinkedMessengers(numRelayers)

	gasStationURL := bridge.gasStationInstance.URL()
	log.Info("started gas station server", "URL", gasStationURL)

	wg := sync.WaitGroup{}
	wg.Add(numRelayers)

	for i := 0; i < numRelayers; i++ {
		generalConfigs := testsRelayers.CreateEthDrtBridgeComponentsConfig(i, workingDir, gasStationURL)
		generalConfigs.Eth.PrivateKeyFile = fmt.Sprintf(relayerETHKeyPathFormat, i)
		argsBridgeComponents := factory.ArgsEthereumToDharitrIBridge{
			Configs: config.Configs{
				GeneralConfig:   generalConfigs,
				ApiRoutesConfig: config.ApiRoutesConfig{},
				FlagsConfig: config.ContextFlagsConfig{
					RestApiInterface: bridgeCore.WebServerOffString,
				},
			},
			Proxy:                         chainSimulator.Proxy(),
			ClientWrapper:                 ethereumChain,
			Messenger:                     messengers[i],
			StatusStorer:                  testsCommon.NewStorerMock(),
			TimeForBootstrap:              time.Second * 5,
			TimeBeforeRepeatJoin:          time.Second * 30,
			MetricsHolder:                 status.NewMetricsHolder(),
			AppStatusHandler:              &statusHandler.AppStatusHandlerStub{},
			DharitrIClientStatusHandler: &testsCommon.StatusHandlerStub{},
		}
		argsBridgeComponents.Configs.GeneralConfig.Eth.SafeContractAddress = ethSafeContractAddress
		argsBridgeComponents.Erc20ContractsHolder = erc20ContractsHolder
		argsBridgeComponents.Configs.GeneralConfig.DharitrI.NetworkAddress = chainSimulator.GetNetworkAddress()
		argsBridgeComponents.Configs.GeneralConfig.DharitrI.SafeContractAddress = drtSafeAddress.Bech32()
		argsBridgeComponents.Configs.GeneralConfig.DharitrI.MultisigContractAddress = drtMultisigAddress.Bech32()
		argsBridgeComponents.Configs.GeneralConfig.DharitrI.GasMap = config.DharitrIGasMapConfig{
			Sign:                   8000000,
			ProposeTransferBase:    11000000,
			ProposeTransferForEach: 5500000,
			ProposeStatusBase:      10000000,
			ProposeStatusForEach:   7000000,
			PerformActionBase:      40000000,
			PerformActionForEach:   5500000,
			ScCallPerByte:          100000,
			ScCallPerformForEach:   10000000,
		}
		relayer, err := factory.NewEthDrtBridgeComponents(argsBridgeComponents)
		require.Nil(bridge, err)

		go func() {
			err = relayer.Start()
			log.LogIfError(err)
			require.Nil(bridge, err)
			wg.Done()
		}()

		bridge.RelayerInstances = append(bridge.RelayerInstances, relayer)
	}

	// ensure all relayers are successfully started before returning the bridge components instance
	wg.Wait()

	return bridge
}

// NewSuiBridgeComponents will create the bridge components (relayers) for Sui
func NewSuiBridgeComponents(
	tb testing.TB,
	workingDir string,
	chainSimulator ChainSimulatorWrapper,
	suiChainSimulator *suiChainSimulatorWrapper,
	numRelayers int,
	packageId string,
	drtSafeAddress *DrtAddress,
	drtMultisigAddress *DrtAddress,
	bridgeObjectId string,
	safeObjectId string,
	treasuryId string,
	bridgeInitialSharedVersion uint64,
	safeInitialSharedVersion uint64,
	treasuryInitialSharedVersion uint64,
) *BridgeComponents {
	bridge := &BridgeComponents{
		TB:                 tb,
		RelayerInstances:   make([]Relayer, 0, numRelayers),
		gasStationInstance: nil,
	}

	messengers := integrationTests.CreateLinkedMessengers(numRelayers)

	wg := sync.WaitGroup{}
	wg.Add(numRelayers)

	for i := 0; i < numRelayers; i++ {
		generalConfigs := testsRelayers.CreateSuiDrtBridgeComponentsConfig(i, workingDir)
		generalConfigs.Sui.PrivateKeyFile = fmt.Sprintf(relayerSuiSeedPathFormat, i)
		argsBridgeComponents := factory.ArgsSuiToDharitrIBridge{
			Configs: config.Configs{
				GeneralConfig:   generalConfigs,
				ApiRoutesConfig: config.ApiRoutesConfig{},
				FlagsConfig: config.ContextFlagsConfig{
					RestApiInterface: bridgeCore.WebServerOffString,
				},
			},
			Proxy:                         chainSimulator.Proxy(),
			SuiProxy:                      suiChainSimulator.proxy,
			Messenger:                     messengers[i],
			StatusStorer:                  testsCommon.NewStorerMock(),
			TimeForBootstrap:              time.Second * 5,
			TimeBeforeRepeatJoin:          time.Second * 30,
			MetricsHolder:                 status.NewMetricsHolder(),
			AppStatusHandler:              &statusHandler.AppStatusHandlerStub{},
			DharitrIClientStatusHandler: &testsCommon.StatusHandlerStub{},
			SuiClientStatusHandler:        &testsCommon.StatusHandlerStub{},
		}
		argsBridgeComponents.Configs.GeneralConfig.Sui.PackageId = packageId
		argsBridgeComponents.Configs.GeneralConfig.Sui.BridgeObjectId = bridgeObjectId
		argsBridgeComponents.Configs.GeneralConfig.Sui.BridgeObjectInitialSharedVersion = bridgeInitialSharedVersion
		argsBridgeComponents.Configs.GeneralConfig.Sui.SafeObjectId = safeObjectId
		argsBridgeComponents.Configs.GeneralConfig.Sui.SafeObjectInitialSharedVersion = safeInitialSharedVersion
		argsBridgeComponents.Configs.GeneralConfig.Sui.TreasuryObjectId = treasuryId
		argsBridgeComponents.Configs.GeneralConfig.Sui.TreasuryObjectInitialSharedVersion = treasuryInitialSharedVersion
		argsBridgeComponents.Configs.GeneralConfig.DharitrI.NetworkAddress = chainSimulator.GetNetworkAddress()
		argsBridgeComponents.Configs.GeneralConfig.DharitrI.SafeContractAddress = drtSafeAddress.Bech32()
		argsBridgeComponents.Configs.GeneralConfig.DharitrI.MultisigContractAddress = drtMultisigAddress.Bech32()
		argsBridgeComponents.Configs.GeneralConfig.DharitrI.GasMap = config.DharitrIGasMapConfig{
			Sign:                   8000000,
			ProposeTransferBase:    11000000,
			ProposeTransferForEach: 5500000,
			ProposeStatusBase:      10000000,
			ProposeStatusForEach:   7000000,
			PerformActionBase:      40000000,
			PerformActionForEach:   5500000,
			ScCallPerByte:          100000,
			ScCallPerformForEach:   10000000,
		}
		relayer, err := factory.NewSuiDrtBridgeComponents(argsBridgeComponents)
		require.Nil(bridge, err)

		go func() {
			err = relayer.Start()
			log.LogIfError(err)
			require.Nil(bridge, err)
			wg.Done()
		}()

		bridge.RelayerInstances = append(bridge.RelayerInstances, relayer)
	}

	// ensure all relayers are successfully started before returning the bridge components instance
	wg.Wait()

	return bridge

}

// CloseRelayers will call close on all created relayers
func (bridge *BridgeComponents) CloseRelayers() {
	if bridge.gasStationInstance != nil {
		bridge.gasStationInstance.Close()
	}

	for _, r := range bridge.RelayerInstances {
		_ = r.Close()
	}
}
