package roleproviders

import "github.com/TerraDharitri/drt-go-sdk/core"

// DharitrIRoleProviderStub -
type DharitrIRoleProviderStub struct {
	IsWhitelistedCalled func(address core.AddressHandler) bool
}

// IsWhitelisted -
func (stub *DharitrIRoleProviderStub) IsWhitelisted(address core.AddressHandler) bool {
	if stub.IsWhitelistedCalled != nil {
		return stub.IsWhitelistedCalled(address)
	}

	return true
}

// IsInterfaceNil -
func (stub *DharitrIRoleProviderStub) IsInterfaceNil() bool {
	return stub == nil
}
