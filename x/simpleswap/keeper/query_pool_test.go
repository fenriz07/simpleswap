package keeper_test

import (
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/fenriz07/simpleswap/testutil/keeper"
	"github.com/fenriz07/simpleswap/testutil/nullify"
	"github.com/fenriz07/simpleswap/x/simpleswap/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestPoolQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.SimpleswapKeeper(t)
	msgs := createNPool(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetPoolRequest
		response *types.QueryGetPoolResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetPoolRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetPoolResponse{Pool: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetPoolRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetPoolResponse{Pool: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetPoolRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Pool(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestPoolQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.SimpleswapKeeper(t)
	msgs := createNPool(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllPoolRequest {
		return &types.QueryAllPoolRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.PoolAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Pool), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Pool),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.PoolAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Pool), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Pool),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.PoolAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Pool),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.PoolAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
