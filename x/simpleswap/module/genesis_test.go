package simpleswap_test

import (
	"testing"

	keepertest "github.com/fenriz07/simpleswap/testutil/keeper"
	"github.com/fenriz07/simpleswap/testutil/nullify"
	simpleswap "github.com/fenriz07/simpleswap/x/simpleswap/module"
	"github.com/fenriz07/simpleswap/x/simpleswap/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		SystemInfo: types.SystemInfo{
			NextId: 96,
		},
		StableCoinsWhiteListList: []types.StableCoinsWhiteList{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		PoolList: []types.Pool{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SimpleswapKeeper(t)
	simpleswap.InitGenesis(ctx, k, genesisState)
	got := simpleswap.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.SystemInfo, got.SystemInfo)
	require.ElementsMatch(t, genesisState.StableCoinsWhiteListList, got.StableCoinsWhiteListList)
	require.ElementsMatch(t, genesisState.PoolList, got.PoolList)
	// this line is used by starport scaffolding # genesis/test/assert
}
