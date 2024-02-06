package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/emaforlin/mAVS/x/mavs/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowVoter(goCtx context.Context, req *types.QueryShowVoterRequest) (*types.QueryShowVoterResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	res, found := k.GetStoredVoting(ctx, req.VotingId)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}
	dni, err := strconv.ParseUint(req.Dni, 10, 64)
	if err != nil {
		return nil, status.Error(codes.Internal, "error parsing dni")
	}
	return &types.QueryShowVoterResponse{
		Voter: res.ElectoralRoll.Voters[dni],
	}, nil
}
