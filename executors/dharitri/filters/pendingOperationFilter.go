package filters

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/TerraDharitri/drt-go-bridge-eth/config"
	"github.com/TerraDharitri/drt-go-bridge-eth/parsers"
	"github.com/TerraDharitri/drt-go-chain-core/core/check"
	logger "github.com/TerraDharitri/drt-go-chain-logger"
	"github.com/TerraDharitri/drt-go-sdk/data"
)

const (
	wildcardString   = "*"
	emptyString      = ""
	ethAddressPrefix = "0x"
)

var ethWildcardString = ""

func init() {
	var ethAddressWildcard = common.Address{}
	ethAddressWildcard.SetBytes([]byte(wildcardString))
	ethWildcardString = ethAddressWildcard.String()
}

type pendingOperationFilter struct {
	allowedEthAddresses []string
	deniedEthAddresses  []string
	allowedDrtAddresses []string
	deniedDrtAddresses  []string
	allowedTokens       []string
	deniedTokens        []string
}

// NewPendingOperationFilter creates a new instance of type pendingOperationFilter
func NewPendingOperationFilter(cfg config.PendingOperationsFilterConfig, log logger.Logger) (*pendingOperationFilter, error) {
	if check.IfNil(log) {
		return nil, errNilLogger
	}
	if len(cfg.AllowedDrtAddresses)+len(cfg.AllowedEthAddresses)+len(cfg.AllowedTokens) == 0 {
		return nil, errNoItemsAllowed
	}

	filter := &pendingOperationFilter{}
	err := filter.parseConfigs(cfg)
	if err != nil {
		return nil, err
	}

	err = filter.checkLists()
	if err != nil {
		return nil, err
	}

	log.Info("NewPendingOperationFilter config options",
		"DeniedEthAddresses", strings.Join(filter.deniedEthAddresses, ", "),
		"DeniedDrtAddresses", strings.Join(filter.deniedDrtAddresses, ", "),
		"DeniedTokens", strings.Join(filter.deniedTokens, ", "),
		"AllowedEthAddresses", strings.Join(filter.allowedEthAddresses, ", "),
		"AllowedDrtAddresses", strings.Join(filter.allowedDrtAddresses, ", "),
		"AllowedTokens", strings.Join(filter.allowedTokens, ", "),
	)

	return filter, nil
}

func (filter *pendingOperationFilter) parseConfigs(cfg config.PendingOperationsFilterConfig) error {
	var err error

	// denied lists do not support wildcard items
	filter.deniedEthAddresses, err = parseList(cfg.DeniedEthAddresses, wildcardString)
	if err != nil {
		return fmt.Errorf("%w in list DeniedEthAddresses", err)
	}

	filter.deniedDrtAddresses, err = parseList(cfg.DeniedDrtAddresses, wildcardString)
	if err != nil {
		return fmt.Errorf("%w in list DeniedDrtAddresses", err)
	}

	filter.deniedTokens, err = parseList(cfg.DeniedTokens, wildcardString)
	if err != nil {
		return fmt.Errorf("%w in list DeniedTokens", err)
	}

	// allowed lists do not support empty items
	filter.allowedEthAddresses, err = parseList(cfg.AllowedEthAddresses, emptyString)
	if err != nil {
		return fmt.Errorf("%w in list AllowedEthAddresses", err)
	}

	filter.allowedDrtAddresses, err = parseList(cfg.AllowedDrtAddresses, emptyString)
	if err != nil {
		return fmt.Errorf("%w in list AllowedDrtAddresses", err)
	}

	filter.allowedTokens, err = parseList(cfg.AllowedTokens, emptyString)
	if err != nil {
		return fmt.Errorf("%w in list AllowedTokens", err)
	}

	return nil
}

func parseList(list []string, unsupportedMarker string) ([]string, error) {
	newList := make([]string, 0, len(list))
	for index, item := range list {
		item = strings.ToLower(item)
		item = strings.Trim(item, "\r\n \t")
		if item == unsupportedMarker {
			return nil, fmt.Errorf("%w %s on item at index %d", errUnsupportedMarker, unsupportedMarker, index)
		}

		newList = append(newList, item)
	}

	return newList, nil
}

func (filter *pendingOperationFilter) checkLists() error {
	err := filter.checkList(filter.allowedEthAddresses, checkEthItemValid)
	if err != nil {
		return fmt.Errorf("%w in list AllowedEthAddresses", err)
	}

	err = filter.checkList(filter.deniedEthAddresses, checkEthItemValid)
	if err != nil {
		return fmt.Errorf("%w in list DeniedEthAddresses", err)
	}

	err = filter.checkList(filter.allowedDrtAddresses, checkDrtItemValid)
	if err != nil {
		return fmt.Errorf("%w in list AllowedDrtAddresses", err)
	}

	err = filter.checkList(filter.deniedDrtAddresses, checkDrtItemValid)
	if err != nil {
		return fmt.Errorf("%w in list DeniedDrtAddresses", err)
	}

	return nil
}

func (filter *pendingOperationFilter) checkList(list []string, checkItem func(item string) error) error {
	for index, item := range list {
		if item == wildcardString {
			continue
		}

		err := checkItem(item)
		if err != nil {
			return fmt.Errorf("%w on item at index %d", err, index)
		}
	}

	return nil
}

func checkDrtItemValid(item string) error {
	_, errNewAddr := data.NewAddressFromBech32String(item)
	return errNewAddr
}

func checkEthItemValid(item string) error {
	if !strings.HasPrefix(item, ethAddressPrefix) {
		return fmt.Errorf("%w (missing %s prefix)", errMissingEthPrefix, ethAddressPrefix)
	}

	return nil
}

// ShouldExecute returns true if the To, From or token are not denied and allowed
func (filter *pendingOperationFilter) ShouldExecute(callData parsers.ProxySCCompleteCallData) bool {
	if check.IfNil(callData.To) {
		return false
	}

	toAddress, err := callData.To.AddressAsBech32String()
	if err != nil {
		return false
	}

	isSpecificallyDenied := filter.stringExistsInList(callData.From.String(), filter.deniedEthAddresses, ethWildcardString) ||
		filter.stringExistsInList(toAddress, filter.deniedDrtAddresses, wildcardString) ||
		filter.stringExistsInList(callData.Token, filter.deniedTokens, wildcardString)
	if isSpecificallyDenied {
		return false
	}

	isAllowed := filter.stringExistsInList(callData.From.String(), filter.allowedEthAddresses, ethWildcardString) ||
		filter.stringExistsInList(toAddress, filter.allowedDrtAddresses, wildcardString) ||
		filter.stringExistsInList(callData.Token, filter.allowedTokens, wildcardString)

	return isAllowed
}

func (filter *pendingOperationFilter) stringExistsInList(needle string, haystack []string, wildcardMarker string) bool {
	needle = strings.ToLower(needle)
	wildcardMarker = strings.ToLower(wildcardMarker)

	for _, item := range haystack {
		if item == wildcardMarker {
			return true
		}

		if item == needle {
			return true
		}
	}

	return false
}

// IsInterfaceNil returns true if there is no value under the interface
func (filter *pendingOperationFilter) IsInterfaceNil() bool {
	return filter == nil
}
