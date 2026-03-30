package framework

import (
	"math/big"

	"github.com/TerraDharitri/drt-go-chain-core/core/pubkeyConverter"
	logger "github.com/TerraDharitri/drt-go-chain-logger"
)

var (
	log                       = logger.GetOrCreate("integrationtests/slowtests")
	addressPubkeyConverter, _ = pubkeyConverter.NewBech32PubkeyConverter(32, "drt")
	zeroValueBigInt           = big.NewInt(0)
)
