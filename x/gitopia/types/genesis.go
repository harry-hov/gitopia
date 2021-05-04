package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		WhoisList: []*Whois{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate
	// Check for duplicated ID in whois
	whoisNameMap := make(map[string]bool)

	for _, elem := range gs.WhoisList {
		if _, ok := whoisNameMap[elem.Name]; ok {
			return fmt.Errorf("duplicated id for whois")
		}
		whoisNameMap[elem.Name] = true
	}

	return nil
}