package keeper

import (
	"context"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/emaforlin/mAVS/x/mavs/directives"
	"github.com/emaforlin/mAVS/x/mavs/types"
)

func (k msgServer) CreateVoting(goCtx context.Context, msg *types.MsgCreateVoting) (*types.MsgCreateVotingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	systemInfo, found := k.Keeper.GetSystemInfo(ctx)
	if !found {
		panic("SystemInfo not found")
	}
	newIndex := strconv.FormatUint(systemInfo.NextId, 10)
	newVoting := directives.NewVoting(msg.Title, directives.VotingTimeFromString(msg.Timewindow), nil, strings.Split(msg.Candidates, " ")...)

	roll := make(map[uint64]*types.Voter)
	roll[0] = &types.Voter{
		ProofId:  "0",
		HasVoted: false,
	}
	storedVoting := types.StoredVoting{
		Index:      newIndex,
		Title:      msg.Title,
		Timewindow: msg.Timewindow,
		Counting:   newVoting.VoteCount,
		Creator:    msg.Creator,
		ElectoralRoll: &types.ElectoralRoll{
			Voters: roll,
		},
	}

	err := storedVoting.Validate()
	if err != nil {
		return nil, err
	}

	k.Keeper.SetStoredVoting(ctx, storedVoting)
	systemInfo.NextId++
	k.Keeper.SetSystemInfo(ctx, systemInfo)

	return &types.MsgCreateVotingResponse{
		VotingIndex: newIndex,
	}, nil
}
