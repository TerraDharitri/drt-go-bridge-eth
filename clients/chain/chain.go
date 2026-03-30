package chain

import (
	"fmt"
	"strings"
)

const (
	peerChainToDharitrINameTemplate   = "%sToDharitrI"
	dharitrIToPeerChainNameTemplate   = "DharitrITo%s"
	baseLogIdTemplate                   = "%sDharitrI-Base"
	dharitrIClientLogIdTemplate       = "%sDharitrI-DharitrIClient"
	dharitrIDataGetterLogIdTemplate   = "%sDharitrI-DharitrIDataGetter"
	peerChainDataGetterLogIdTemplate    = "%sDharitrI-%sDataGetter"
	peerChainClientLogIdTemplate        = "%sDharitrI-%sClient"
	dharitrIRoleProviderLogIdTemplate = "%sDharitrI-DharitrIRoleProvider"
	peerChainRoleProviderLogIdTemplate  = "%sDharitrI-%sRoleProvider"
	broadcasterLogIdTemplate            = "%sDharitrI-Broadcaster"
)

// Chain defines all the chain supported
type Chain string

const (
	// DharitrI is the string representation of the DharitrI chain
	DharitrI Chain = "msx"

	// Ethereum is the string representation of the Ethereum chain
	Ethereum Chain = "Ethereum"

	// Bsc is the string representation of the Binance smart chain
	Bsc Chain = "Bsc"

	// Polygon is the string representation of the Polygon chain
	Polygon Chain = "Polygon"

	// Sui is the string representation of the Sui chain
	Sui Chain = "Sui"
)

// ToLower returns the lowercase string of chain
func (c Chain) ToLower() string {
	return strings.ToLower(string(c))
}

// PeerChainToDharitrIName returns the string using chain value and peerChainToDharitrINameTemplate
func (c Chain) PeerChainToDharitrIName() string {
	return fmt.Sprintf(peerChainToDharitrINameTemplate, c)
}

// DharitrIToPeerChainName returns the string using chain value and dharitrIToPeerChainNameTemplate
func (c Chain) DharitrIToPeerChainName() string {
	return fmt.Sprintf(dharitrIToPeerChainNameTemplate, c)
}

// BaseLogId returns the string using chain value and baseLogIdTemplate
func (c Chain) BaseLogId() string {
	return fmt.Sprintf(baseLogIdTemplate, c)
}

// DharitrIClientLogId returns the string using chain value and dharitrIClientLogIdTemplate
func (c Chain) DharitrIClientLogId() string {
	return fmt.Sprintf(dharitrIClientLogIdTemplate, c)
}

// DharitrIDataGetterLogId returns the string using chain value and dharitrIDataGetterLogIdTemplate
func (c Chain) DharitrIDataGetterLogId() string {
	return fmt.Sprintf(dharitrIDataGetterLogIdTemplate, c)
}

// PeerChainDataGetterLogId returns the string using chain value and peerChainDataGetterLogIdTemplate
func (c Chain) PeerChainDataGetterLogId() string {
	return fmt.Sprintf(peerChainDataGetterLogIdTemplate, c, c)
}

// PeerChainClientLogId returns the string using chain value and peerChainClientLogIdTemplate
func (c Chain) PeerChainClientLogId() string {
	return fmt.Sprintf(peerChainClientLogIdTemplate, c, c)
}

// DharitrIRoleProviderLogId returns the string using chain value and dharitrIRoleProviderLogIdTemplate
func (c Chain) DharitrIRoleProviderLogId() string {
	return fmt.Sprintf(dharitrIRoleProviderLogIdTemplate, c)
}

// PeerChainRoleProviderLogId returns the string using chain value and peerChainRoleProviderLogIdTemplate
func (c Chain) PeerChainRoleProviderLogId() string {
	return fmt.Sprintf(peerChainRoleProviderLogIdTemplate, c, c)
}

// BroadcasterLogId returns the string using chain value and broadcasterLogIdTemplate
func (c Chain) BroadcasterLogId() string {
	return fmt.Sprintf(broadcasterLogIdTemplate, c)
}
