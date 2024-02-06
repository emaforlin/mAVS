package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/emaforlin/mAVS/x/mavs/types"
)

func (k msgServer) AddVoter(goCtx context.Context, msg *types.MsgAddVoter) (*types.MsgAddVoterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// TODO: Handling the message
	sysInfo, found := k.Keeper.GetSystemInfo(ctx)
	if !found {
		panic("SystemInfo not found")
	}

	requiredVotingId, err := strconv.ParseUint(msg.VotingId, 10, 64)
	if err != nil {
		return nil, err
	}

	if requiredVotingId > sysInfo.NextId {
		return nil, types.ErrInvalidVotingId
	}

	voting, found := k.Keeper.GetStoredVoting(ctx, msg.VotingId)
	if !found {
		return nil, types.ErrInvalidVotingId
	}

	if msg.Creator != voting.Creator {
		return nil, sdkerrors.ErrUnauthorized
	}

	dni, err := strconv.Atoi(msg.Dni)
	if err != nil {
		return nil, types.ErrInvalidVoterId
	}
	voting.ElectoralRoll.Voters[uint64(dni)] = &types.Voter{
		ProofId:  msg.ProofId,
		HasVoted: false,
	}

	err = voting.Validate()
	if err != nil {
		return nil, err
	}

	k.Keeper.SetStoredVoting(ctx, voting)

	return &types.MsgAddVoterResponse{
		Status: "ok",
	}, nil
}
