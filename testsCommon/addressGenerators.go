package testsCommon

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	sdkCore "github.com/TerraDharitri/drt-go-sdk/core"
	"github.com/TerraDharitri/drt-go-sdk/data"
)

const suiAddressBytesLength = 32

// CreateRandomEthereumAddress will create a random Ethereum address
func CreateRandomEthereumAddress() common.Address {
	buff := make([]byte, len(common.Address{}))
	_, _ = rand.Read(buff)

	return common.BytesToAddress(buff)
}

// CreateRandomDharitrIAddress will create a random DharitrI address
func CreateRandomDharitrIAddress() sdkCore.AddressHandler {
	buff := make([]byte, 32)
	_, _ = rand.Read(buff)

	return data.NewAddressFromBytes(buff)
}

// CreateRandomDharitrISCAddress will create a random DharitrI smart contract address
func CreateRandomDharitrISCAddress() sdkCore.AddressHandler {
	buff := make([]byte, 22)
	_, _ = rand.Read(buff)

	firstPart := append(make([]byte, 8), []byte{5, 0}...)

	return data.NewAddressFromBytes(append(firstPart, buff...))
}

// CreateRandomSuiAddressBytes will create a random Sui address bytes
func CreateRandomSuiAddressBytes() [suiAddressBytesLength]byte {
	buff := make([]byte, suiAddressBytesLength)
	_, _ = rand.Read(buff)

	return [suiAddressBytesLength]byte(buff)
}

// CreateRandomCoinId will create a random Sui coin ID
func CreateRandomCoinId() string {
	addrBytes := CreateRandomSuiAddressBytes()
	return fmt.Sprintf("0x%s::coin::COIN", hex.EncodeToString(addrBytes[:]))
}
