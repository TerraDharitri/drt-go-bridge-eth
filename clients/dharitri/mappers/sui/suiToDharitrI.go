package sui

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/TerraDharitri/drt-go-bridge-eth/clients"
	"github.com/TerraDharitri/drt-go-bridge-eth/clients/dharitri/mappers"
	"github.com/TerraDharitri/drt-go-chain-core/core/check"
)

type suiToDharitrI struct {
	dg mappers.DataGetter
}

// NewSuiToDharitrIMapper returns a new instance of suiToDharitrI
func NewSuiToDharitrIMapper(dg mappers.DataGetter) (*suiToDharitrI, error) {
	if check.IfNil(dg) {
		return nil, clients.ErrNilDataGetter
	}

	return &suiToDharitrI{
		dg: dg,
	}, nil
}

// ConvertToken will return drt token id given a specific sui coin type
func (mapper *suiToDharitrI) ConvertToken(ctx context.Context, sourceBytes []byte) ([]byte, error) {

	response, err := mapper.dg.GetTokenIdForSuiCoin(ctx, sourceBytes)
	if err != nil {
		return nil, err
	}

	if len(response) == 0 {
		return nil, fmt.Errorf("%w for provided %s", mappers.ErrUnknownToken, hex.EncodeToString(sourceBytes))
	}

	return response[0], nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (mapper *suiToDharitrI) IsInterfaceNil() bool {
	return mapper == nil
}
