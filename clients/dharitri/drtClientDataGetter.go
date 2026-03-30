package dharitri

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"sync"

	"github.com/TerraDharitri/drt-go-bridge-eth/clients"
	bridgeCore "github.com/TerraDharitri/drt-go-bridge-eth/core"
	"github.com/TerraDharitri/drt-go-bridge-eth/errors"
	"github.com/TerraDharitri/drt-go-chain-core/core/check"
	logger "github.com/TerraDharitri/drt-go-chain-logger"
	"github.com/TerraDharitri/drt-go-sdk/builders"
	"github.com/TerraDharitri/drt-go-sdk/core"
	"github.com/TerraDharitri/drt-go-sdk/data"
)

const (
	okCodeAfterExecution                                      = "ok"
	internalError                                             = "internal error"
	getCurrentTxBatchFuncName                                 = "getCurrentTxBatch"
	getBatchFuncName                                          = "getBatch"
	wasTransferActionProposedFuncName                         = "wasTransferActionProposed"
	wasActionExecutedFuncName                                 = "wasActionExecuted"
	getActionIdForTransferBatchFuncName                       = "getActionIdForTransferBatch"
	wasSetCurrentTransactionBatchStatusActionProposedFuncName = "wasSetCurrentTransactionBatchStatusActionProposed"
	getStatusesAfterExecutionFuncName                         = "getStatusesAfterExecution"
	getActionIdForSetCurrentTransactionBatchStatusFuncName    = "getActionIdForSetCurrentTransactionBatchStatus"
	getTokenIdForErc20AddressFuncName                         = "getTokenIdForErc20Address"
	getErc20AddressForTokenIdFuncName                         = "getErc20AddressForTokenId"
	quorumReachedFuncName                                     = "quorumReached"
	getLastExecutedEthBatchIdFuncName                         = "getLastExecutedEthBatchId"
	getLastExecutedEthTxId                                    = "getLastExecutedEthTxId"
	signedFuncName                                            = "signed"
	getAllStakedRelayersFuncName                              = "getAllStakedRelayers"
	isPausedFuncName                                          = "isPaused"
	isMintBurnTokenFuncName                                   = "isMintBurnToken"
	isNativeTokenFuncName                                     = "isNativeToken"
	getTotalBalances                                          = "getTotalBalances"
	getMintBalances                                           = "getMintBalances"
	getBurnBalances                                           = "getBurnBalances"
	getAllKnownTokens                                         = "getAllKnownTokens"
	getLastBatchId                                            = "getLastBatchId"
)

// ArgsDRTClientDataGetter is the arguments DTO used in the NewDRTClientDataGetter constructor
type ArgsDRTClientDataGetter struct {
	MultisigContractAddress core.AddressHandler
	SafeContractAddress     core.AddressHandler
	RelayerAddress          core.AddressHandler
	Proxy                   Proxy
	Log                     logger.Logger
}

type drtClientDataGetter struct {
	multisigContractAddress       core.AddressHandler
	safeContractAddress           core.AddressHandler
	bech32MultisigContractAddress string
	relayerAddress                core.AddressHandler
	proxy                         Proxy
	log                           logger.Logger
	mutNodeStatus                 sync.Mutex
	wasShardIDFetched             bool
	shardID                       uint32
}

// NewDRTClientDataGetter creates a new instance of the dataGetter type
func NewDRTClientDataGetter(args ArgsDRTClientDataGetter) (*drtClientDataGetter, error) {
	if check.IfNil(args.Log) {
		return nil, errNilLogger
	}
	if check.IfNil(args.Proxy) {
		return nil, errNilProxy
	}
	if check.IfNil(args.RelayerAddress) {
		return nil, fmt.Errorf("%w for the RelayerAddress argument", errNilAddressHandler)
	}
	if check.IfNil(args.MultisigContractAddress) {
		return nil, fmt.Errorf("%w for the MultisigContractAddress argument", errNilAddressHandler)
	}
	if check.IfNil(args.SafeContractAddress) {
		return nil, fmt.Errorf("%w for the SafeContractAddress argument", errNilAddressHandler)
	}
	bech32Address, err := args.MultisigContractAddress.AddressAsBech32String()
	if err != nil {
		return nil, fmt.Errorf("%w for %x", err, args.MultisigContractAddress.AddressBytes())
	}

	return &drtClientDataGetter{
		multisigContractAddress:       args.MultisigContractAddress,
		safeContractAddress:           args.SafeContractAddress,
		bech32MultisigContractAddress: bech32Address,
		relayerAddress:                args.RelayerAddress,
		proxy:                         args.Proxy,
		log:                           args.Log,
	}, nil
}

// ExecuteQueryReturningBytes will try to execute the provided query and return the result as slice of byte slices
func (dataGetter *drtClientDataGetter) ExecuteQueryReturningBytes(ctx context.Context, request *data.VmValueRequest) ([][]byte, error) {
	if request == nil {
		return nil, errNilRequest
	}

	response, err := dataGetter.proxy.ExecuteVMQuery(ctx, request)
	if err != nil {
		dataGetter.log.Error("got error on VMQuery", "FuncName", request.FuncName,
			"Args", request.Args, "SC address", request.Address, "Caller", request.CallerAddr, "error", err)
		return nil, err
	}
	dataGetter.log.Debug("executed VMQuery", "FuncName", request.FuncName,
		"Args", request.Args, "SC address", request.Address, "Caller", request.CallerAddr,
		"response.ReturnCode", response.Data.ReturnCode,
		"response.ReturnData", fmt.Sprintf("%+v", response.Data.ReturnData))
	if response.Data.ReturnCode != okCodeAfterExecution {
		return nil, errors.NewQueryResponseError(
			response.Data.ReturnCode,
			response.Data.ReturnMessage,
			request.FuncName,
			request.Address,
			request.Args...,
		)
	}
	return response.Data.ReturnData, nil
}

// GetCurrentNonce will get from the shard containing the multisig contract the latest block's nonce
func (dataGetter *drtClientDataGetter) GetCurrentNonce(ctx context.Context) (uint64, error) {
	shardID, err := dataGetter.getShardID(ctx)
	if err != nil {
		return 0, err
	}

	nodeStatus, err := dataGetter.proxy.GetNetworkStatus(ctx, shardID)
	if err != nil {
		return 0, err
	}
	if nodeStatus == nil {
		return 0, errNilNodeStatusResponse
	}

	return nodeStatus.Nonce, nil
}

func (dataGetter *drtClientDataGetter) getShardID(ctx context.Context) (uint32, error) {
	dataGetter.mutNodeStatus.Lock()
	defer dataGetter.mutNodeStatus.Unlock()

	if dataGetter.wasShardIDFetched {
		return dataGetter.shardID, nil
	}

	var err error
	dataGetter.shardID, err = dataGetter.proxy.GetShardOfAddress(ctx, dataGetter.bech32MultisigContractAddress)
	if err == nil {
		dataGetter.wasShardIDFetched = true
	}

	return dataGetter.shardID, err
}

// ExecuteQueryReturningBool will try to execute the provided query and return the result as bool
func (dataGetter *drtClientDataGetter) ExecuteQueryReturningBool(ctx context.Context, request *data.VmValueRequest) (bool, error) {
	response, err := dataGetter.ExecuteQueryReturningBytes(ctx, request)
	if err != nil {
		return false, err
	}

	if len(response) == 0 {
		return false, nil
	}

	return dataGetter.parseBool(response[0], request.FuncName, request.Address, request.Args...)
}

func (dataGetter *drtClientDataGetter) parseBool(buff []byte, funcName string, address string, args ...string) (bool, error) {
	if len(buff) == 0 {
		return false, nil
	}

	result, err := strconv.ParseBool(fmt.Sprintf("%d", buff[0]))
	if err != nil {
		return false, errors.NewQueryResponseError(
			internalError,
			fmt.Sprintf("error converting the received bytes to bool, %s", err.Error()),
			funcName,
			address,
			args...,
		)
	}

	return result, nil
}

// ExecuteQueryReturningUint64 will try to execute the provided query and return the result as uint64
func (dataGetter *drtClientDataGetter) ExecuteQueryReturningUint64(ctx context.Context, request *data.VmValueRequest) (uint64, error) {
	response, err := dataGetter.ExecuteQueryReturningBytes(ctx, request)
	if err != nil {
		return 0, err
	}

	if len(response) == 0 {
		return 0, nil
	}
	if len(response[0]) == 0 {
		return 0, nil
	}

	num, err := parseUInt64FromByteSlice(response[0])
	if err != nil {
		return 0, errors.NewQueryResponseError(
			internalError,
			err.Error(),
			request.FuncName,
			request.Address,
			request.Args...,
		)
	}

	return num, nil
}

// ExecuteQueryReturningBigInt will try to execute the provided query and return the result as big.Int
func (dataGetter *drtClientDataGetter) ExecuteQueryReturningBigInt(ctx context.Context, request *data.VmValueRequest) (*big.Int, error) {
	response, err := dataGetter.ExecuteQueryReturningBytes(ctx, request)
	if err != nil {
		return nil, err
	}

	if len(response) == 0 {
		return big.NewInt(0), nil
	}
	if len(response[0]) == 0 {
		return big.NewInt(0), nil
	}

	num := big.NewInt(0).SetBytes(response[0])
	return num, nil
}

func parseUInt64FromByteSlice(bytes []byte) (uint64, error) {
	num := big.NewInt(0).SetBytes(bytes)
	if !num.IsUint64() {
		return 0, errNotUint64Bytes
	}

	return num.Uint64(), nil
}

func (dataGetter *drtClientDataGetter) executeQueryFromBuilder(ctx context.Context, builder builders.VMQueryBuilder) ([][]byte, error) {
	vmValuesRequest, err := builder.ToVmValueRequest()
	if err != nil {
		return nil, err
	}

	return dataGetter.ExecuteQueryReturningBytes(ctx, vmValuesRequest)
}

func (dataGetter *drtClientDataGetter) executeQueryUint64FromBuilder(ctx context.Context, builder builders.VMQueryBuilder) (uint64, error) {
	vmValuesRequest, err := builder.ToVmValueRequest()
	if err != nil {
		return 0, err
	}

	return dataGetter.ExecuteQueryReturningUint64(ctx, vmValuesRequest)
}

func (dataGetter *drtClientDataGetter) executeQueryBigIntFromBuilder(ctx context.Context, builder builders.VMQueryBuilder) (*big.Int, error) {
	vmValuesRequest, err := builder.ToVmValueRequest()
	if err != nil {
		return nil, err
	}

	return dataGetter.ExecuteQueryReturningBigInt(ctx, vmValuesRequest)
}

func (dataGetter *drtClientDataGetter) executeQueryBoolFromBuilder(ctx context.Context, builder builders.VMQueryBuilder) (bool, error) {
	vmValuesRequest, err := builder.ToVmValueRequest()
	if err != nil {
		return false, err
	}

	return dataGetter.ExecuteQueryReturningBool(ctx, vmValuesRequest)
}

func (dataGetter *drtClientDataGetter) createMultisigDefaultVmQueryBuilder() builders.VMQueryBuilder {
	return builders.NewVMQueryBuilder().Address(dataGetter.multisigContractAddress).CallerAddress(dataGetter.relayerAddress)
}

func (dataGetter *drtClientDataGetter) createSafeDefaultVmQueryBuilder() builders.VMQueryBuilder {
	return builders.NewVMQueryBuilder().Address(dataGetter.safeContractAddress).CallerAddress(dataGetter.relayerAddress)
}

// GetCurrentBatchAsDataBytes will assemble a builder and query the proxy for the current pending batch
func (dataGetter *drtClientDataGetter) GetCurrentBatchAsDataBytes(ctx context.Context) ([][]byte, error) {
	builder := dataGetter.createMultisigDefaultVmQueryBuilder()
	builder.Function(getCurrentTxBatchFuncName)

	return dataGetter.executeQueryFromBuilder(ctx, builder)
}

// GetBatchAsDataBytes will assemble a builder and query the proxy for the batch info
func (dataGetter *drtClientDataGetter) GetBatchAsDataBytes(ctx context.Context, batchID uint64) ([][]byte, error) {
	builder := dataGetter.createMultisigDefaultVmQueryBuilder()
	builder.Function(getBatchFuncName)
	builder.ArgInt64(int64(batchID))

	return dataGetter.executeQueryFromBuilder(ctx, builder)
}

// GetTokenIdForErc20Address will assemble a builder and query the proxy for a token id given a specific erc20 address
func (dataGetter *drtClientDataGetter) GetTokenIdForErc20Address(ctx context.Context, erc20Address []byte) ([][]byte, error) {
	builder := dataGetter.createMultisigDefaultVmQueryBuilder()
	builder.Function(getTokenIdForErc20AddressFuncName)
	builder.ArgBytes(erc20Address)

	return dataGetter.executeQueryFromBuilder(ctx, builder)
}

// GetERC20AddressForTokenId will assemble a builder and query the proxy for an erc20 address given a specific token id
func (dataGetter *drtClientDataGetter) GetERC20AddressForTokenId(ctx context.Context, tokenId []byte) ([][]byte, error) {
	builder := dataGetter.createMultisigDefaultVmQueryBuilder()
	builder.Function(getErc20AddressForTokenIdFuncName)
	builder.ArgBytes(tokenId)
	return dataGetter.executeQueryFromBuilder(ctx, builder)
}

// GetTokenIdForSuiCoin will assemble a builder and query the proxy for a token id given a specific sui coin
func (dataGetter *drtClientDataGetter) GetTokenIdForSuiCoin(ctx context.Context, suiCoin []byte) ([][]byte, error) {
	builder := dataGetter.createMultisigDefaultVmQueryBuilder()
	builder.Function(getTokenIdForErc20AddressFuncName)
	builder.ArgBytes(suiCoin)

	return dataGetter.executeQueryFromBuilder(ctx, builder)
}

// GetSuiCoinForTokenId will assemble a builder and query the proxy for a sui coin given a specific token id
func (dataGetter *drtClientDataGetter) GetSuiCoinForTokenId(ctx context.Context, tokenId []byte) ([][]byte, error) {
	builder := dataGetter.createMultisigDefaultVmQueryBuilder()
	builder.Function(getErc20AddressForTokenIdFuncName)
	builder.ArgBytes(tokenId)
	return dataGetter.executeQueryFromBuilder(ctx, builder)
}

// WasProposedTransfer returns true if the transfer action proposed was triggered
func (dataGetter *drtClientDataGetter) WasProposedTransfer(ctx context.Context, batch *bridgeCore.TransferBatch) (bool, error) {
	if batch == nil {
		return false, clients.ErrNilBatch
	}

	builder := dataGetter.createMultisigDefaultVmQueryBuilder()
	builder.Function(wasTransferActionProposedFuncName).ArgInt64(int64(batch.ID))
	dataGetter.addBatchInfo(builder, batch)

	return dataGetter.executeQueryBoolFromBuilder(ctx, builder)
}

// WasExecuted returns true if the provided actionID was executed or not
func (dataGetter *drtClientDataGetter) WasExecuted(ctx context.Context, actionID uint64) (bool, error) {
	builder := dataGetter.createMultisigDefaultVmQueryBuilder()
	builder.Function(wasActionExecutedFuncName).ArgInt64(int64(actionID))

	return dataGetter.executeQueryBoolFromBuilder(ctx, builder)
}

// GetActionIDForProposeTransfer returns the action ID for the proposed transfer operation
func (dataGetter *drtClientDataGetter) GetActionIDForProposeTransfer(ctx context.Context, batch *bridgeCore.TransferBatch) (uint64, error) {
	if batch == nil {
		return 0, clients.ErrNilBatch
	}

	builder := dataGetter.createMultisigDefaultVmQueryBuilder()
	builder.Function(getActionIdForTransferBatchFuncName).ArgInt64(int64(batch.ID))
	dataGetter.addBatchInfo(builder, batch)

	return dataGetter.executeQueryUint64FromBuilder(ctx, builder)
}

// WasProposedSetStatus returns true if the proposed set status was triggered
func (dataGetter *drtClientDataGetter) WasProposedSetStatus(ctx context.Context, batch *bridgeCore.TransferBatch) (bool, error) {
	if batch == nil {
		return false, clients.ErrNilBatch
	}

	builder := dataGetter.createMultisigDefaultVmQueryBuilder()
	builder.Function(wasSetCurrentTransactionBatchStatusActionProposedFuncName).ArgInt64(int64(batch.ID))
	for _, stat := range batch.Statuses {
		builder.ArgBytes([]byte{stat})
	}

	return dataGetter.executeQueryBoolFromBuilder(ctx, builder)
}

// GetTransactionsStatuses will return the transactions statuses from the batch ID
func (dataGetter *drtClientDataGetter) GetTransactionsStatuses(ctx context.Context, batchID uint64) ([]byte, error) {
	builder := dataGetter.createMultisigDefaultVmQueryBuilder()
	builder.Function(getStatusesAfterExecutionFuncName).ArgInt64(int64(batchID))

	values, err := dataGetter.executeQueryFromBuilder(ctx, builder)
	if err != nil {
		return nil, err
	}
	if len(values) == 0 {
		return nil, fmt.Errorf("%w for batch ID %v", errNoStatusForBatchID, batchID)
	}

	isFinished, err := dataGetter.parseBool(values[0], getStatusesAfterExecutionFuncName, dataGetter.bech32MultisigContractAddress)
	if err != nil {
		return nil, err
	}
	if !isFinished {
		return nil, fmt.Errorf("%w for batch ID %v", errBatchNotFinished, batchID)
	}

	results := make([]byte, len(values)-1)
	for i := 1; i < len(values); i++ {
		results[i-1], err = getStatusFromBuff(values[i])
		if err != nil {
			return nil, fmt.Errorf("%w for result index %d", err, i-1)
		}
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("%w status is finished, no results are given", errMalformedBatchResponse)
	}

	return results, nil
}

// GetActionIDForSetStatusOnPendingTransfer returns the action ID for setting the status on the pending transfer batch
func (dataGetter *drtClientDataGetter) GetActionIDForSetStatusOnPendingTransfer(ctx context.Context, batch *bridgeCore.TransferBatch) (uint64, error) {
	if batch == nil {
		return 0, clients.ErrNilBatch
	}

	builder := dataGetter.createMultisigDefaultVmQueryBuilder()
	builder.Function(getActionIdForSetCurrentTransactionBatchStatusFuncName).ArgInt64(int64(batch.ID))
	for _, stat := range batch.Statuses {
		builder.ArgBytes([]byte{stat})
	}

	return dataGetter.executeQueryUint64FromBuilder(ctx, builder)
}

// QuorumReached returns true if the provided action ID reached the set quorum
func (dataGetter *drtClientDataGetter) QuorumReached(ctx context.Context, actionID uint64) (bool, error) {
	builder := dataGetter.createMultisigDefaultVmQueryBuilder()
	builder.Function(quorumReachedFuncName).ArgInt64(int64(actionID))

	return dataGetter.executeQueryBoolFromBuilder(ctx, builder)
}

// GetLastExecutedEthBatchID returns the last executed Ethereum batch ID
func (dataGetter *drtClientDataGetter) GetLastExecutedEthBatchID(ctx context.Context) (uint64, error) {
	builder := dataGetter.createMultisigDefaultVmQueryBuilder().Function(getLastExecutedEthBatchIdFuncName)

	return dataGetter.executeQueryUint64FromBuilder(ctx, builder)
}

// GetLastExecutedEthTxID returns the last executed Ethereum deposit ID
func (dataGetter *drtClientDataGetter) GetLastExecutedEthTxID(ctx context.Context) (uint64, error) {
	builder := dataGetter.createMultisigDefaultVmQueryBuilder().Function(getLastExecutedEthTxId)

	return dataGetter.executeQueryUint64FromBuilder(ctx, builder)
}

// WasSigned returns true if the action was already signed by the current relayer
func (dataGetter *drtClientDataGetter) WasSigned(ctx context.Context, actionID uint64) (bool, error) {
	builder := dataGetter.createMultisigDefaultVmQueryBuilder()
	builder.Function(signedFuncName).ArgAddress(dataGetter.relayerAddress).ArgInt64(int64(actionID))

	return dataGetter.executeQueryBoolFromBuilder(ctx, builder)
}

// GetAllStakedRelayers returns all staked relayers defined in DharitrI SC
func (dataGetter *drtClientDataGetter) GetAllStakedRelayers(ctx context.Context) ([][]byte, error) {
	builder := dataGetter.createMultisigDefaultVmQueryBuilder()
	builder.Function(getAllStakedRelayersFuncName)

	return dataGetter.executeQueryFromBuilder(ctx, builder)
}

// IsPaused returns true if the multisig contract is paused
func (dataGetter *drtClientDataGetter) IsPaused(ctx context.Context) (bool, error) {
	builder := dataGetter.createMultisigDefaultVmQueryBuilder()
	builder.Function(isPausedFuncName)

	return dataGetter.executeQueryBoolFromBuilder(ctx, builder)
}

func (dataGetter *drtClientDataGetter) isMintBurnToken(ctx context.Context, token []byte) (bool, error) {
	builder := dataGetter.createSafeDefaultVmQueryBuilder()
	builder.Function(isMintBurnTokenFuncName).ArgBytes(token)

	return dataGetter.executeQueryBoolFromBuilder(ctx, builder)
}

// isNativeToken returns true if the token is native
func (dataGetter *drtClientDataGetter) isNativeToken(ctx context.Context, token []byte) (bool, error) {
	builder := dataGetter.createSafeDefaultVmQueryBuilder()
	builder.Function(isNativeTokenFuncName).ArgBytes(token)

	return dataGetter.executeQueryBoolFromBuilder(ctx, builder)
}

func (dataGetter *drtClientDataGetter) getTotalBalances(ctx context.Context, token []byte) (*big.Int, error) {
	builder := dataGetter.createSafeDefaultVmQueryBuilder()
	builder.Function(getTotalBalances).ArgBytes(token)

	return dataGetter.executeQueryBigIntFromBuilder(ctx, builder)
}

func (dataGetter *drtClientDataGetter) getMintBalances(ctx context.Context, token []byte) (*big.Int, error) {
	builder := dataGetter.createSafeDefaultVmQueryBuilder()
	builder.Function(getMintBalances).ArgBytes(token)

	return dataGetter.executeQueryBigIntFromBuilder(ctx, builder)
}

func (dataGetter *drtClientDataGetter) getBurnBalances(ctx context.Context, token []byte) (*big.Int, error) {
	builder := dataGetter.createSafeDefaultVmQueryBuilder()
	builder.Function(getBurnBalances).ArgBytes(token)

	return dataGetter.executeQueryBigIntFromBuilder(ctx, builder)
}

func (dataGetter *drtClientDataGetter) addBatchInfo(builder builders.VMQueryBuilder, batch *bridgeCore.TransferBatch) {
	for _, dt := range batch.Deposits {
		builder.ArgBytes(dt.FromBytes).
			ArgBytes(dt.ToBytes).
			ArgBytes(dt.DestinationTokenBytes).
			ArgBigInt(dt.Amount).
			ArgInt64(int64(dt.Nonce)).
			ArgBytes(dt.Data)
	}
}

func getStatusFromBuff(buff []byte) (byte, error) {
	if len(buff) == 0 {
		return 0, errMalformedBatchResponse
	}

	return buff[len(buff)-1], nil
}

// GetAllKnownTokens returns all registered tokens
func (dataGetter *drtClientDataGetter) GetAllKnownTokens(ctx context.Context) ([][]byte, error) {
	builder := dataGetter.createSafeDefaultVmQueryBuilder()
	builder.Function(getAllKnownTokens)

	return dataGetter.executeQueryFromBuilder(ctx, builder)
}

// GetLastDrtBatchID returns the highest batch ID the safe contract reached. This might be a WIP batch that is not executable yet
func (dataGetter *drtClientDataGetter) GetLastDrtBatchID(ctx context.Context) (uint64, error) {
	builder := dataGetter.createSafeDefaultVmQueryBuilder()
	builder.Function(getLastBatchId)

	return dataGetter.executeQueryUint64FromBuilder(ctx, builder)
}

// IsInterfaceNil returns true if there is no value under the interface
func (dataGetter *drtClientDataGetter) IsInterfaceNil() bool {
	return dataGetter == nil
}
