package chain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_peerChainToDharitrIName(t *testing.T) {
	assert.Equal(t, "EthereumToDharitrI", Ethereum.PeerChainToDharitrIName())
	assert.Equal(t, "BscToDharitrI", Bsc.PeerChainToDharitrIName())
	assert.Equal(t, "SuiToDharitrI", Sui.PeerChainToDharitrIName())
}

func Test_dharitrIToPeerChainName(t *testing.T) {
	assert.Equal(t, "DharitrIToEthereum", Ethereum.DharitrIToPeerChainName())
	assert.Equal(t, "DharitrIToBsc", Bsc.DharitrIToPeerChainName())
	assert.Equal(t, "DharitrIToSui", Sui.DharitrIToPeerChainName())
}

func Test_baseLogId(t *testing.T) {
	assert.Equal(t, "EthereumDharitrI-Base", Ethereum.BaseLogId())
	assert.Equal(t, "BscDharitrI-Base", Bsc.BaseLogId())
	assert.Equal(t, "SuiDharitrI-Base", Sui.BaseLogId())
}

func Test_dharitrIClientLogId(t *testing.T) {
	assert.Equal(t, "EthereumDharitrI-DharitrIClient", Ethereum.DharitrIClientLogId())
	assert.Equal(t, "BscDharitrI-DharitrIClient", Bsc.DharitrIClientLogId())
	assert.Equal(t, "SuiDharitrI-DharitrIClient", Sui.DharitrIClientLogId())
}

func Test_dharitrIDataGetterLogId(t *testing.T) {
	assert.Equal(t, "EthereumDharitrI-DharitrIDataGetter", Ethereum.DharitrIDataGetterLogId())
	assert.Equal(t, "BscDharitrI-DharitrIDataGetter", Bsc.DharitrIDataGetterLogId())
	assert.Equal(t, "SuiDharitrI-DharitrIDataGetter", Sui.DharitrIDataGetterLogId())
}

func Test_peerChainClientLogId(t *testing.T) {
	assert.Equal(t, "EthereumDharitrI-EthereumClient", Ethereum.PeerChainClientLogId())
	assert.Equal(t, "BscDharitrI-BscClient", Bsc.PeerChainClientLogId())
	assert.Equal(t, "SuiDharitrI-SuiClient", Sui.PeerChainClientLogId())
}

func Test_dharitrIRoleProviderLogId(t *testing.T) {
	assert.Equal(t, "EthereumDharitrI-DharitrIRoleProvider", Ethereum.DharitrIRoleProviderLogId())
	assert.Equal(t, "BscDharitrI-DharitrIRoleProvider", Bsc.DharitrIRoleProviderLogId())
	assert.Equal(t, "SuiDharitrI-DharitrIRoleProvider", Sui.DharitrIRoleProviderLogId())
}

func Test_peerChainRoleProviderLogId(t *testing.T) {
	assert.Equal(t, "EthereumDharitrI-EthereumRoleProvider", Ethereum.PeerChainRoleProviderLogId())
	assert.Equal(t, "BscDharitrI-BscRoleProvider", Bsc.PeerChainRoleProviderLogId())
	assert.Equal(t, "SuiDharitrI-SuiRoleProvider", Sui.PeerChainRoleProviderLogId())
}

func Test_broadcasterLogId(t *testing.T) {
	assert.Equal(t, "EthereumDharitrI-Broadcaster", Ethereum.BroadcasterLogId())
	assert.Equal(t, "BscDharitrI-Broadcaster", Bsc.BroadcasterLogId())
	assert.Equal(t, "SuiDharitrI-Broadcaster", Sui.BroadcasterLogId())
}

func TestToLower(t *testing.T) {
	assert.Equal(t, "msx", DharitrI.ToLower())
	assert.Equal(t, "ethereum", Ethereum.ToLower())
	assert.Equal(t, "bsc", Bsc.ToLower())
	assert.Equal(t, "sui", Sui.ToLower())
}
