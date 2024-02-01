package mavs_test

import (
	"testing"

	keepertest "github.com/emaforlin/mAVS/testutil/keeper"
	"github.com/emaforlin/mAVS/testutil/nullify"
	mavs "github.com/emaforlin/mAVS/x/mavs/module"
	"github.com/emaforlin/mAVS/x/mavs/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		SystemInfo: types.SystemInfo{
			NextId: 58,
		},
		StoredVotingList: []types.StoredVoting{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MavsKeeper(t)
	mavs.InitGenesis(ctx, k, genesisState)
	got := mavs.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.SystemInfo, got.SystemInfo)
	require.ElementsMatch(t, genesisState.StoredVotingList, got.StoredVotingList)
	// this line is used by starport scaffolding # genesis/test/assert
}
