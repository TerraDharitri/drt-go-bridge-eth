package factory

import (
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/block-vision/sui-go-sdk/signer"
	"github.com/TerraDharitri/drt-go-bridge-eth/bridges"
	"github.com/TerraDharitri/drt-go-bridge-eth/bridges/disabled"
	dharitritoeth "github.com/TerraDharitri/drt-go-bridge-eth/bridges/steps/fromDharitrI"
	ethtodharitri "github.com/TerraDharitri/drt-go-bridge-eth/bridges/steps/toDharitrI"
	"github.com/TerraDharitri/drt-go-bridge-eth/bridges/topology"
	"github.com/TerraDharitri/drt-go-bridge-eth/clients/chain"
	"github.com/TerraDharitri/drt-go-bridge-eth/clients/dharitri"
	mapper "github.com/TerraDharitri/drt-go-bridge-eth/clients/dharitri/mappers/sui"
	roleproviders "github.com/TerraDharitri/drt-go-bridge-eth/clients/roleProviders"
	suiClient "github.com/TerraDharitri/drt-go-bridge-eth/clients/sui"
	"github.com/TerraDharitri/drt-go-bridge-eth/config"
	"github.com/TerraDharitri/drt-go-bridge-eth/core"
	"github.com/TerraDharitri/drt-go-bridge-eth/core/converters"
	"github.com/TerraDharitri/drt-go-bridge-eth/p2p"
	"github.com/TerraDharitri/drt-go-bridge-eth/stateMachine"
	"github.com/TerraDharitri/drt-go-bridge-eth/status"
	chainCore "github.com/TerraDharitri/drt-go-chain-core/core"
	"github.com/TerraDharitri/drt-go-chain-core/core/check"
	logger "github.com/TerraDharitri/drt-go-chain-logger"
	"github.com/TerraDharitri/drt-go-sdk/core/polling"
)

// ArgsSuiToDharitrIBridge is the arguments DTO used for creating a Sui to DharitrI bridge
type ArgsSuiToDharitrIBridge struct {
	Configs                       config.Configs
	Messenger                     p2p.NetMessenger
	StatusStorer                  core.Storer
	Proxy                         dharitri.Proxy
	DharitrIClientStatusHandler core.StatusHandler
	SuiProxy                      suiClient.Proxy
	SuiClientStatusHandler        core.StatusHandler
	TimeForBootstrap              time.Duration
	TimeBeforeRepeatJoin          time.Duration
	MetricsHolder                 core.MetricsHolder
	AppStatusHandler              chainCore.AppStatusHandler
}

type suiDrtBridgeComponents struct {
	*baseBridgeComponents
	chain           chain.Chain
	suiApi          suiClient.Proxy
	suiClient       bridges.PeerChainClient
	suiDataGetter   suiDataGetter
	suiSigner       *signer.Signer
	suiRoleProvider PeerChainRoleProvider
	suiPackageId    string
}

func NewSuiDrtBridgeComponents(args ArgsSuiToDharitrIBridge) (*suiDrtBridgeComponents, error) {
	err := checkArgsSui(args)
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

	components := &suiDrtBridgeComponents{
		baseBridgeComponents: baseComponents,
		chain:                args.Configs.GeneralConfig.Sui.Chain,
		suiApi:               args.SuiProxy,
	}

	err = components.initBaseComponents(commonBridgeArgs)
	if err != nil {
		return nil, err
	}

	err = components.createSuiKeysAndAddresses(args.Configs.GeneralConfig.Sui)
	if err != nil {
		return nil, err
	}

	err = components.createSuiDataGetter(args)
	if err != nil {
		return nil, err
	}

	err = components.createSuiRoleProvider(args)
	if err != nil {
		return nil, err
	}

	err = components.createSuiClient(args)
	if err != nil {
		return nil, err
	}

	err = components.createSuiToDharitrIBridge(args)
	if err != nil {
		return nil, err
	}

	err = components.createSuiToDharitrIStateMachine()
	if err != nil {
		return nil, err
	}

	err = components.createDharitrIToSuiBridge(args)
	if err != nil {
		return nil, err
	}

	err = components.createDharitrIToSuiStateMachine()
	if err != nil {
		return nil, err
	}

	return components, nil
}

func checkArgsSui(args ArgsSuiToDharitrIBridge) error {
	if args.SuiProxy == nil {
		return fmt.Errorf("%w for Sui proxy", errNilProxy)
	}
	if check.IfNil(args.SuiClientStatusHandler) {
		return errNilStatusHandler
	}

	return nil
}

func (components *suiDrtBridgeComponents) initBaseComponents(args ArgsBridgeCommon) error {
	suiToDharitrIName := components.chain.PeerChainToDharitrIName()
	baseLogId := components.chain.BaseLogId()
	components.baseLogger = core.NewLoggerWithIdentifier(logger.GetOrCreate(suiToDharitrIName), baseLogId)

	dataGetterLogId := components.chain.DharitrIDataGetterLogId()
	err := components.createDataGetter(dataGetterLogId)
	if err != nil {
		return err
	}

	clientLogId := components.chain.DharitrIClientLogId()
	err = components.createDharitrIRoleProvider(args, clientLogId)
	if err != nil {
		return err
	}

	roleProviderLogId := components.chain.DharitrIRoleProviderLogId()
	tokensMapper, err := mapper.NewDharitrIToSuiMapper(components.drtDataGetter)
	if err != nil {
		return err
	}
	err = components.createDharitrIClient(args, roleProviderLogId, tokensMapper)
	if err != nil {
		return err
	}

	return nil
}

func (components *suiDrtBridgeComponents) createSuiKeysAndAddresses(suiConfigs config.SuiConfig) error {
	privKey, err := loadPrivateKeyFromFile(suiConfigs.PrivateKeyFile)
	if err != nil {
		fmt.Println("Error loading private key from file:", err)
		return err
	}

	components.suiSigner = signer.NewSigner(privKey)
	components.suiPackageId = suiConfigs.PackageId

	return nil
}

func (components *suiDrtBridgeComponents) createSuiDataGetter(args ArgsSuiToDharitrIBridge) error {
	suiConfig := args.Configs.GeneralConfig.Sui
	suiDataGetterLogId := components.chain.PeerChainDataGetterLogId()
	argsSuiDataGetter := suiClient.ArgsSuiClientDataGetter{
		PackageId:                    components.suiPackageId,
		SafeObjectId:                 suiConfig.SafeObjectId,
		SafeInitialSharedVersion:     suiConfig.SafeObjectInitialSharedVersion,
		BridgeObjectId:               suiConfig.BridgeObjectId,
		BridgeInitialSharedVersion:   suiConfig.BridgeObjectInitialSharedVersion,
		TreasuryObjectId:             suiConfig.TreasuryObjectId,
		TreasuryInitialSharedVersion: suiConfig.TreasuryObjectInitialSharedVersion,
		RelayerAddress:               components.suiSigner.Address,
		Proxy:                        components.suiApi,
		Log:                          core.NewLoggerWithIdentifier(logger.GetOrCreate(suiDataGetterLogId), suiDataGetterLogId),
	}

	var err error
	components.suiDataGetter, err = suiClient.NewSuiClientDataGetter(argsSuiDataGetter)

	return err
}

func (components *suiDrtBridgeComponents) createSuiRoleProvider(args ArgsSuiToDharitrIBridge) error {
	configs := args.Configs.GeneralConfig
	suiRoleProviderLogId := components.chain.PeerChainRoleProviderLogId()
	log := core.NewLoggerWithIdentifier(logger.GetOrCreate(suiRoleProviderLogId), suiRoleProviderLogId)
	argsRoleProvider := roleproviders.ArgsSuiRoleProvider{
		DataGetter: components.suiDataGetter,
		Log:        log,
	}

	var err error
	components.suiRoleProvider, err = roleproviders.NewSuiRoleProvider(argsRoleProvider)
	if err != nil {
		return err
	}

	argsPollingHandler := polling.ArgsPollingHandler{
		Log:              log,
		Name:             string(components.chain) + " role provider",
		PollingInterval:  time.Duration(configs.Relayer.RoleProvider.PollingIntervalInMillis) * time.Millisecond,
		PollingWhenError: pollingDurationOnError,
		Executor:         components.suiRoleProvider,
	}

	pollingHandler, err := polling.NewPollingHandler(argsPollingHandler)
	if err != nil {
		return err
	}

	components.addClosableComponent(pollingHandler)
	components.pollingHandlers = append(components.pollingHandlers, pollingHandler)

	return nil
}

func (components *suiDrtBridgeComponents) createSuiClient(args ArgsSuiToDharitrIBridge) error {
	suiConfig := args.Configs.GeneralConfig.Sui

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

	broadcasterLogId := components.chain.BroadcasterLogId()
	suiToDharitrIName := components.chain.PeerChainToDharitrIName()
	argsBroadcaster := p2p.ArgsBroadcaster{
		Messenger:              args.Messenger,
		Log:                    core.NewLoggerWithIdentifier(logger.GetOrCreate(broadcasterLogId), broadcasterLogId),
		DharitrIRoleProvider: components.dharitrIRoleProvider,
		SignatureProcessor:     components.suiRoleProvider,
		KeyGen:                 keyGen,
		SingleSigner:           singleSigner,
		PrivateKey:             components.dharitrIRelayerPrivateKey,
		Name:                   suiToDharitrIName,
		AntifloodComponents:    antifloodComponents,
	}

	components.broadcaster, err = p2p.NewBroadcaster(argsBroadcaster)
	if err != nil {
		return err
	}

	tokensMapper, err := mapper.NewSuiToDharitrIMapper(components.drtDataGetter)
	if err != nil {
		return err
	}

	signaturesHolder := bridges.NewSignatureHolder()
	components.toDharitrISignaturesHolder = signaturesHolder
	err = components.broadcaster.AddBroadcastClient(signaturesHolder)
	if err != nil {
		return err
	}

	suiClientLogId := components.chain.PeerChainClientLogId()
	suiClientLogger := core.NewLoggerWithIdentifier(logger.GetOrCreate(suiClientLogId), suiClientLogId)

	argsTxHandler := suiClient.ArgsTxHandler{
		Proxy:     components.suiApi,
		Signer:    components.suiSigner,
		GasPrice:  1000,
		GasBudget: 75000000,
		Logger:    suiClientLogger,
	}
	suiTxHandler, err := suiClient.NewTransactionHandler(argsTxHandler)
	if err != nil {
		return err
	}

	argsSuiClient := suiClient.ArgsSuiClient{
		Proxy:                        components.suiApi,
		TxHandler:                    suiTxHandler,
		Log:                          suiClientLogger,
		Signer:                       components.suiSigner,
		PackageId:                    components.suiPackageId,
		SafeObjectId:                 suiConfig.SafeObjectId,
		SafeInitialSharedVersion:     suiConfig.SafeObjectInitialSharedVersion,
		BridgeObjectId:               suiConfig.BridgeObjectId,
		BridgeInitialSharedVersion:   suiConfig.BridgeObjectInitialSharedVersion,
		TreasuryObjectId:             suiConfig.TreasuryObjectId,
		TreasuryInitialSharedVersion: suiConfig.TreasuryObjectInitialSharedVersion,
		Broadcaster:                  components.broadcaster,
		TokensMapper:                 tokensMapper,
		SignatureHolder:              signaturesHolder,
		StatusHandler:                args.SuiClientStatusHandler,
		ClientAvailabilityAllowDelta: suiConfig.ClientAvailabilityAllowDelta,
	}

	components.suiClient, err = suiClient.NewSuiClient(argsSuiClient)

	return err
}

func (components *suiDrtBridgeComponents) createSuiToDharitrIBridge(args ArgsSuiToDharitrIBridge) error {
	suiToDharitrIName := components.chain.PeerChainToDharitrIName()
	log := core.NewLoggerWithIdentifier(logger.GetOrCreate(suiToDharitrIName), suiToDharitrIName)

	configs, found := args.Configs.GeneralConfig.StateMachine[suiToDharitrIName]
	if !found {
		return fmt.Errorf("%w for %q", errMissingConfig, suiToDharitrIName)
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

	components.toDharitrIStatusHandler, err = status.NewStatusHandler(suiToDharitrIName, components.statusStorer)
	if err != nil {
		return err
	}

	err = components.metricsHolder.AddStatusHandler(components.toDharitrIStatusHandler)
	if err != nil {
		return err
	}

	timeForTransferExecution := time.Second * time.Duration(args.Configs.GeneralConfig.Sui.IntervalToWaitForTransferInSeconds)

	balanceValidator, err := components.createBalanceValidator(components.baseLogger, components.suiClient)
	if err != nil {
		return err
	}

	argsBridgeExecutor := bridges.ArgsBridgeExecutor{
		Log:                          log,
		TopologyProvider:             topologyHandler,
		DharitrIClient:             components.dharitrIClient,
		PeerChainClient:              components.suiClient,
		StatusHandler:                components.toDharitrIStatusHandler,
		TimeForWaitOnPeerClient:      timeForTransferExecution,
		SignaturesHolder:             disabled.NewDisabledSignaturesHolder(),
		BalanceValidator:             balanceValidator,
		MaxQuorumRetriesOnPeerClient: args.Configs.GeneralConfig.Sui.MaxRetriesOnQuorumReached,
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

func (components *suiDrtBridgeComponents) createSuiToDharitrIStateMachine() error {
	suiToDharitrIName := components.chain.PeerChainToDharitrIName()
	log := core.NewLoggerWithIdentifier(logger.GetOrCreate(suiToDharitrIName), suiToDharitrIName)

	argsStateMachine := stateMachine.ArgsStateMachine{
		StateMachineName:     suiToDharitrIName,
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
		Name:             suiToDharitrIName + " State machine",
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

func (components *suiDrtBridgeComponents) createDharitrIToSuiBridge(args ArgsSuiToDharitrIBridge) error {
	dharitrIToSuiName := components.chain.DharitrIToPeerChainName()
	log := core.NewLoggerWithIdentifier(logger.GetOrCreate(dharitrIToSuiName), dharitrIToSuiName)

	configs, found := args.Configs.GeneralConfig.StateMachine[dharitrIToSuiName]
	if !found {
		return fmt.Errorf("%w for %q", errMissingConfig, dharitrIToSuiName)
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

	components.fromDharitrIStatusHandler, err = status.NewStatusHandler(dharitrIToSuiName, components.statusStorer)
	if err != nil {
		return err
	}

	err = components.metricsHolder.AddStatusHandler(components.fromDharitrIStatusHandler)
	if err != nil {
		return err
	}

	timeForWaitOnSui := time.Second * time.Duration(args.Configs.GeneralConfig.Sui.IntervalToWaitForTransferInSeconds)

	balanceValidator, err := components.createBalanceValidator(components.baseLogger, components.suiClient)
	if err != nil {
		return err
	}

	argsBridgeExecutor := bridges.ArgsBridgeExecutor{
		Log:                          log,
		TopologyProvider:             topologyHandler,
		DharitrIClient:             components.dharitrIClient,
		PeerChainClient:              components.suiClient,
		StatusHandler:                components.fromDharitrIStatusHandler,
		TimeForWaitOnPeerClient:      timeForWaitOnSui,
		SignaturesHolder:             components.toDharitrISignaturesHolder,
		BalanceValidator:             balanceValidator,
		MaxQuorumRetriesOnPeerClient: args.Configs.GeneralConfig.Sui.MaxRetriesOnQuorumReached,
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

func (components *suiDrtBridgeComponents) createDharitrIToSuiStateMachine() error {
	dharitrIToSuiName := components.chain.DharitrIToPeerChainName()
	log := core.NewLoggerWithIdentifier(logger.GetOrCreate(dharitrIToSuiName), dharitrIToSuiName)

	argsStateMachine := stateMachine.ArgsStateMachine{
		StateMachineName:     dharitrIToSuiName,
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
		Name:             dharitrIToSuiName + " State machine",
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
func (components *suiDrtBridgeComponents) Start() error {
	return components.start()
}

// Close will close the bridge
func (components *suiDrtBridgeComponents) Close() error {
	return components.close()
}

// PeerChainRelayerAddress returns the Sui address associated to this relayer
func (components *suiDrtBridgeComponents) PeerChainRelayerAddress() string {
	return components.suiSigner.Address
}

func loadPrivateKeyFromFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	privKey := converters.TrimWhiteSpaceCharacters(string(data))
	return hex.DecodeString(privKey)
}
