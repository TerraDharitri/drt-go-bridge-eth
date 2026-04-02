package batchProcessor

import (
	"math/big"

	bridgeCore "github.com/TerraDharitri/drt-go-bridge-eth/core"
)

// Direction is the direction of the transfer
type Direction string

const (
	// FromDharitrI is the direction of the transfer
	FromDharitrI Direction = "FromDharitrI"
	// ToDharitrI is the direction of the transfer
	ToDharitrI Direction = "ToDharitrI"
)

// ArgListsBatch is a struct that contains the batch data in a format that is easy to use
type ArgListsBatch struct {
	PeerTokens    [][]byte
	Recipients    [][]byte
	DrtTokenBytes [][]byte
	Amounts       []*big.Int
	Nonces        []*big.Int
	Direction     Direction
}

// ExtractListFromDrt will extract the batch data into a format that is easy to use
// The transfer is from DharitrI
func ExtractListFromDrt(batch *bridgeCore.TransferBatch) *ArgListsBatch {
	arg := &ArgListsBatch{
		Direction: FromDharitrI,
	}

	for _, dt := range batch.Deposits {
		recipient := dt.ToBytes
		arg.Recipients = append(arg.Recipients, recipient)

		token := dt.DestinationTokenBytes
		arg.PeerTokens = append(arg.PeerTokens, token)

		amount := big.NewInt(0).Set(dt.Amount)
		arg.Amounts = append(arg.Amounts, amount)

		nonce := big.NewInt(0).SetUint64(dt.Nonce)
		arg.Nonces = append(arg.Nonces, nonce)

		arg.DrtTokenBytes = append(arg.DrtTokenBytes, dt.SourceTokenBytes)
	}

	return arg
}

// ExtractListToDrt will extract the batch data into a format that is easy to use
// The transfer is to DharitrI
func ExtractListToDrt(batch *bridgeCore.TransferBatch) *ArgListsBatch {
	arg := &ArgListsBatch{
		Direction: ToDharitrI,
	}

	for _, dt := range batch.Deposits {
		recipient := dt.ToBytes
		arg.Recipients = append(arg.Recipients, recipient)

		token := dt.SourceTokenBytes
		arg.PeerTokens = append(arg.PeerTokens, token)

		amount := big.NewInt(0).Set(dt.Amount)
		arg.Amounts = append(arg.Amounts, amount)

		nonce := big.NewInt(0).SetUint64(dt.Nonce)
		arg.Nonces = append(arg.Nonces, nonce)

		arg.DrtTokenBytes = append(arg.DrtTokenBytes, dt.DestinationTokenBytes)
	}

	return arg
}
