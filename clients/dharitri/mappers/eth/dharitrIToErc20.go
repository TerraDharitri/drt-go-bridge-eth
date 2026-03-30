package eth

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/TerraDharitri/drt-go-bridge-eth/clients"
	"github.com/TerraDharitri/drt-go-bridge-eth/clients/dharitri/mappers"
	"github.com/TerraDharitri/drt-go-chain-core/core/check"
)

type dharitrIToErc20 struct {
	dg mappers.DataGetter
}

// NewDharitrIToErc20Mapper returns a new instance of dharitrIToErc20
func NewDharitrIToErc20Mapper(dg mappers.DataGetter) (*dharitrIToErc20, error) {
	if check.IfNil(dg) {
		return nil, clients.ErrNilDataGetter
	}

	return &dharitrIToErc20{
		dg: dg,
	}, nil
}

// ConvertToken will return drt token id given a specific erc20 address
func (mapper *dharitrIToErc20) ConvertToken(ctx context.Context, sourceBytes []byte) ([]byte, error) {

	response, err := mapper.dg.GetERC20AddressForTokenId(ctx, sourceBytes)
	if err != nil {
		return nil, err
	}

	if len(response) == 0 {
		return nil, fmt.Errorf("%w for provided %s", mappers.ErrUnknownToken, hex.EncodeToString(sourceBytes))
	}

	return response[0], nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (mapper *dharitrIToErc20) IsInterfaceNil() bool {
	return mapper == nil
}
