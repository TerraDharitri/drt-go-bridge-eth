package main

import (
	"fmt"
	"os"
	"os/signal"
	"path"
	"runtime"
	"syscall"
	"time"

	"github.com/block-vision/sui-go-sdk/sui"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/TerraDharitri/drt-go-bridge-eth/clients/ethereum"
	"github.com/TerraDharitri/drt-go-bridge-eth/clients/ethereum/contract"
	"github.com/TerraDharitri/drt-go-bridge-eth/clients/ethereum/wrappers"
	"github.com/TerraDharitri/drt-go-bridge-eth/config"
	"github.com/TerraDharitri/drt-go-bridge-eth/core"
	"github.com/TerraDharitri/drt-go-bridge-eth/factory"
	"github.com/TerraDharitri/drt-go-bridge-eth/p2p"
	"github.com/TerraDharitri/drt-go-bridge-eth/status"
	"github.com/TerraDharitri/drt-go-chain-communication/p2p/libp2p"
	chainCore "github.com/TerraDharitri/drt-go-chain-core/core"
	"github.com/TerraDharitri/drt-go-chain-core/core/check"
	"github.com/TerraDharitri/drt-go-chain-core/data/typeConverters/uint64ByteSlice"
	"github.com/TerraDharitri/drt-go-chain-core/marshal"
	factoryMarshaller "github.com/TerraDharitri/drt-go-chain-core/marshal/factory"
	"github.com/TerraDharitri/drt-go-chain-crypto/signing"
	"github.com/TerraDharitri/drt-go-chain-crypto/signing/secp256k1"
	"github.com/TerraDharitri/drt-go-chain-crypto/signing/secp256k1/singlesig"
	chainFactory "github.com/TerraDharitri/drt-go-chain/cmd/node/factory"
	chainCommon "github.com/TerraDharitri/drt-go-chain/common"
	p2pConfig "github.com/TerraDharitri/drt-go-chain/p2p/config"
	p2pFactory "github.com/TerraDharitri/drt-go-chain/p2p/factory"
	"github.com/TerraDharitri/drt-go-chain/statusHandler"
	"github.com/TerraDharitri/drt-go-chain/statusHandler/persister"
	"github.com/TerraDharitri/drt-go-chain/storage/cache"
	"github.com/TerraDharitri/drt-go-chain/update/disabled"
	logger "github.com/TerraDharitri/drt-go-chain-logger"
	"github.com/TerraDharitri/drt-go-chain-logger/file"
	"github.com/TerraDharitri/drt-go-sdk/blockchain"
	sdkCore "github.com/TerraDharitri/drt-go-sdk/core"
	"github.com/urfave/cli"
)

const (
	filePathPlaceholder      = "[path]"
	defaultLogsPath          = "logs"
	logFilePrefix            = "dharitri-eth-bridge"
	p2pPeerNetworkDiscoverer = "optimized"
	nilListSharderType       = "NilListSharder"
	disabledWatcher          = "disabled"
	dbPath                   = "db"
	timeForBootstrap         = time.Second * 20
	timeBeforeRepeatJoin     = time.Minute * 5
)

var log = logger.GetOrCreate("main")

// appVersion should be populated at build time using ldflags
// Usage examples:
// linux/mac:
//
//	go build -i -v -ldflags="-X main.appVersion=$(git describe --tags --long --dirty)"
//
// windows:
//
//	for /f %i in ('git describe --tags --long --dirty') do set VERS=%i
//	go build -i -v -ldflags="-X main.appVersion=%VERS%"
var appVersion = chainCommon.UnVersionedAppString

func main() {
	app := cli.NewApp()
	app.Name = "Relay CLI app"
	app.Usage = "This is the entry point for the bridge relay"
	app.Flags = getFlags()
	machineID := chainCore.GetAnonymizedMachineID(app.Name)
	app.Version = fmt.Sprintf("%s/%s/%s-%s/%s", appVersion, runtime.Version(), runtime.GOOS, runtime.GOARCH, machineID)
	app.Authors = []cli.Author{
		{
			Name:  "Team Dharitri",
			Email: "contact@dharitri.org",
		},
	}

	app.Action = func(c *cli.Context) error {
		return startRelay(c, app.Version)
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}

func startRelay(ctx *cli.Context, version string) error {
	flagsConfig := getFlagsConfig(ctx)

	fileLogging, errLogger := attachFileLogger(log, flagsConfig)
	if errLogger != nil {
		return errLogger
	}

	log.Info("starting bridge node", "version", version, "pid", os.Getpid())

	err := logger.SetLogLevel(flagsConfig.LogLevel)
	if err != nil {
		return err
	}

	cfg, err := loadConfig(flagsConfig.ConfigurationFile)
	if err != nil {
		return err
	}

	apiRoutesConfig, err := loadApiConfig(flagsConfig.ConfigurationApiFile)
	if err != nil {
		return err
	}
	log.Debug("config", "file", flagsConfig.ConfigurationApiFile)

	if !check.IfNil(fileLogging) {
		timeLogLifeSpan := time.Second * time.Duration(cfg.Logs.LogFileLifeSpanInSec)
		sizeLogLifeSpanInMB := uint64(cfg.Logs.LogFileLifeSpanInMB)
		err = fileLogging.ChangeFileLifeSpan(timeLogLifeSpan, sizeLogLifeSpanInMB)
		if err != nil {
			return err
		}
	}

	dbFullPath := path.Join(flagsConfig.WorkingDir, dbPath)
	statusStorer, err := factory.CreateUnitStorer(cfg.Relayer.StatusMetricsStorage, dbFullPath)
	if err != nil {
		return err
	}

	metricsHolder := status.NewMetricsHolder()

	dharitrIClientStatusHandler, err := status.NewStatusHandler(core.DharitrIClientStatusHandlerName, statusStorer)
	if err != nil {
		return err
	}
	err = metricsHolder.AddStatusHandler(dharitrIClientStatusHandler)
	if err != nil {
		return err
	}

	if len(cfg.DharitrI.NetworkAddress) == 0 {
		return fmt.Errorf("empty DharitrI.NetworkAddress in config file")
	}

	argsProxy := blockchain.ArgsProxy{
		ProxyURL:            cfg.DharitrI.NetworkAddress,
		SameScState:         false,
		ShouldBeSynced:      false,
		FinalityCheck:       cfg.DharitrI.Proxy.FinalityCheck,
		AllowedDeltaToFinal: cfg.DharitrI.Proxy.MaxNoncesDelta,
		CacheExpirationTime: time.Second * time.Duration(cfg.DharitrI.Proxy.CacherExpirationSeconds),
		EntityType:          sdkCore.RestAPIEntityType(cfg.DharitrI.Proxy.RestAPIEntityType),
	}
	proxy, err := blockchain.NewProxy(argsProxy)
	if err != nil {
		return err
	}

	marshaller, err := factoryMarshaller.NewMarshalizer(cfg.Relayer.Marshalizer.Type)
	if err != nil {
		return err
	}

	messenger, err := buildNetMessenger(cfg, marshaller)
	if err != nil {
		return err
	}

	configs := config.Configs{
		GeneralConfig:   cfg,
		ApiRoutesConfig: apiRoutesConfig,
		FlagsConfig:     flagsConfig,
	}

	var appStatusHandlers []chainCore.AppStatusHandler
	statusMetrics := statusHandler.NewStatusMetrics()
	appStatusHandlers = append(appStatusHandlers, statusMetrics)

	persistentHandler, err := persister.NewPersistentStatusHandler(marshaller, uint64ByteSlice.NewBigEndianConverter())
	if err != nil {
		return err
	}
	appStatusHandlers = append(appStatusHandlers, persistentHandler)
	appStatusHandler, err := statusHandler.NewAppStatusFacadeWithHandlers(appStatusHandlers...)
	if err != nil {
		return err
	}

	argsCommonBridge := factory.ArgsBridgeCommon{
		Configs:                       configs,
		StatusStorer:                  statusStorer,
		MetricsHolder:                 metricsHolder,
		Messenger:                     messenger,
		Proxy:                         proxy,
		AppStatusHandler:              appStatusHandler,
		DharitrIClientStatusHandler: dharitrIClientStatusHandler,
	}

	var bridge factory.BridgeComponents
	switch {
	case cfg.Eth.Enabled && !cfg.Sui.Enabled:
		bridge, err = setupEthComponents(argsCommonBridge)
	case cfg.Sui.Enabled && !cfg.Eth.Enabled:
		bridge, err = setupSuiComponents(argsCommonBridge)
	default:
		return fmt.Errorf(
			"invalid configuration: Eth.Enabled=%v, Sui.Enabled=%v; exactly one must be true",
			cfg.Eth.Enabled,
			cfg.Sui.Enabled,
		)
	}
	if err != nil {
		return fmt.Errorf("failed to setup bridge: %w", err)
	}

	webServer, err := factory.StartWebServer(configs, metricsHolder)
	if err != nil {
		return err
	}

	log.Info("Starting relay")

	err = bridge.Start()
	if err != nil {
		return err
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs

	log.Info("application closing, calling Close on all subcomponents...")

	var lastErr error
	err = bridge.Close()
	if err != nil {
		lastErr = err
	}

	err = webServer.Close()
	if err != nil {
		lastErr = err
	}

	return lastErr
}

func setupEthComponents(args factory.ArgsBridgeCommon) (factory.BridgeComponents, error) {
	ethClientStatusHandler, err := status.NewStatusHandler(core.EthClientStatusHandlerName, args.StatusStorer)
	if err != nil {
		return nil, err
	}
	err = args.MetricsHolder.AddStatusHandler(ethClientStatusHandler)
	if err != nil {
		return nil, err
	}

	ethClient, err := ethclient.Dial(args.Configs.GeneralConfig.Eth.NetworkAddress)
	if err != nil {
		return nil, err
	}

	bridgeEthAddress := ethCommon.HexToAddress(args.Configs.GeneralConfig.Eth.MultisigContractAddress)
	multiSigInstance, err := contract.NewBridge(bridgeEthAddress, ethClient)
	if err != nil {
		return nil, err
	}

	safeEthAddress := ethCommon.HexToAddress(args.Configs.GeneralConfig.Eth.SafeContractAddress)
	safeInstance, err := contract.NewERC20Safe(safeEthAddress, ethClient)
	if err != nil {
		return nil, err
	}

	argsContractsHolder := ethereum.ArgsErc20SafeContractsHolder{
		EthClient:              ethClient,
		EthClientStatusHandler: ethClientStatusHandler,
	}
	erc20ContractsHolder, err := ethereum.NewErc20SafeContractsHolder(argsContractsHolder)
	if err != nil {
		return nil, err
	}

	argsClientWrapper := wrappers.ArgsEthereumChainWrapper{
		StatusHandler:    ethClientStatusHandler,
		MultiSigContract: multiSigInstance,
		SafeContract:     safeInstance,
		BlockchainClient: ethClient,
	}

	clientWrapper, err := wrappers.NewEthereumChainWrapper(argsClientWrapper)
	if err != nil {
		return nil, err
	}

	ethDrtBridgeArgs := factory.ArgsEthereumToDharitrIBridge{
		Configs:                       args.Configs,
		Messenger:                     args.Messenger,
		StatusStorer:                  args.StatusStorer,
		Proxy:                         args.Proxy,
		Erc20ContractsHolder:          erc20ContractsHolder,
		ClientWrapper:                 clientWrapper,
		TimeForBootstrap:              timeForBootstrap,
		TimeBeforeRepeatJoin:          timeBeforeRepeatJoin,
		MetricsHolder:                 args.MetricsHolder,
		AppStatusHandler:              args.AppStatusHandler,
		DharitrIClientStatusHandler: args.DharitrIClientStatusHandler,
	}

	return factory.NewEthDrtBridgeComponents(ethDrtBridgeArgs)
}

func setupSuiComponents(args factory.ArgsBridgeCommon) (factory.BridgeComponents, error) {
	suiConfig := args.Configs.GeneralConfig.Sui

	suiClientStatusHandler, err := status.NewStatusHandler(core.SuiClientStatusHandlerName, args.StatusStorer)
	if err != nil {
		return nil, err
	}
	err = args.MetricsHolder.AddStatusHandler(suiClientStatusHandler)
	if err != nil {
		return nil, err
	}

	if len(suiConfig.NetworkAddress) == 0 {
		return nil, fmt.Errorf("empty Sui.NetworkAddress in config file")
	}
	clientWithProxy := sui.NewSuiClient(suiConfig.NetworkAddress)

	suiDrtBridgeArgs := factory.ArgsSuiToDharitrIBridge{
		Configs:                       args.Configs,
		Messenger:                     args.Messenger,
		StatusStorer:                  args.StatusStorer,
		Proxy:                         args.Proxy,
		DharitrIClientStatusHandler: args.DharitrIClientStatusHandler,
		SuiProxy:                      clientWithProxy,
		SuiClientStatusHandler:        suiClientStatusHandler,
		TimeForBootstrap:              timeForBootstrap,
		TimeBeforeRepeatJoin:          timeBeforeRepeatJoin,
		MetricsHolder:                 args.MetricsHolder,
		AppStatusHandler:              args.AppStatusHandler,
	}

	return factory.NewSuiDrtBridgeComponents(suiDrtBridgeArgs)
}

func loadConfig(filepath string) (config.Config, error) {
	cfg := config.Config{}
	err := chainCore.LoadTomlFile(&cfg, filepath)
	if err != nil {
		return config.Config{}, err
	}

	return cfg, nil
}

// LoadApiConfig returns a ApiRoutesConfig by reading the config file provided
func loadApiConfig(filepath string) (config.ApiRoutesConfig, error) {
	cfg := config.ApiRoutesConfig{}
	err := chainCore.LoadTomlFile(&cfg, filepath)
	if err != nil {
		return config.ApiRoutesConfig{}, err
	}

	return cfg, nil
}

func attachFileLogger(log logger.Logger, flagsConfig config.ContextFlagsConfig) (chainFactory.FileLoggingHandler, error) {
	var fileLogging chainFactory.FileLoggingHandler
	var err error
	if flagsConfig.SaveLogFile {
		argsFileLogging := file.ArgsFileLogging{
			WorkingDir:      flagsConfig.WorkingDir,
			DefaultLogsPath: defaultLogsPath,
			LogFilePrefix:   logFilePrefix,
		}
		fileLogging, err = file.NewFileLogging(argsFileLogging)
		if err != nil {
			return nil, fmt.Errorf("%w creating a log file", err)
		}
	}

	err = logger.SetDisplayByteSlice(logger.ToHex)
	log.LogIfError(err)
	logger.ToggleLoggerName(flagsConfig.EnableLogName)
	logLevelFlagValue := flagsConfig.LogLevel
	err = logger.SetLogLevel(logLevelFlagValue)
	if err != nil {
		return nil, err
	}

	if flagsConfig.DisableAnsiColor {
		err = logger.RemoveLogObserver(os.Stdout)
		if err != nil {
			return nil, err
		}

		err = logger.AddLogObserver(os.Stdout, &logger.PlainFormatter{})
		if err != nil {
			return nil, err
		}
	}
	log.Trace("logger updated", "level", logLevelFlagValue, "disable ANSI color", flagsConfig.DisableAnsiColor)

	return fileLogging, nil
}

func buildNetMessenger(cfg config.Config, marshalizer marshal.Marshalizer) (p2p.NetMessenger, error) {
	nodeConfig := p2pConfig.NodeConfig{
		Port:                       cfg.P2P.Port,
		MaximumExpectedPeerCount:   0,
		ThresholdMinConnectedPeers: 0,
		Transports:                 cfg.P2P.Transports,
		ResourceLimiter:            cfg.P2P.ResourceLimiter,
	}
	peerDiscoveryConfig := p2pConfig.KadDhtPeerDiscoveryConfig{
		Enabled:                          true,
		RefreshIntervalInSec:             5,
		ProtocolIDs:                       []string{cfg.P2P.ProtocolID},
		InitialPeerList:                  cfg.P2P.InitialPeerList,
		BucketSize:                       0,
		RoutingTableRefreshIntervalInSec: 300,
		Type:                             p2pPeerNetworkDiscoverer,
	}

	p2pCfg := p2pConfig.P2PConfig{
		Node:                nodeConfig,
		KadDhtPeerDiscovery: peerDiscoveryConfig,
		Sharding: p2pConfig.ShardingConfig{
			TargetPeerCount:         0,
			MaxIntraShardValidators: 0,
			MaxCrossShardValidators: 0,
			MaxIntraShardObservers:  0,
			MaxCrossShardObservers:  0,
			Type:                    nilListSharderType,
		},
	}

	p2pLog := logger.GetOrCreate("p2p")
	topRatedCache, err := cache.NewLRUCache(cfg.PeersRatingConfig.TopRatedCacheCapacity)
	if err != nil {
		return nil, err
	}
	badRatedCache, err := cache.NewLRUCache(cfg.PeersRatingConfig.BadRatedCacheCapacity)
	if err != nil {
		return nil, err
	}
	argsPeersRatingHandler := p2pFactory.ArgPeersRatingHandler{
		TopRatedCache: topRatedCache,
		BadRatedCache: badRatedCache,
		Logger:        p2pLog,
	}
	peersRatingHandler, err := p2pFactory.NewPeersRatingHandler(argsPeersRatingHandler)
	if err != nil {
		return nil, err
	}

	p2pSingleSigner := &singlesig.Secp256k1Signer{}
	p2pKeyGen := signing.NewKeyGenerator(secp256k1.NewSecp256k1())
	p2pPrivKey, _ := p2pKeyGen.GeneratePair()

	args := libp2p.ArgsNetworkMessenger{
		Marshaller:            marshalizer,
		P2pConfig:             p2pCfg,
		SyncTimer:             &libp2p.LocalSyncTimer{},
		PreferredPeersHolder:  disabled.NewPreferredPeersHolder(),
		PeersRatingHandler:    peersRatingHandler,
		ConnectionWatcherType: disabledWatcher,
		P2pPrivateKey:         p2pPrivKey,
		P2pSingleSigner:       p2pSingleSigner,
		P2pKeyGenerator:       p2pKeyGen,
		Logger:                p2pLog,
	}

	return libp2p.NewNetworkMessenger(args)
}
