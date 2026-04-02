package balanceValidator

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/TerraDharitri/drt-go-bridge-eth/clients"
	bridgeCore "github.com/TerraDharitri/drt-go-bridge-eth/core"
	"github.com/TerraDharitri/drt-go-bridge-eth/core/batchProcessor"
	"github.com/TerraDharitri/drt-go-chain-core/core/check"
	logger "github.com/TerraDharitri/drt-go-chain-logger"
)

// ArgsBalanceValidator represents the DTO struct used in the NewBalanceValidator constructor function
type ArgsBalanceValidator struct {
	Log              logger.Logger
	DharitrIClient DharitrIClient
	PeerChainClient  PeerChainClient
}

type balanceValidator struct {
	log              logger.Logger
	dharitrIClient DharitrIClient
	peerChainClient  PeerChainClient
}

// NewBalanceValidator creates a new instance of type balanceValidator
func NewBalanceValidator(args ArgsBalanceValidator) (*balanceValidator, error) {
	err := checkArgs(args)
	if err != nil {
		return nil, err
	}

	return &balanceValidator{
		log:              args.Log,
		dharitrIClient: args.DharitrIClient,
		peerChainClient:  args.PeerChainClient,
	}, nil
}

func checkArgs(args ArgsBalanceValidator) error {
	if check.IfNil(args.Log) {
		return ErrNilLogger
	}
	if check.IfNil(args.DharitrIClient) {
		return ErrNilDharitrIClient
	}
	if check.IfNil(args.PeerChainClient) {
		return ErrNilEthereumClient
	}

	return nil
}

// CheckToken returns error if the bridge can not happen to the provided token due to faulty balance values in the contracts
func (validator *balanceValidator) CheckToken(ctx context.Context, token []byte, drtToken []byte, amount *big.Int, direction batchProcessor.Direction) error {
	err := validator.checkRequiredBalance(ctx, token, drtToken, amount, direction)
	if err != nil {
		return err
	}

	isMintBurnOnEthereum, err := validator.isMintBurnOnEthereum(ctx, token)
	if err != nil {
		return err
	}

	isMintBurnOnDharitrI, err := validator.isMintBurnOnDharitrI(ctx, drtToken)
	if err != nil {
		return err
	}

	isNativeOnEthereum, err := validator.isNativeOnEthereum(ctx, token)
	if err != nil {
		return err
	}

	isNativeOnDharitrI, err := validator.isNativeOnDharitrI(ctx, drtToken)
	if err != nil {
		return err
	}

	if !isNativeOnEthereum && !isMintBurnOnEthereum {
		return fmt.Errorf("%w isNativeOnEthereum = %v, isMintBurnOnEthereum = %v", ErrInvalidSetup, isNativeOnEthereum, isMintBurnOnEthereum)
	}

	if !isNativeOnDharitrI && !isMintBurnOnDharitrI {
		return fmt.Errorf("%w isNativeOnDharitrI = %v, isMintBurnOnDharitrI = %v", ErrInvalidSetup, isNativeOnDharitrI, isMintBurnOnDharitrI)
	}

	if isNativeOnEthereum == isNativeOnDharitrI {
		return fmt.Errorf("%w isNativeOnEthereum = %v, isNativeOnDharitrI = %v", ErrInvalidSetup, isNativeOnEthereum, isNativeOnDharitrI)
	}

	ethAmount, err := validator.computeEthAmount(ctx, token, isMintBurnOnEthereum, isNativeOnEthereum)
	if err != nil {
		return err
	}
	drtAmount, err := validator.computeDrtAmount(ctx, drtToken, isMintBurnOnDharitrI, isNativeOnDharitrI)
	if err != nil {
		return err
	}

	validator.log.Debug("balanceValidator.CheckToken",
		"ERC20 token", token,
		"ERC20 balance", ethAmount.String(),
		"DCDT token", drtToken,
		"DCDT balance", drtAmount.String(),
		"amount", amount.String(),
	)

	if ethAmount.Cmp(drtAmount) != 0 {
		return fmt.Errorf("%w, balance for ERC20 token %s is %s and the balance for DCDT token %s is %s, direction %s",
			ErrBalanceMismatch, token, ethAmount.String(), drtToken, drtAmount.String(), direction)
	}
	return nil
}

func (validator *balanceValidator) checkRequiredBalance(ctx context.Context, token []byte, drtToken []byte, amount *big.Int, direction batchProcessor.Direction) error {
	switch direction {
	case batchProcessor.FromDharitrI:
		return validator.peerChainClient.CheckRequiredBalance(ctx, token, amount)
	case batchProcessor.ToDharitrI:
		return validator.dharitrIClient.CheckRequiredBalance(ctx, drtToken, amount)
	default:
		return fmt.Errorf("%w, direction: %s", ErrInvalidDirection, direction)
	}
}

func (validator *balanceValidator) isMintBurnOnEthereum(ctx context.Context, erc20Address []byte) (bool, error) {
	isMintBurn, err := validator.peerChainClient.MintBurnTokens(ctx, erc20Address)
	if err != nil {
		return false, err
	}

	return isMintBurn, nil
}

func (validator *balanceValidator) isNativeOnEthereum(ctx context.Context, token []byte) (bool, error) {
	isNative, err := validator.peerChainClient.NativeTokens(ctx, token)
	if err != nil {
		return false, err
	}
	return isNative, nil
}

func (validator *balanceValidator) isMintBurnOnDharitrI(ctx context.Context, token []byte) (bool, error) {
	isMintBurn, err := validator.dharitrIClient.IsMintBurnToken(ctx, token)
	if err != nil {
		return false, err
	}
	return isMintBurn, nil
}

func (validator *balanceValidator) isNativeOnDharitrI(ctx context.Context, token []byte) (bool, error) {
	isNative, err := validator.dharitrIClient.IsNativeToken(ctx, token)
	if err != nil {
		return false, err
	}
	return isNative, nil
}

func (validator *balanceValidator) computeEthAmount(
	ctx context.Context,
	token []byte,
	isMintBurn bool,
	isNative bool,
) (*big.Int, error) {
	ethAmountInPendingBatches, err := validator.getTotalTransferAmountInPendingEthBatches(ctx, token)
	if err != nil {
		return nil, err
	}

	if !isMintBurn {
		// we need to subtract all locked balances on the Ethereum side (all pending, un-executed batches) so the balances
		// with the minted DharitrI tokens will match
		total, errTotal := validator.peerChainClient.TotalBalances(ctx, token)
		if errTotal != nil {
			return nil, errTotal
		}

		return total.Sub(total, ethAmountInPendingBatches), nil
	}

	burnBalances, err := validator.peerChainClient.BurnBalances(ctx, token)
	if err != nil {
		return nil, err
	}
	mintBalances, err := validator.peerChainClient.MintBalances(ctx, token)
	if err != nil {
		return nil, err
	}

	// we need to cancel out what was burned in advance when the deposit was registered in the contract
	burnBalances.Sub(burnBalances, ethAmountInPendingBatches)

	var ethAmount *big.Int
	if isNative {
		ethAmount = big.NewInt(0).Sub(burnBalances, mintBalances)
	} else {
		ethAmount = big.NewInt(0).Sub(mintBalances, burnBalances)
	}

	if ethAmount.Cmp(big.NewInt(0)) < 0 {
		return big.NewInt(0), fmt.Errorf("%w, ethAmount: %s", ErrNegativeAmount, ethAmount.String())
	}
	return ethAmount, nil
}

func (validator *balanceValidator) computeDrtAmount(
	ctx context.Context,
	token []byte,
	isMintBurn bool,
	isNative bool,
) (*big.Int, error) {
	drtAmountInPendingBatches, err := validator.getTotalTransferAmountInPendingDrtBatches(ctx, token)
	if err != nil {
		return nil, err
	}

	if !isMintBurn {
		// we need to subtract all locked balances on the DharitrI side (all pending, un-executed batches) so the balances
		// with the minted Ethereum tokens will match
		total, errTotal := validator.dharitrIClient.TotalBalances(ctx, token)
		if errTotal != nil {
			return nil, errTotal
		}

		return total.Sub(total, drtAmountInPendingBatches), nil
	}

	burnBalances, err := validator.dharitrIClient.BurnBalances(ctx, token)
	if err != nil {
		return nil, err
	}
	mintBalances, err := validator.dharitrIClient.MintBalances(ctx, token)
	if err != nil {
		return nil, err
	}
	var drtAmount *big.Int

	// we need to cancel out what was burned in advance when the deposit was registered in the contract
	burnBalances.Sub(burnBalances, drtAmountInPendingBatches)

	if isNative {
		drtAmount = big.NewInt(0).Sub(burnBalances, mintBalances)
	} else {
		drtAmount = big.NewInt(0).Sub(mintBalances, burnBalances)
	}

	if drtAmount.Cmp(big.NewInt(0)) < 0 {
		return big.NewInt(0), fmt.Errorf("%w, drtAmount: %s", ErrNegativeAmount, drtAmount.String())
	}
	return drtAmount, nil
}

func getTotalAmountFromBatch(batch *bridgeCore.TransferBatch, token []byte) *big.Int {
	amount := big.NewInt(0)
	for _, deposit := range batch.Deposits {
		if bytes.Equal(deposit.SourceTokenBytes, token) {
			amount.Add(amount, deposit.Amount)
		}
	}

	return amount
}

func (validator *balanceValidator) getTotalTransferAmountInPendingDrtBatches(ctx context.Context, drtToken []byte) (*big.Int, error) {
	batchID, err := validator.dharitrIClient.GetLastDrtBatchID(ctx)
	if err != nil {
		return nil, err
	}

	var batch *bridgeCore.TransferBatch
	amount := big.NewInt(0)
	for {
		batch, err = validator.dharitrIClient.GetBatch(ctx, batchID)
		if errors.Is(err, clients.ErrNoBatchAvailable) {
			return amount, nil
		}
		if err != nil {
			return nil, err
		}

		wasExecuted, errWasExecuted := validator.peerChainClient.WasExecuted(ctx, batch.ID)
		if errWasExecuted != nil {
			return nil, errWasExecuted
		}
		if wasExecuted {
			return amount, nil
		}

		amountFromBatch := getTotalAmountFromBatch(batch, drtToken)
		amount.Add(amount, amountFromBatch)
		batchID-- // go to the previous batch
	}
}

func (validator *balanceValidator) getTotalTransferAmountInPendingEthBatches(ctx context.Context, token []byte) (*big.Int, error) {
	batchID, err := validator.dharitrIClient.GetLastExecutedEthBatchID(ctx)
	if err != nil {
		return nil, err
	}

	var batch *bridgeCore.TransferBatch
	amount := big.NewInt(0)
	for {
		batch, _, err = validator.peerChainClient.GetBatch(ctx, batchID+1) // we take all batches, regardless if they are final or not
		if err != nil {
			return nil, err
		}

		isBatchInvalid := batch.ID != batchID+1 || len(batch.Deposits) == 0
		if isBatchInvalid {
			return amount, nil
		}

		amountFromBatch := getTotalAmountFromBatch(batch, token)
		amount.Add(amount, amountFromBatch)
		batchID++
	}
}

// IsInterfaceNil returns true if there is no value under the interface
func (validator *balanceValidator) IsInterfaceNil() bool {
	return validator == nil
}
