package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "github.com/fenriz07/simpleswap/testutil/keeper"
	"github.com/fenriz07/simpleswap/testutil/nullify"
	"github.com/fenriz07/simpleswap/x/simpleswap/keeper"
	"github.com/fenriz07/simpleswap/x/simpleswap/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNStableCoinsWhiteList(keeper keeper.Keeper, ctx context.Context, n int) []types.StableCoinsWhiteList {
	items := make([]types.StableCoinsWhiteList, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetStableCoinsWhiteList(ctx, items[i])
	}
	return items
}

func TestStableCoinsWhiteListGet(t *testing.T) {
	keeper, ctx := keepertest.SimpleswapKeeper(t)
	items := createNStableCoinsWhiteList(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetStableCoinsWhiteList(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestStableCoinsWhiteListRemove(t *testing.T) {
	keeper, ctx := keepertest.SimpleswapKeeper(t)
	items := createNStableCoinsWhiteList(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveStableCoinsWhiteList(ctx,
			item.Index,
		)
		_, found := keeper.GetStableCoinsWhiteList(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestStableCoinsWhiteListGetAll(t *testing.T) {
	keeper, ctx := keepertest.SimpleswapKeeper(t)
	items := createNStableCoinsWhiteList(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllStableCoinsWhiteList(ctx)),
	)
}
