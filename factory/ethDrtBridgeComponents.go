package factory

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/TerraDharitri/drt-go-bridge-eth/bridges"
	"github.com/TerraDharitri/drt-go-bridge-eth/bridges/disabled"
	"github.com/TerraDharitri/drt-go-bridge-eth/bridges/steps/fromDharitrI"
	"github.com/TerraDharitri/drt-go-bridge-eth/bridges/steps/toDharitrI"
	"github.com/TerraDharitri/drt-go-bridge-eth/bridges/topology"
	"github.com/TerraDharitri/drt-go-bridge-eth/clients/chain"
	"github.com/TerraDharitri/drt-go-bridge-eth/clients/ethereum"
	"github.com/TerraDharitri/drt-go-bridge-eth/clients/gasManagement"
	"github.com/TerraDharitri/drt-go-bridge-eth/clients/gasManagement/factory"
	"github.com/TerraDharitri/drt-go-bridge-eth/clients/dharitri"
	"github.com/TerraDharitri/drt-go-bridge-eth/clients/dharitri/mappers/eth"
	roleproviders "github.com/TerraDharitri/drt-go-bridge-eth/clients/roleProviders"
	"github.com/TerraDharitri/drt-go-bridge-eth/config"
	"github.com/TerraDharitri/drt-go-bridge-eth/core"
	"github.com/TerraDharitri/drt-go-bridge-eth/p2p"
	"github.com/TerraDharitri/drt-go-bridge-eth/stateMachine"
	"github.com/TerraDharitri/drt-go-bridge-eth/status"
	chainCore "github.com/TerraDharitri/drt-go-chain-core/core"
	"github.com/TerraDharitri/drt-go-chain-core/core/check"
	logger "github.com/TerraDharitri/drt-go-chain-logger"
	"github.com/TerraDharitri/drt-go-sdk/core/polling"
)

// ArgsEthereumToDharitrIBridge is the arguments DTO used for creating an Ethereum to DharitrI bridge
type ArgsEthereumToDharitrIBridge struct {
	Configs                       config.Configs
	Messenger                     p2p.NetMessenger
	StatusStorer                  core.Storer
	Proxy                         dharitri.Proxy
	DharitrIClientStatusHandler core.StatusHandler
	Erc20ContractsHolder          ethereum.Erc20ContractsHolder
	ClientWrapper                 ethereum.ClientWrapper
	TimeForBootstrap              time.Duration
	TimeBeforeRepeatJoin          time.Duration
	MetricsHolder                 core.MetricsHolder
	AppStatusHandler              chainCore.AppStatusHandler
}

type ethDrtBridgeComponents struct {
	*baseBridgeComponents
	evmCompatibleChain     chain.Chain
	ethClient              bridges.PeerChainClient
	ethereumRelayerAddress common.Address
	ethereumRoleProvider   PeerChainRoleProvider
}

func NewEthDrtBridgeComponents(args ArgsEthereumToDharitrIBridge) (*ethDrtBridgeComponents, error) {
	err := checkArgsEthereum(args)
	if err != nil {
		return nil, err
	}

	commonBridgeArgs := ArgsBridgeCommon{
		Configs:                       args.Configs,
		Messenger:                     args.Messenger,
		StatusStorer:                  args.StatusStorer,
		Proxy:                         args.Proxy,
		DharitrIClientStatusHandler: args.DharitrIClientStatusHandler,
		TimeForBootstrap:              args.TimeForBootstrap,
		TimeBeforeRepeatJoin:          args.TimeBeforeRepeatJoin,
		MetricsHolder:                 args.MetricsHolder,
		AppStatusHandler:              args.AppStatusHandler,
	}
	baseComponents, err := NewBaseComponents(commonBridgeArgs)
	if err != nil {
		return nil, err
	}

	components := &ethDrtBridgeComponents{
		baseBridgeComponents: baseComponents,
		evmCompatibleChain:   args.Configs.GeneralConfig.Eth.Chain,
	}

	err = components.initBaseComponents(commonBridgeArgs)
	if err != nil {
		return nil, err
	}

	err = components.createEthereumRoleProvider(args)
	if err != nil {
		return nil, err
	}

	err = components.createEthereumClient(args)
	if err != nil {
		return nil, err
	}

	err = components.createEthereumToDharitrIBridge(args)
	if err != nil {
		return nil, err
	}

	err = components.createEthereumToDharitrIStateMachine()
	if err != nil {
		return nil, err
	}

	err = components.createDharitrIToEthereumBridge(args)
	if err != nil {
		return nil, err
	}

	err = components.createDharitrIToEthereumStateMachine()
	if err != nil {
		return nil, err
	}

	return components, nil
}

func checkArgsEthereum(args ArgsEthereumToDharitrIBridge) error {
	if check.IfNil(args.ClientWrapper) {
		return errNilEthClient
	}
	if check.IfNil(args.Erc20ContractsHolder) {
		return errNilErc20ContractsHolder
	}
	return nil
}

func (components *ethDrtBridgeComponents) initBaseComponents(args ArgsBridgeCommon) error {
	ethToDharitrIName := components.evmCompatibleChain.PeerChainToDharitrIName()
	baseLogId := components.evmCompatibleChain.BaseLogId()
	components.baseLogger = core.NewLoggerWithIdentifier(logger.GetOrCreate(ethToDharitrIName), baseLogId)

	dataGetterLogId := components.evmCompatibleChain.DharitrIDataGetterLogId()
	err := components.createDataGetter(dataGetterLogId)
	if err != nil {
		return err
	}

	clientLogId := components.evmCompatibleChain.DharitrIClientLogId()
	err = components.createDharitrIRoleProvider(args, clientLogId)
	if err != nil {
		return err
	}

	roleProviderLogId := components.evmCompatibleChain.DharitrIRoleProviderLogId()
	tokensMapper, err := eth.NewDharitrIToErc20Mapper(components.drtDataGetter)
	if err != nil {
		return err
	}
	err = components.createDharitrIClient(args, roleProviderLogId, tokensMapper)
	if err != nil {
		return err
	}

	return nil
}

func (components *ethDrtBridgeComponents) createEthereumRoleProvider(args ArgsEthereumToDharitrIBridge) error {
	configs := args.Configs.GeneralConfig
	ethRoleProviderLogId := components.evmCompatibleChain.PeerChainRoleProviderLogId()
	log := core.NewLoggerWithIdentifier(logger.GetOrCreate(ethRoleProviderLogId), ethRoleProviderLogId)
	argsRoleProvider := roleproviders.ArgsEthereumRoleProvider{
		EthereumChainInteractor: args.ClientWrapper,
		Log:                     log,
	}

	var err error
	components.ethereumRoleProvider, err = roleproviders.NewEthereumRoleProvider(argsRoleProvider)
	if err != nil {
		return err
	}

	argsPollingHandler := polling.ArgsPollingHandler{
		Log:              log,
		Name:             string(components.evmCompatibleChain) + " role provider",
		PollingInterval:  time.Duration(configs.Relayer.RoleProvider.PollingIntervalInMillis) * time.Millisecond,
		PollingWhenError: pollingDurationOnError,
		Executor:         components.ethereumRoleProvider,
	}

	pollingHandler, err := polling.NewPollingHandler(argsPollingHandler)
	if err != nil {
		return err
	}

	components.addClosableComponent(pollingHandler)
	components.pollingHandlers = append(components.pollingHandlers, pollingHandler)

	return nil
}

func (components *ethDrtBridgeComponents) createEthereumClient(args ArgsEthereumToDharitrIBridge) error {
	ethereumConfigs := args.Configs.GeneralConfig.Eth

	gasStationConfig := ethereumConfigs.GasStation
	argsGasStation := gasManagement.ArgsGasStation{
		RequestURL:             gasStationConfig.URL,
		RequestPollingInterval: time.Duration(gasStationConfig.PollingIntervalInSeconds) * time.Second,
		RequestRetryDelay:      time.Duration(gasStationConfig.RequestRetryDelayInSeconds) * time.Second,
		MaximumFetchRetries:    gasStationConfig.MaxFetchRetries,
		RequestTime:            time.Duration(gasStationConfig.RequestTimeInSeconds) * time.Second,
		MaximumGasPrice:        gasStationConfig.MaximumAllowedGasPrice,
		GasPriceSelector:       core.EthGasPriceSelector(gasStationConfig.GasPriceSelector),
		GasPriceMultiplier:     gasStationConfig.GasPriceMultiplier,
	}

	gs, err := factory.CreateGasStation(argsGasStation, gasStationConfig.Enabled)
	if err != nil {
		return err
	}

	components.addClosableComponent(gs)

	antifloodComponents, err := components.createAntifloodComponents(args.Configs.GeneralConfig.P2P.AntifloodConfig)
	if err != nil {
		return err
	}

	peerDenialEvaluator, err := p2p.NewPeerDenialEvaluator(antifloodComponents.BlacklistHandler, antifloodComponents.PubKeysCacher)
	if err != nil {
		return err
	}
	err = args.Messenger.SetPeerDenialEvaluator(peerDenialEvaluator)
	if err != nil {
		return err
	}

	broadcasterLogId := components.evmCompatibleChain.BroadcasterLogId()
	ethToDharitrIName := components.evmCompatibleChain.PeerChainToDharitrIName()
	argsBroadcaster := p2p.ArgsBroadcaster{
		Messenger:              args.Messenger,
		Log:                    core.NewLoggerWithIdentifier(logger.GetOrCreate(broadcasterLogId), broadcasterLogId),
		DharitrIRoleProvider: components.dharitrIRoleProvider,
		SignatureProcessor:     components.ethereumRoleProvider,
		KeyGen:                 keyGen,
		SingleSigner:           singleSigner,
		PrivateKey:             components.dharitrIRelayerPrivateKey,
		Name:                   ethToDharitrIName,
		AntifloodComponents:    antifloodComponents,
	}

	components.broadcaster, err = p2p.NewBroadcaster(argsBroadcaster)
	if err != nil {
		return err
	}

	cryptoHandler, err := ethereum.NewCryptoHandler(ethereumConfigs.PrivateKeyFile)
	if err != nil {
		return err
	}

	components.ethereumRelayerAddress = cryptoHandler.GetAddress()

	tokensMapper, err := eth.NewErc20ToDharitrIMapper(components.drtDataGetter)
	if err != nil {
		return err
	}

	signaturesHolder := bridges.NewSignatureHolder()
	components.toDharitrISignaturesHolder = signaturesHolder
	err = components.broadcaster.AddBroadcastClient(signaturesHolder)
	if err != nil {
		return err
	}

	safeContractAddress := common.HexToAddress(ethereumConfigs.SafeContractAddress)

	ethClientLogId := components.evmCompatibleChain.PeerChainClientLogId()
	argsEthClient := ethereum.ArgsEthereumClient{
		ClientWrapper:                args.ClientWrapper,
		Erc20ContractsHandler:        args.Erc20ContractsHolder,
		Log:                          core.NewLoggerWithIdentifier(logger.GetOrCreate(ethClientLogId), ethClientLogId),
		AddressConverter:             components.addressConverter,
		Broadcaster:                  components.broadcaster,
		CryptoHandler:                cryptoHandler,
		TokensMapper:                 tokensMapper,
		SignatureHolder:              signaturesHolder,
		SafeContractAddress:          safeContractAddress,
		GasHandler:                   gs,
		TransferGasLimitBase:         ethereumConfigs.GasLimitBase,
		TransferGasLimitForEach:      ethereumConfigs.GasLimitForEach,
		ClientAvailabilityAllowDelta: ethereumConfigs.ClientAvailabilityAllowDelta,
		EventsBlockRangeFrom:         ethereumConfigs.EventsBlockRangeFrom,
		EventsBlockRangeTo:           ethereumConfigs.EventsBlockRangeTo,
	}

	components.ethClient, err = ethereum.NewEthereumClient(argsEthClient)

	return err
}

func (components *ethDrtBridgeComponents) createEthereumToDharitrIBridge(args ArgsEthereumToDharitrIBridge) error {
	ethToDharitrIName := components.evmCompatibleChain.PeerChainToDharitrIName()
	log := core.NewLoggerWithIdentifier(logger.GetOrCreate(ethToDharitrIName), ethToDharitrIName)

	configs, found := args.Configs.GeneralConfig.StateMachine[ethToDharitrIName]
	if !found {
		return fmt.Errorf("%w for %q", errMissingConfig, ethToDharitrIName)
	}

	components.toDharitrIStepDuration = time.Duration(configs.StepDurationInMillis) * time.Millisecond

	argsTopologyHandler := topology.ArgsTopologyHandler{
		PublicKeysProvider: components.dharitrIRoleProvider,
		Timer:              components.timer,
		IntervalForLeader:  time.Second * time.Duration(configs.IntervalForLeaderInSeconds),
		AddressBytes:       components.dharitrIRelayerAddress.AddressBytes(),
		Log:                log,
		AddressConverter:   components.addressConverter,
	}

	topologyHandler, err := topology.NewTopologyHandler(argsTopologyHandler)
	if err != nil {
		return err
	}

	components.toDharitrIStatusHandler, err = status.NewStatusHandler(ethToDharitrIName, components.statusStorer)
	if err != nil {
		return err
	}

	err = components.metricsHolder.AddStatusHandler(components.toDharitrIStatusHandler)
	if err != nil {
		return err
	}

	timeForTransferExecution := time.Second * time.Duration(args.Configs.GeneralConfig.Eth.IntervalToWaitForTransferInSeconds)

	balanceValidator, err := components.createBalanceValidator(components.baseLogger, components.ethClient)
	if err != nil {
		return err
	}

	argsBridgeExecutor := bridges.ArgsBridgeExecutor{
		Log:                          log,
		TopologyProvider:             topologyHandler,
		DharitrIClient:             components.dharitrIClient,
		PeerChainClient:              components.ethClient,
		StatusHandler:                components.toDharitrIStatusHandler,
		TimeForWaitOnPeerClient:      timeForTransferExecution,
		SignaturesHolder:             disabled.NewDisabledSignaturesHolder(),
		BalanceValidator:             balanceValidator,
		MaxQuorumRetriesOnPeerClient: args.Configs.GeneralConfig.Eth.MaxRetriesOnQuorumReached,
		MaxQuorumRetriesOnDharitrI: args.Configs.GeneralConfig.DharitrI.MaxRetriesOnQuorumReached,
		MaxRestriesOnWasProposed:     args.Configs.GeneralConfig.DharitrI.MaxRetriesOnWasTransferProposed,
	}

	bridge, err := bridges.NewBridgeExecutor(argsBridgeExecutor)
	if err != nil {
		return err
	}

	components.toDharitrIMachineStates, err = ethtodharitri.CreateSteps(bridge)
	if err != nil {
		return err
	}

	return nil
}

func (components *ethDrtBridgeComponents) createEthereumToDharitrIStateMachine() error {
	ethToDharitrIName := components.evmCompatibleChain.PeerChainToDharitrIName()
	log := core.NewLoggerWithIdentifier(logger.GetOrCreate(ethToDharitrIName), ethToDharitrIName)

	argsStateMachine := stateMachine.ArgsStateMachine{
		StateMachineName:     ethToDharitrIName,
		Steps:                components.toDharitrIMachineStates,
		StartStateIdentifier: ethtodharitri.GettingPendingBatchFromPeerChain,
		Log:                  log,
		StatusHandler:        components.toDharitrIStatusHandler,
	}

	var err error
	components.toDharitrIStateMachine, err = stateMachine.NewStateMachine(argsStateMachine)
	if err != nil {
		return err
	}

	argsPollingHandler := polling.ArgsPollingHandler{
		Log:              log,
		Name:             ethToDharitrIName + " State machine",
		PollingInterval:  components.toDharitrIStepDuration,
		PollingWhenError: pollingDurationOnError,
		Executor:         components.toDharitrIStateMachine,
	}

	pollingHandler, err := polling.NewPollingHandler(argsPollingHandler)
	if err != nil {
		return err
	}

	components.addClosableComponent(pollingHandler)
	components.pollingHandlers = append(components.pollingHandlers, pollingHandler)

	return nil
}

func (components *ethDrtBridgeComponents) createDharitrIToEthereumBridge(args ArgsEthereumToDharitrIBridge) error {
	dharitrIToEthName := components.evmCompatibleChain.DharitrIToPeerChainName()
	log := core.NewLoggerWithIdentifier(logger.GetOrCreate(dharitrIToEthName), dharitrIToEthName)

	configs, found := args.Configs.GeneralConfig.StateMachine[dharitrIToEthName]
	if !found {
		return fmt.Errorf("%w for %q", errMissingConfig, dharitrIToEthName)
	}

	components.fromDharitrIStepDuration = time.Duration(configs.StepDurationInMillis) * time.Millisecond
	argsTopologyHandler := topology.ArgsTopologyHandler{
		PublicKeysProvider: components.dharitrIRoleProvider,
		Timer:              components.timer,
		IntervalForLeader:  time.Second * time.Duration(configs.IntervalForLeaderInSeconds),
		AddressBytes:       components.dharitrIRelayerAddress.AddressBytes(),
		Log:                log,
		AddressConverter:   components.addressConverter,
	}

	topologyHandler, err := topology.NewTopologyHandler(argsTopologyHandler)
	if err != nil {
		return err
	}

	components.fromDharitrIStatusHandler, err = status.NewStatusHandler(dharitrIToEthName, components.statusStorer)
	if err != nil {
		return err
	}

	err = components.metricsHolder.AddStatusHandler(components.fromDharitrIStatusHandler)
	if err != nil {
		return err
	}

	timeForWaitOnEthereum := time.Second * time.Duration(args.Configs.GeneralConfig.Eth.IntervalToWaitForTransferInSeconds)

	balanceValidator, err := components.createBalanceValidator(components.baseLogger, components.ethClient)
	if err != nil {
		return err
	}

	argsBridgeExecutor := bridges.ArgsBridgeExecutor{
		Log:                          log,
		TopologyProvider:             topologyHandler,
		DharitrIClient:             components.dharitrIClient,
		PeerChainClient:              components.ethClient,
		StatusHandler:                components.fromDharitrIStatusHandler,
		TimeForWaitOnPeerClient:      timeForWaitOnEthereum,
		SignaturesHolder:             components.toDharitrISignaturesHolder,
		BalanceValidator:             balanceValidator,
		MaxQuorumRetriesOnPeerClient: args.Configs.GeneralConfig.Eth.MaxRetriesOnQuorumReached,
		MaxQuorumRetriesOnDharitrI: args.Configs.GeneralConfig.DharitrI.MaxRetriesOnQuorumReached,
		MaxRestriesOnWasProposed:     args.Configs.GeneralConfig.DharitrI.MaxRetriesOnWasTransferProposed,
	}

	bridge, err := bridges.NewBridgeExecutor(argsBridgeExecutor)
	if err != nil {
		return err
	}

	components.fromDharitrIMachineStates, err = dharitritoeth.CreateSteps(bridge)
	if err != nil {
		return err
	}

	return nil
}

func (components *ethDrtBridgeComponents) createDharitrIToEthereumStateMachine() error {
	dharitrIToEthName := components.evmCompatibleChain.DharitrIToPeerChainName()
	log := core.NewLoggerWithIdentifier(logger.GetOrCreate(dharitrIToEthName), dharitrIToEthName)

	argsStateMachine := stateMachine.ArgsStateMachine{
		StateMachineName:     dharitrIToEthName,
		Steps:                components.fromDharitrIMachineStates,
		StartStateIdentifier: dharitritoeth.GettingPendingBatchFromDharitrI,
		Log:                  log,
		StatusHandler:        components.fromDharitrIStatusHandler,
	}

	var err error
	components.fromDharitrIStateMachine, err = stateMachine.NewStateMachine(argsStateMachine)
	if err != nil {
		return err
	}

	argsPollingHandler := polling.ArgsPollingHandler{
		Log:              log,
		Name:             dharitrIToEthName + " State machine",
		PollingInterval:  components.fromDharitrIStepDuration,
		PollingWhenError: pollingDurationOnError,
		Executor:         components.fromDharitrIStateMachine,
	}

	pollingHandler, err := polling.NewPollingHandler(argsPollingHandler)
	if err != nil {
		return err
	}

	components.addClosableComponent(pollingHandler)
	components.pollingHandlers = append(components.pollingHandlers, pollingHandler)

	return nil
}

// Start will start the bridge
func (components *ethDrtBridgeComponents) Start() error {
	return components.start()
}

// Close will close the bridge
func (components *ethDrtBridgeComponents) Close() error {
	return components.close()
}

// PeerChainRelayerAddress returns the Ethereum's address associated to this relayer
func (components *ethDrtBridgeComponents) PeerChainRelayerAddress() string {
	return components.ethereumRelayerAddress.String()
}
