package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		SystemInfo: SystemInfo{
			NextId: 5,
		},
		StableCoinsWhiteListList: []StableCoinsWhiteList{
			{
				Index:     "1",
				Coin:      "usdc",
				Available: true,
			},
			{
				Index:     "2",
				Coin:      "usdt",
				Available: false,
			},
			{
				Index:     "3",
				Coin:      "dai",
				Available: true,
			},
			{
				Index:     "4",
				Coin:      "fdusd",
				Available: false,
			},
		},
		PoolList: []Pool{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in stableCoinsWhiteList
	stableCoinsWhiteListIndexMap := make(map[string]struct{})

	for _, elem := range gs.StableCoinsWhiteListList {
		index := string(StableCoinsWhiteListKey(elem.Index))
		if _, ok := stableCoinsWhiteListIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for stableCoinsWhiteList")
		}
		stableCoinsWhiteListIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in pool
	poolIndexMap := make(map[string]struct{})

	for _, elem := range gs.PoolList {
		index := string(PoolKey(elem.Index))
		if _, ok := poolIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for pool")
		}
		poolIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
