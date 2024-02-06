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

	// try to get the systemInfo
	sysInfo, found := k.Keeper.GetSystemInfo(ctx)
	if !found {
		panic("SystemInfo not found")
	}

	// parse voting id to uint
	requiredVotingId, err := strconv.ParseUint(msg.VotingId, 10, 64)
	if err != nil {
		return nil, err
	}

	// checks that voting id is valid
	if requiredVotingId > sysInfo.NextId {
		return nil, types.ErrInvalidVotingId
	}

	// try to retrive voting with the required id
	voting, found := k.Keeper.GetStoredVoting(ctx, msg.VotingId)
	if !found {
		return nil, types.ErrInvalidVotingId
	}

	// accept new voters only from the voting creator
	if msg.Creator != voting.Creator {
		return nil, sdkerrors.ErrUnauthorized
	}

	// parse the dni
	dni, err := strconv.ParseUint(msg.Dni, 10, 64)
	if err != nil {
		return nil, types.ErrInvalidVoter
	}

	// insert the new voter in the electoral roll of the stored voting
	voting.ElectoralRoll.Voters[dni] = &types.Voter{
		ProofId:  msg.ProofId,
		HasVoted: false,
	}

	// look for other errors
	err = voting.Validate()
	if err != nil {
		return nil, err
	}

	// save changes
	k.Keeper.SetStoredVoting(ctx, voting)

	return &types.MsgAddVoterResponse{
		Status: "ok",
	}, nil
}
