package sui

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/TerraDharitri/drt-go-bridge-eth/clients"
	"github.com/TerraDharitri/drt-go-bridge-eth/clients/dharitri/mappers"
	"github.com/TerraDharitri/drt-go-chain-core/core/check"
)

type dharitrIToSui struct {
	dg mappers.DataGetter
}

// NewDharitrIToSuiMapper returns a new instance of dharitrIToSui
func NewDharitrIToSuiMapper(dg mappers.DataGetter) (*dharitrIToSui, error) {
	if check.IfNil(dg) {
		return nil, clients.ErrNilDataGetter
	}

	return &dharitrIToSui{
		dg: dg,
	}, nil
}

// ConvertToken will return drt token id given a specific sui coin type
func (mapper *dharitrIToSui) ConvertToken(ctx context.Context, sourceBytes []byte) ([]byte, error) {

	response, err := mapper.dg.GetSuiCoinForTokenId(ctx, sourceBytes)
	if err != nil {
		return nil, err
	}

	if len(response) == 0 {
		return nil, fmt.Errorf("%w for provided %s", mappers.ErrUnknownToken, hex.EncodeToString(sourceBytes))
	}

	return response[0], nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (mapper *dharitrIToSui) IsInterfaceNil() bool {
	return mapper == nil
}
