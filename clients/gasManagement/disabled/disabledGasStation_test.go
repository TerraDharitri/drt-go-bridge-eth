package disabled

import (
	"math/big"
	"testing"

	"github.com/TerraDharitri/drt-go-chain-core/core/check"
	"github.com/stretchr/testify/assert"
)

func TestNewDisabledGasStation(t *testing.T) {
	dgs := &DisabledGasStation{}

	assert.False(t, check.IfNil(dgs))

	gasPrice, err := dgs.GetCurrentGasPrice()
	assert.Equal(t, big.NewInt(defaultDisabledGasPrice), gasPrice)
	assert.Nil(t, err)

	err = dgs.Close()
	assert.Nil(t, err)
}
