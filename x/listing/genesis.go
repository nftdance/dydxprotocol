package listing

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nftdance/dydxprotocol/x/listing/keeper"
	"github.com/nftdance/dydxprotocol/x/listing/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.InitializeForGenesis(ctx)

	if err := k.SetMarketsHardCap(ctx, genState.HardCapForMarkets); err != nil {
		panic(err)
	}

	if err := k.SetListingVaultDepositParams(ctx, genState.ListingVaultDepositParams); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	return genesis
}
