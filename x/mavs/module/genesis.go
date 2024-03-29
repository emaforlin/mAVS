package mavs

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/emaforlin/mAVS/x/mavs/keeper"
	"github.com/emaforlin/mAVS/x/mavs/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	k.SetSystemInfo(ctx, genState.SystemInfo)
	// Set all the storedVoting
	for _, elem := range genState.StoredVotingList {
		k.SetStoredVoting(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all systemInfo
	systemInfo, found := k.GetSystemInfo(ctx)
	if found {
		genesis.SystemInfo = systemInfo
	}
	genesis.StoredVotingList = k.GetAllStoredVoting(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
