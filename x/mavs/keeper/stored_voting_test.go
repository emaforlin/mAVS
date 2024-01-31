package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "github.com/emaforlin/mAVS/testutil/keeper"
	"github.com/emaforlin/mAVS/testutil/nullify"
	"github.com/emaforlin/mAVS/x/mavs/keeper"
	"github.com/emaforlin/mAVS/x/mavs/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNStoredVoting(keeper keeper.Keeper, ctx context.Context, n int) []types.StoredVoting {
	items := make([]types.StoredVoting, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetStoredVoting(ctx, items[i])
	}
	return items
}

func TestStoredVotingGet(t *testing.T) {
	keeper, ctx := keepertest.MavsKeeper(t)
	items := createNStoredVoting(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetStoredVoting(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestStoredVotingRemove(t *testing.T) {
	keeper, ctx := keepertest.MavsKeeper(t)
	items := createNStoredVoting(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveStoredVoting(ctx,
			item.Index,
		)
		_, found := keeper.GetStoredVoting(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestStoredVotingGetAll(t *testing.T) {
	keeper, ctx := keepertest.MavsKeeper(t)
	items := createNStoredVoting(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllStoredVoting(ctx)),
	)
}
