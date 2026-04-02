package factory

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/TerraDharitri/drt-go-bridge-eth/bridges"
	"github.com/TerraDharitri/drt-go-bridge-eth/clients"
	balanceValidatorManagement "github.com/TerraDharitri/drt-go-bridge-eth/clients/balanceValidator"
	"github.com/TerraDharitri/drt-go-bridge-eth/clients/dharitri"
	roleproviders "github.com/TerraDharitri/drt-go-bridge-eth/clients/roleProviders"
	"github.com/TerraDharitri/drt-go-bridge-eth/config"
	"github.com/TerraDharitri/drt-go-bridge-eth/core"
	"github.com/TerraDharitri/drt-go-bridge-eth/core/converters"
	"github.com/TerraDharitri/drt-go-bridge-eth/core/timer"
	"github.com/TerraDharitri/drt-go-bridge-eth/p2p"
	chainCore "github.com/TerraDharitri/drt-go-chain-core/core"
	"github.com/TerraDharitri/drt-go-chain-core/core/check"
	crypto "github.com/TerraDharitri/drt-go-chain-crypto"
	"github.com/TerraDharitri/drt-go-chain-crypto/signing"
	"github.com/TerraDharitri/drt-go-chain-crypto/signing/ed25519"
	"github.com/TerraDharitri/drt-go-chain-crypto/signing/ed25519/singlesig"
	chainConfig "github.com/TerraDharitri/drt-go-chain/config"
	antifloodFactory "github.com/TerraDharitri/drt-go-chain/process/throttle/antiflood/factory"
	logger "github.com/TerraDharitri/drt-go-chain-logger"
	sdkCore "github.com/TerraDharitri/drt-go-sdk/core"
	"github.com/TerraDharitri/drt-go-sdk/core/polling"
	"github.com/TerraDharitri/drt-go-sdk/data"
	"github.com/TerraDharitri/drt-go-sdk/interactors"
)

const (
	minTimeForBootstrap     = time.Millisecond * 100
	minTimeBeforeRepeatJoin = time.Second * 30
	pollingDurationOnError  = time.Second * 5
)

var suite = ed25519.NewEd25519()
var keyGen = signing.NewKeyGenerator(suite)
var singleSigner = &singlesig.Ed25519Signer{}

type ArgsBridgeCommon struct {
	Configs                       config.Configs
	Messenger                     p2p.NetMessenger
	StatusStorer                  core.Storer
	Proxy                         dharitri.Proxy
	DharitrIClientStatusHandler core.StatusHandler
	TimeForBootstrap              time.Duration
	TimeBeforeRepeatJoin          time.Duration
	MetricsHolder                 core.MetricsHolder
	AppStatusHandler              chainCore.AppStatusHandler
}

type baseBridgeComponents struct {
	baseLogger                        logger.Logger
	messenger                         p2p.NetMessenger
	statusStorer                      core.Storer
	dharitrIClient                  bridges.DharitrIClient
	dharitrIMultisigContractAddress sdkCore.AddressHandler
	dharitrISafeContractAddress     sdkCore.AddressHandler
	dharitrIRelayerPrivateKey       crypto.PrivateKey
	dharitrIRelayerAddress          sdkCore.AddressHandler
	drtDataGetter                      dataGetter
	proxy                             dharitri.Proxy
	dharitrIRoleProvider            DharitrIRoleProvider
	broadcaster                       Broadcaster
	timer                             core.Timer
	timeForBootstrap                  time.Duration
	metricsHolder                     core.MetricsHolder
	addressConverter                  core.AddressConverter

	toDharitrIMachineStates    core.MachineStates
	toDharitrIStepDuration     time.Duration
	toDharitrIStatusHandler    core.StatusHandler
	toDharitrIStateMachine     StateMachine
	toDharitrISignaturesHolder bridges.SignaturesHolder

	fromDharitrIMachineStates core.MachineStates
	fromDharitrIStepDuration  time.Duration
	fromDharitrIStatusHandler core.StatusHandler
	fromDharitrIStateMachine  StateMachine

	mutClosableHandlers sync.RWMutex
	closableHandlers    []io.Closer

	pollingHandlers []PollingHandler

	timeBeforeRepeatJoin time.Duration
	cancelFunc           func()
	appStatusHandler     chainCore.AppStatusHandler
}

// NewBaseComponents creates a base component holder
func NewBaseComponents(args ArgsBridgeCommon) (*baseBridgeComponents, error) {
	err := checkCommonArgs(args)
	if err != nil {
		return nil, err
	}

	components := &baseBridgeComponents{
		messenger:            args.Messenger,
		statusStorer:         args.StatusStorer,
		closableHandlers:     make([]io.Closer, 0),
		proxy:                args.Proxy,
		timer:                timer.NewNTPTimer(),
		timeForBootstrap:     args.TimeForBootstrap,
		timeBeforeRepeatJoin: args.TimeBeforeRepeatJoin,
		metricsHolder:        args.MetricsHolder,
		appStatusHandler:     args.AppStatusHandler,
	}

	addressConverter, err := converters.NewAddressConverter()
	if err != nil {
		return nil, clients.ErrNilAddressConverter
	}
	components.addressConverter = addressConverter

	components.addClosableComponent(components.timer)

	err = components.createDharitrIKeysAndAddresses(args.Configs.GeneralConfig.DharitrI)
	if err != nil {
		return nil, err
	}

	return components, nil
}

func checkCommonArgs(args ArgsBridgeCommon) error {
	if check.IfNil(args.Proxy) {
		return fmt.Errorf("%w for DharitrI proxy", errNilProxy)
	}
	if check.IfNil(args.Messenger) {
		return errNilMessenger
	}
	if check.IfNil(args.StatusStorer) {
		return errNilStatusStorer
	}
	if args.TimeForBootstrap < minTimeForBootstrap {
		return fmt.Errorf("%w for TimeForBootstrap, received: %v, minimum: %v", errInvalidValue, args.TimeForBootstrap, minTimeForBootstrap)
	}
	if args.TimeBeforeRepeatJoin < minTimeBeforeRepeatJoin {
		return fmt.Errorf("%w for TimeBeforeRepeatJoin, received: %v, minimum: %v", errInvalidValue, args.TimeBeforeRepeatJoin, minTimeBeforeRepeatJoin)
	}
	if check.IfNil(args.MetricsHolder) {
		return errNilMetricsHolder
	}
	if check.IfNil(args.AppStatusHandler) {
		return errNilStatusHandler
	}

	return nil
}

func (components *baseBridgeComponents) addClosableComponent(closable io.Closer) {
	components.mutClosableHandlers.Lock()
	components.closableHandlers = append(components.closableHandlers, closable)
	components.mutClosableHandlers.Unlock()
}

func (components *baseBridgeComponents) createDharitrIKeysAndAddresses(chainConfigs config.DharitrIConfig) error {
	wallet := interactors.NewWallet()
	dharitrIPrivateKeyBytes, err := wallet.LoadPrivateKeyFromPemFile(chainConfigs.PrivateKeyFile)
	if err != nil {
		return err
	}

	components.dharitrIRelayerPrivateKey, err = keyGen.PrivateKeyFromByteArray(dharitrIPrivateKeyBytes)
	if err != nil {
		return err
	}

	components.dharitrIRelayerAddress, err = wallet.GetAddressFromPrivateKey(dharitrIPrivateKeyBytes)
	if err != nil {
		return err
	}

	components.dharitrIMultisigContractAddress, err = data.NewAddressFromBech32String(chainConfigs.MultisigContractAddress)
	if err != nil {
		return fmt.Errorf("%w for chainConfigs.MultisigContractAddress", err)
	}

	components.dharitrISafeContractAddress, err = data.NewAddressFromBech32String(chainConfigs.SafeContractAddress)
	if err != nil {
		return fmt.Errorf("%w for chainConfigs.SafeContractAddress", err)
	}

	return nil
}

func (components *baseBridgeComponents) createDataGetter(logId string) error {
	argsDRTClientDataGetter := dharitri.ArgsDRTClientDataGetter{
		MultisigContractAddress: components.dharitrIMultisigContractAddress,
		SafeContractAddress:     components.dharitrISafeContractAddress,
		RelayerAddress:          components.dharitrIRelayerAddress,
		Proxy:                   components.proxy,
		Log:                     core.NewLoggerWithIdentifier(logger.GetOrCreate(logId), logId),
	}

	var err error
	components.drtDataGetter, err = dharitri.NewDRTClientDataGetter(argsDRTClientDataGetter)

	return err
}

func (components *baseBridgeComponents) createDharitrIClient(args ArgsBridgeCommon, logId string, tokensMapper dharitri.TokensMapper) error {
	chainConfigs := args.Configs.GeneralConfig.DharitrI

	clientArgs := dharitri.ClientArgs{
		GasMapConfig:                 chainConfigs.GasMap,
		Proxy:                        args.Proxy,
		Log:                          core.NewLoggerWithIdentifier(logger.GetOrCreate(logId), logId),
		RelayerPrivateKey:            components.dharitrIRelayerPrivateKey,
		MultisigContractAddress:      components.dharitrIMultisigContractAddress,
		SafeContractAddress:          components.dharitrISafeContractAddress,
		IntervalToResendTxsInSeconds: chainConfigs.IntervalToResendTxsInSeconds,
		TokensMapper:                 tokensMapper,
		RoleProvider:                 components.dharitrIRoleProvider,
		StatusHandler:                args.DharitrIClientStatusHandler,
		ClientAvailabilityAllowDelta: chainConfigs.ClientAvailabilityAllowDelta,
	}

	var err error
	components.dharitrIClient, err = dharitri.NewClient(clientArgs)
	components.addClosableComponent(components.dharitrIClient)

	return err
}

func (components *baseBridgeComponents) createDharitrIRoleProvider(args ArgsBridgeCommon, logId string) error {
	configs := args.Configs.GeneralConfig
	log := core.NewLoggerWithIdentifier(logger.GetOrCreate(logId), logId)

	argsRoleProvider := roleproviders.ArgsDharitrIRoleProvider{
		DataGetter: components.drtDataGetter,
		Log:        log,
	}

	var err error
	components.dharitrIRoleProvider, err = roleproviders.NewDharitrIRoleProvider(argsRoleProvider)
	if err != nil {
		return err
	}

	argsPollingHandler := polling.ArgsPollingHandler{
		Log:              log,
		Name:             "DharitrI role provider",
		PollingInterval:  time.Duration(configs.Relayer.RoleProvider.PollingIntervalInMillis) * time.Millisecond,
		PollingWhenError: pollingDurationOnError,
		Executor:         components.dharitrIRoleProvider,
	}

	pollingHandler, err := polling.NewPollingHandler(argsPollingHandler)
	if err != nil {
		return err
	}

	components.addClosableComponent(pollingHandler)
	components.pollingHandlers = append(components.pollingHandlers, pollingHandler)

	return nil
}

func (components *baseBridgeComponents) createAntifloodComponents(antifloodConfig chainConfig.AntifloodConfig) (*antifloodFactory.AntiFloodComponents, error) {
	var err error
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer func() {
		if err != nil {
			cancelFunc()
		}
	}()

	cfg := chainConfig.Config{
		Antiflood: antifloodConfig,
	}
	antiFloodComponents, err := antifloodFactory.NewP2PAntiFloodComponents(ctx, cfg, components.appStatusHandler, components.messenger.ID())
	if err != nil {
		return nil, err
	}
	return antiFloodComponents, nil
}

func (components *baseBridgeComponents) createBalanceValidator(logger logger.Logger, client bridges.PeerChainClient) (bridges.BalanceValidator, error) {
	argsBalanceValidator := balanceValidatorManagement.ArgsBalanceValidator{
		Log:              logger,
		DharitrIClient: components.dharitrIClient,
		PeerChainClient:  client,
	}

	return balanceValidatorManagement.NewBalanceValidator(argsBalanceValidator)
}

func (components *baseBridgeComponents) startPollingHandlers() error {
	for _, pollingHandler := range components.pollingHandlers {
		err := pollingHandler.StartProcessingLoop()
		if err != nil {
			return err
		}
	}

	return nil
}

func (components *baseBridgeComponents) startBroadcastJoinRetriesLoop(ctx context.Context) {
	broadcastTimer := time.NewTimer(components.timeBeforeRepeatJoin)
	defer broadcastTimer.Stop()

	for {
		broadcastTimer.Reset(components.timeBeforeRepeatJoin)

		select {
		case <-broadcastTimer.C:
			components.baseLogger.Info("broadcast again join topic")
			components.broadcaster.BroadcastJoinTopic()
		case <-ctx.Done():
			components.baseLogger.Info("closing broadcast join topic loop")
			return

		}
	}
}

func (components *baseBridgeComponents) start() error {
	err := components.messenger.Bootstrap()
	if err != nil {
		return err
	}

	components.baseLogger.Info("waiting for p2p bootstrap", "time", components.timeForBootstrap)
	time.Sleep(components.timeForBootstrap)

	err = components.broadcaster.RegisterOnTopics()
	if err != nil {
		return err
	}

	components.broadcaster.BroadcastJoinTopic()

	err = components.startPollingHandlers()
	if err != nil {
		return err
	}

	var ctx context.Context
	ctx, components.cancelFunc = context.WithCancel(context.Background())
	go components.startBroadcastJoinRetriesLoop(ctx)

	return nil
}

func (components *baseBridgeComponents) close() error {
	components.mutClosableHandlers.RLock()
	defer components.mutClosableHandlers.RUnlock()

	if components.cancelFunc != nil {
		components.cancelFunc()
	}

	var lastError error
	for _, closable := range components.closableHandlers {
		if closable == nil {
			components.baseLogger.Warn("programming error, nil closable component")
			continue
		}

		err := closable.Close()
		if err != nil {
			lastError = err

			components.baseLogger.Error("error closing component", "error", err)
		}
	}

	return lastError
}

// DharitrIRelayerAddress returns the DharitrI's address associated to this relayer
func (components *baseBridgeComponents) DharitrIRelayerAddress() sdkCore.AddressHandler {
	return components.dharitrIRelayerAddress
}
