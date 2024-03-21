package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/fenriz07/simpleswap/testutil/keeper"
	"github.com/fenriz07/simpleswap/x/simpleswap/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.SimpleswapKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
