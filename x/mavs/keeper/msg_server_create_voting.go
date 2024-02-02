package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/emaforlin/mAVS/x/mavs/types"
)

func (k msgServer) CreateVoting(goCtx context.Context, msg *types.MsgCreateVoting) (*types.MsgCreateVotingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateVotingResponse{}, nil
}
