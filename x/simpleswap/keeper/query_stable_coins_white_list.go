package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/fenriz07/simpleswap/x/simpleswap/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) StableCoinsWhiteListAll(ctx context.Context, req *types.QueryAllStableCoinsWhiteListRequest) (*types.QueryAllStableCoinsWhiteListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var stableCoinsWhiteLists []types.StableCoinsWhiteList

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	stableCoinsWhiteListStore := prefix.NewStore(store, types.KeyPrefix(types.StableCoinsWhiteListKeyPrefix))

	pageRes, err := query.Paginate(stableCoinsWhiteListStore, req.Pagination, func(key []byte, value []byte) error {
		var stableCoinsWhiteList types.StableCoinsWhiteList
		if err := k.cdc.Unmarshal(value, &stableCoinsWhiteList); err != nil {
			return err
		}

		stableCoinsWhiteLists = append(stableCoinsWhiteLists, stableCoinsWhiteList)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStableCoinsWhiteListResponse{StableCoinsWhiteList: stableCoinsWhiteLists, Pagination: pageRes}, nil
}

func (k Keeper) StableCoinsWhiteList(ctx context.Context, req *types.QueryGetStableCoinsWhiteListRequest) (*types.QueryGetStableCoinsWhiteListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetStableCoinsWhiteList(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetStableCoinsWhiteListResponse{StableCoinsWhiteList: val}, nil
}
