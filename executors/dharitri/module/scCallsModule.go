package module

import (
	"time"

	"github.com/TerraDharitri/drt-go-bridge-eth/config"
	"github.com/TerraDharitri/drt-go-bridge-eth/executors/dharitri"
	"github.com/TerraDharitri/drt-go-bridge-eth/executors/dharitri/filters"
	"github.com/TerraDharitri/drt-go-bridge-eth/parsers"
	"github.com/TerraDharitri/drt-go-chain-crypto/signing"
	"github.com/TerraDharitri/drt-go-chain-crypto/signing/ed25519"
	"github.com/TerraDharitri/drt-go-chain-crypto/signing/ed25519/singlesig"
	logger "github.com/TerraDharitri/drt-go-chain-logger"
	"github.com/TerraDharitri/drt-go-sdk/blockchain"
	sdkCore "github.com/TerraDharitri/drt-go-sdk/core"
	"github.com/TerraDharitri/drt-go-sdk/core/polling"
	"github.com/TerraDharitri/drt-go-sdk/interactors"
	"github.com/TerraDharitri/drt-go-sdk/interactors/nonceHandlerV2"
)

var suite = ed25519.NewEd25519()
var keyGen = signing.NewKeyGenerator(suite)
var singleSigner = &singlesig.Ed25519Signer{}

type scCallsModule struct {
	nonceTxsHandler  nonceTransactionsHandler
	pollingHandler   pollingHandler
	executorInstance executor
}

// NewScCallsModule creates a starts a new scCallsModule instance
func NewScCallsModule(cfg config.ScCallsModuleConfig, log logger.Logger, chCloseApp chan struct{}) (*scCallsModule, error) {
	filter, err := filters.NewPendingOperationFilter(cfg.Filter, log)
	if err != nil {
		return nil, err
	}

	argsProxy := blockchain.ArgsProxy{
		ProxyURL:            cfg.NetworkAddress,
		SameScState:         false,
		ShouldBeSynced:      false,
		FinalityCheck:       cfg.ProxyFinalityCheck,
		AllowedDeltaToFinal: cfg.ProxyMaxNoncesDelta,
		CacheExpirationTime: time.Second * time.Duration(cfg.ProxyCacherExpirationSeconds),
		EntityType:          sdkCore.RestAPIEntityType(cfg.ProxyRestAPIEntityType),
	}

	proxy, err := blockchain.NewProxy(argsProxy)
	if err != nil {
		return nil, err
	}

	module := &scCallsModule{}

	argNonceHandler := nonceHandlerV2.ArgsNonceTransactionsHandlerV2{
		Proxy:            proxy,
		IntervalToResend: time.Second * time.Duration(cfg.IntervalToResendTxsInSeconds),
	}
	module.nonceTxsHandler, err = nonceHandlerV2.NewNonceTransactionHandlerV2(argNonceHandler)
	if err != nil {
		return nil, err
	}

	wallet := interactors.NewWallet()
	dharitrIPrivateKeyBytes, err := wallet.LoadPrivateKeyFromPemFile(cfg.PrivateKeyFile)
	if err != nil {
		return nil, err
	}

	privateKey, err := keyGen.PrivateKeyFromByteArray(dharitrIPrivateKeyBytes)
	if err != nil {
		return nil, err
	}

	argsExecutor := dharitri.ArgsScCallExecutor{
		ScProxyBech32Address:            cfg.ScProxyBech32Address,
		Proxy:                           proxy,
		Codec:                           &parsers.DharitriCodec{},
		Filter:                          filter,
		Log:                             log,
		ExtraGasToExecute:               cfg.ExtraGasToExecute,
		MaxGasLimitToUse:                cfg.MaxGasLimitToUse,
		GasLimitForOutOfGasTransactions: cfg.GasLimitForOutOfGasTransactions,
		NonceTxHandler:                  module.nonceTxsHandler,
		PrivateKey:                      privateKey,
		SingleSigner:                    singleSigner,
		CloseAppChan:                    chCloseApp,
		TransactionChecks:               cfg.TransactionChecks,
	}
	module.executorInstance, err = dharitri.NewScCallExecutor(argsExecutor)
	if err != nil {
		return nil, err
	}

	argsPollingHandler := polling.ArgsPollingHandler{
		Log:              log,
		Name:             "DharitrI SC calls",
		PollingInterval:  time.Duration(cfg.PollingIntervalInMillis) * time.Millisecond,
		PollingWhenError: time.Duration(cfg.PollingIntervalInMillis) * time.Millisecond,
		Executor:         module.executorInstance,
	}

	module.pollingHandler, err = polling.NewPollingHandler(argsPollingHandler)
	if err != nil {
		return nil, err
	}

	err = module.pollingHandler.StartProcessingLoop()
	if err != nil {
		return nil, err
	}

	return module, nil
}

// GetNumSentTransaction returns the total sent transactions
func (module *scCallsModule) GetNumSentTransaction() uint32 {
	return module.executorInstance.GetNumSentTransaction()
}

// Close closes any components started
func (module *scCallsModule) Close() error {
	errPollingHandler := module.pollingHandler.Close()
	errNonceTxsHandler := module.nonceTxsHandler.Close()

	if errPollingHandler != nil {
		return errPollingHandler
	}
	return errNonceTxsHandler
}
