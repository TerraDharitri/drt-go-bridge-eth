package batchProcessor

import (
	"math/big"
	"testing"

	bridgeCore "github.com/TerraDharitri/drt-go-bridge-eth/core"
	"github.com/stretchr/testify/assert"
)

func TestExtractListToDrt(t *testing.T) {
	t.Parallel()

	testBatch := &bridgeCore.TransferBatch{
		ID: 37,
		Deposits: []*bridgeCore.DepositTransfer{
			{
				Nonce:                 1,
				ToBytes:               []byte("to 1"),
				FromBytes:             []byte("from 1"),
				SourceTokenBytes:      []byte("source token 1"),
				DestinationTokenBytes: []byte("destination token 1"),
				Amount:                big.NewInt(11),
			},
			{
				Nonce:                 2,
				ToBytes:               []byte("to 2"),
				FromBytes:             []byte("from 2"),
				SourceTokenBytes:      []byte("source token 2"),
				DestinationTokenBytes: []byte("destination token 2"),
				Amount:                big.NewInt(22),
			},
		},
		Statuses: nil,
	}

	args := ExtractListToDrt(testBatch)

	expectedPeerTokens := [][]byte{
		[]byte("source token 1"),
		[]byte("source token 2"),
	}
	assert.Equal(t, expectedPeerTokens, args.PeerTokens)

	expectedRecipients := [][]byte{
		[]byte("to 1"),
		[]byte("to 2"),
	}
	assert.Equal(t, expectedRecipients, args.Recipients)

	expectedDrtTokenBytes := [][]byte{
		[]byte("destination token 1"),
		[]byte("destination token 2"),
	}
	assert.Equal(t, expectedDrtTokenBytes, args.DrtTokenBytes)

	expectedAmounts := []*big.Int{
		big.NewInt(11),
		big.NewInt(22),
	}
	assert.Equal(t, expectedAmounts, args.Amounts)

	expectedNonces := []*big.Int{
		big.NewInt(1),
		big.NewInt(2),
	}
	assert.Equal(t, expectedNonces, args.Nonces)
}

func TestExtractListFromDrt(t *testing.T) {
	t.Parallel()

	testBatch := &bridgeCore.TransferBatch{
		ID: 37,
		Deposits: []*bridgeCore.DepositTransfer{
			{
				Nonce:                 1,
				ToBytes:               []byte("to 1"),
				FromBytes:             []byte("from 1"),
				SourceTokenBytes:      []byte("source token 1"),
				DestinationTokenBytes: []byte("destination token 1"),
				Amount:                big.NewInt(11),
			},
			{
				Nonce:                 2,
				ToBytes:               []byte("to 2"),
				FromBytes:             []byte("from 2"),
				SourceTokenBytes:      []byte("source token 2"),
				DestinationTokenBytes: []byte("destination token 2"),
				Amount:                big.NewInt(22),
			},
		},
		Statuses: nil,
	}

	args := ExtractListFromDrt(testBatch)

	expectedPeerTokens := [][]byte{
		[]byte("destination token 1"),
		[]byte("destination token 2"),
	}
	assert.Equal(t, expectedPeerTokens, args.PeerTokens)

	expectedRecipients := [][]byte{
		[]byte("to 1"),
		[]byte("to 2"),
	}
	assert.Equal(t, expectedRecipients, args.Recipients)

	expectedDrtTokenBytes := [][]byte{
		[]byte("source token 1"),
		[]byte("source token 2"),
	}
	assert.Equal(t, expectedDrtTokenBytes, args.DrtTokenBytes)

	expectedAmounts := []*big.Int{
		big.NewInt(11),
		big.NewInt(22),
	}
	assert.Equal(t, expectedAmounts, args.Amounts)

	expectedNonces := []*big.Int{
		big.NewInt(1),
		big.NewInt(2),
	}
	assert.Equal(t, expectedNonces, args.Nonces)
}
