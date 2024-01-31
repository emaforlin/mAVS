package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/emaforlin/mAVS/x/mavs/types"
)

// SetStoredVoting set a specific storedVoting in the store from its index
func (k Keeper) SetStoredVoting(ctx context.Context, storedVoting types.StoredVoting) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StoredVotingKeyPrefix))
	b := k.cdc.MustMarshal(&storedVoting)
	store.Set(types.StoredVotingKey(
		storedVoting.Index,
	), b)
}

// GetStoredVoting returns a storedVoting from its index
func (k Keeper) GetStoredVoting(
	ctx context.Context,
	index string,

) (val types.StoredVoting, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StoredVotingKeyPrefix))

	b := store.Get(types.StoredVotingKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStoredVoting removes a storedVoting from the store
func (k Keeper) RemoveStoredVoting(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StoredVotingKeyPrefix))
	store.Delete(types.StoredVotingKey(
		index,
	))
}

// GetAllStoredVoting returns all storedVoting
func (k Keeper) GetAllStoredVoting(ctx context.Context) (list []types.StoredVoting) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StoredVotingKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.StoredVoting
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
