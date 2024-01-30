package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/emaforlin/mAVS/testutil/keeper"
	"github.com/emaforlin/mAVS/x/mavs/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.MavsKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
