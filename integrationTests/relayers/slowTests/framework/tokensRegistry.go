package framework

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

type tokensRegistry struct {
	testing.TB
	mut    sync.RWMutex
	tokens map[string]*TokenData
}

// NewTokenRegistry creates a new instance of type tokens registry
func NewTokenRegistry(tb testing.TB) *tokensRegistry {
	return &tokensRegistry{
		TB:     tb,
		tokens: make(map[string]*TokenData, 100),
	}
}

// AddToken will add a new test token
func (registry *tokensRegistry) AddToken(params IssueTokenParams) {
	registry.mut.Lock()
	defer registry.mut.Unlock()

	_, found := registry.tokens[params.AbstractTokenIdentifier]
	require.False(registry, found, "can not register more than one instance of the same abstract token identifier %s", params.AbstractTokenIdentifier)

	newToken := &TokenData{
		AbstractTokenIdentifier:     params.AbstractTokenIdentifier,
		DrtUniversalTokenTicker:     params.DrtUniversalTokenTicker,
		DrtChainSpecificTokenTicker: params.DrtChainSpecificTokenDisplayName,
		PeerChainTokenName:          params.PeerChainTokenName,
		PeerChainTokenSymbol:        params.PeerChainTokenSymbol,
	}

	registry.tokens[params.AbstractTokenIdentifier] = newToken
}

// RegisterUniversalToken will save the universal token identifier
func (registry *tokensRegistry) RegisterUniversalToken(abstractTokenIdentifier string, drtUniversalToken string) {
	registry.mut.Lock()
	defer registry.mut.Unlock()

	data, found := registry.tokens[abstractTokenIdentifier]
	require.True(registry, found, "abstract token identifier not registered %s", abstractTokenIdentifier)

	data.DrtUniversalToken = drtUniversalToken
}

// RegisterChainSpecificToken will save the chain specific token identifier
func (registry *tokensRegistry) RegisterChainSpecificToken(abstractTokenIdentifier string, drtChainSpecificToken string) {
	registry.mut.Lock()
	defer registry.mut.Unlock()

	data, found := registry.tokens[abstractTokenIdentifier]
	require.True(registry, found, "abstract token identifier not registered %s", abstractTokenIdentifier)

	data.DrtChainSpecificToken = drtChainSpecificToken
}

// RegisterPeerChainAddressAndInfo will save under the mutex lock the provided peer chain token address and its info
func (registry *tokensRegistry) RegisterPeerChainAddressAndInfo(
	abstractTokenIdentifier string,
	peerChainAddress []byte,
	chainTokenInfo interface{},
) {
	registry.mut.Lock()
	defer registry.mut.Unlock()

	data, found := registry.tokens[abstractTokenIdentifier]
	require.True(registry, found, "abstract token identifier not registered %s", abstractTokenIdentifier)

	data.PeerChainTokenAddress = peerChainAddress
	data.PeerChainTokenInfo = chainTokenInfo
}

// GetTokenData will return the token data based on the abstract identifier provided
func (registry *tokensRegistry) GetTokenData(abstractTokenIdentifier string) *TokenData {
	registry.mut.RLock()
	defer registry.mut.RUnlock()

	return registry.tokens[abstractTokenIdentifier]
}
