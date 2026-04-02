package eth

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/TerraDharitri/drt-go-bridge-eth/clients"
	"github.com/TerraDharitri/drt-go-bridge-eth/clients/dharitri/mappers"
	"github.com/TerraDharitri/drt-go-chain-core/core/check"
)

type erc20ToDharitrI struct {
	dg mappers.DataGetter
}

// NewErc20ToDharitrIMapper returns a new instance of erc20ToDharitrI
func NewErc20ToDharitrIMapper(dg mappers.DataGetter) (*erc20ToDharitrI, error) {
	if check.IfNil(dg) {
		return nil, clients.ErrNilDataGetter
	}

	return &erc20ToDharitrI{
		dg: dg,
	}, nil
}

// ConvertToken will return drt token id given a specific erc20 address
func (mapper *erc20ToDharitrI) ConvertToken(ctx context.Context, sourceBytes []byte) ([]byte, error) {

	response, err := mapper.dg.GetTokenIdForErc20Address(ctx, sourceBytes)
	if err != nil {
		return nil, err
	}

	if len(response) == 0 {
		return nil, fmt.Errorf("%w for provided %s", mappers.ErrUnknownToken, hex.EncodeToString(sourceBytes))
	}

	return response[0], nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (mapper *erc20ToDharitrI) IsInterfaceNil() bool {
	return mapper == nil
}
