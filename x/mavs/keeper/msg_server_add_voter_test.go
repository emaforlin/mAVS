package keeper_test

import (
	"context"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	keepertest "github.com/emaforlin/mAVS/testutil/keeper"
	"github.com/emaforlin/mAVS/x/mavs/keeper"
	mavs "github.com/emaforlin/mAVS/x/mavs/module"
	"github.com/emaforlin/mAVS/x/mavs/types"
	"github.com/stretchr/testify/require"
)

// setups a new msgServer with voting already created
func setupMsgServerAddVoter(t testing.TB) (keeper.Keeper, types.MsgServer, context.Context) {
	k, ctx := keepertest.MavsKeeper(t)
	mavs.InitGenesis(ctx, k, *types.DefaultGenesis())
	kpr := keeper.NewMsgServerImpl(k)
	kpr.CreateVoting(ctx, &DefaultTestingMsgVoting)

	return k, kpr, ctx
}

func TestAddVoter(t *testing.T) {
	keeper, msgServer, ctx := setupMsgServerAddVoter(t)

	storedVoting, found := keeper.GetStoredVoting(ctx, "1")
	require.True(t, found)

	res, err := msgServer.AddVoter(ctx, &types.MsgAddVoter{
		Creator:  storedVoting.Creator,
		VotingId: storedVoting.Index,
		Dni:      "12345678",
		ProofId:  "fakeProofId",
	})
	require.Nil(t, err)
	require.EqualValues(t, res.Status, "ok")
}

func TestWrong1AddVoter(t *testing.T) {
	keeper, msgServer, ctx := setupMsgServerAddVoter(t)

	storedVoting, found := keeper.GetStoredVoting(ctx, "1")
	require.True(t, found)

	_, err := msgServer.AddVoter(ctx, &types.MsgAddVoter{
		Creator:  "invalid creator",
		VotingId: storedVoting.Index,
		Dni:      "12345678",
		ProofId:  "fakeProofId",
	})
	require.EqualValues(t, err, sdkerrors.ErrUnauthorized)
}
