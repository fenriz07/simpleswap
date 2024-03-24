package simpleswap

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/fenriz07/simpleswap/x/simpleswap/keeper"
	"github.com/fenriz07/simpleswap/x/simpleswap/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetSystemInfo(ctx, genState.SystemInfo)
	// Set all the stableCoinsWhiteList
	for _, elem := range genState.StableCoinsWhiteListList {
		k.SetStableCoinsWhiteList(ctx, elem)
	}
	// Set all the pool
	for _, elem := range genState.PoolList {
		k.SetPool(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all systemInfo
	systemInfo, found := k.GetSystemInfo(ctx)
	if found {
		genesis.SystemInfo = systemInfo
	}
	genesis.StableCoinsWhiteListList = k.GetAllStableCoinsWhiteList(ctx)
	genesis.PoolList = k.GetAllPool(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
