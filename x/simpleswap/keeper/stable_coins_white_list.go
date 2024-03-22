package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/fenriz07/simpleswap/x/simpleswap/types"
)

// SetStableCoinsWhiteList set a specific stableCoinsWhiteList in the store from its index
func (k Keeper) SetStableCoinsWhiteList(ctx context.Context, stableCoinsWhiteList types.StableCoinsWhiteList) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StableCoinsWhiteListKeyPrefix))
	b := k.cdc.MustMarshal(&stableCoinsWhiteList)
	store.Set(types.StableCoinsWhiteListKey(
		stableCoinsWhiteList.Index,
	), b)
}

// GetStableCoinsWhiteList returns a stableCoinsWhiteList from its index
func (k Keeper) GetStableCoinsWhiteList(
	ctx context.Context,
	index string,

) (val types.StableCoinsWhiteList, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StableCoinsWhiteListKeyPrefix))

	b := store.Get(types.StableCoinsWhiteListKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStableCoinsWhiteList removes a stableCoinsWhiteList from the store
func (k Keeper) RemoveStableCoinsWhiteList(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StableCoinsWhiteListKeyPrefix))
	store.Delete(types.StableCoinsWhiteListKey(
		index,
	))
}

// GetAllStableCoinsWhiteList returns all stableCoinsWhiteList
func (k Keeper) GetAllStableCoinsWhiteList(ctx context.Context) (list []types.StableCoinsWhiteList) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StableCoinsWhiteListKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.StableCoinsWhiteList
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
