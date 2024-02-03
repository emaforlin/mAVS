package keeper_test

import (
	"context"
	"testing"
	"time"

	keepertest "github.com/emaforlin/mAVS/testutil/keeper"
	"github.com/emaforlin/mAVS/x/mavs/directives"
	"github.com/emaforlin/mAVS/x/mavs/keeper"
	mavs "github.com/emaforlin/mAVS/x/mavs/module"
	"github.com/emaforlin/mAVS/x/mavs/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServerCreateVoting(t testing.TB) (keeper.Keeper, types.MsgServer, context.Context) {
	k, ctx := keepertest.MavsKeeper(t)
	mavs.InitGenesis(ctx, k, *types.DefaultGenesis())
	return k, keeper.NewMsgServerImpl(k), ctx
}

func TestCreateVoting(t *testing.T) {
	_, msgServer, context := setupMsgServerCreateVoting(t)

	tw := directives.NewVotingTime(time.Now(), time.Now().Add(6*time.Hour))

	createResponse, err := msgServer.CreateVoting(context, &types.MsgCreateVoting{
		Creator:    alice,
		Title:      "test voting",
		Timewindow: tw.String(),
		Candidates: "candidate1 candidate2 candidate3",
	})

	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateVotingResponse{
		VotingIndex: "1",
	}, *createResponse)
}

func TestCreate1VotingHasSaved(t *testing.T) {
	keeper, msgSrvr, context := setupMsgServerCreateVoting(t)

	tw := directives.NewVotingTime(time.Now(), time.Now().Add(2*time.Hour))
	msgSrvr.CreateVoting(context, &types.MsgCreateVoting{
		Creator:    bob,
		Title:      "Voting saved",
		Timewindow: tw.String(),
		Candidates: "c1 c2 c3",
	})
	systemInfo, found := keeper.GetSystemInfo(context)
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{
		NextId: 2,
	}, systemInfo)

	voting1, found1 := keeper.GetStoredVoting(context, "1")
	require.True(t, found1)
	require.EqualValues(t, types.StoredVoting{
		Index:      "1",
		Timewindow: tw.String(),
		Title:      "Voting saved",
		Counting: map[string]uint64{
			"blank": 0,
			"c1":    0,
			"c2":    0,
			"c3":    0,
		},
	}, voting1)

}
