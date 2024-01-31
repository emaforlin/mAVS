package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/emaforlin/mAVS/x/mavs/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) StoredVotingAll(ctx context.Context, req *types.QueryAllStoredVotingRequest) (*types.QueryAllStoredVotingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var storedVotings []types.StoredVoting

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	storedVotingStore := prefix.NewStore(store, types.KeyPrefix(types.StoredVotingKeyPrefix))

	pageRes, err := query.Paginate(storedVotingStore, req.Pagination, func(key []byte, value []byte) error {
		var storedVoting types.StoredVoting
		if err := k.cdc.Unmarshal(value, &storedVoting); err != nil {
			return err
		}

		storedVotings = append(storedVotings, storedVoting)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStoredVotingResponse{StoredVoting: storedVotings, Pagination: pageRes}, nil
}

func (k Keeper) StoredVoting(ctx context.Context, req *types.QueryGetStoredVotingRequest) (*types.QueryGetStoredVotingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetStoredVoting(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetStoredVotingResponse{StoredVoting: val}, nil
}
