package subaccounts

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	indexerevents "github.com/nftdance/dydxprotocol/indexer/events"
	"github.com/nftdance/dydxprotocol/indexer/indexer_manager"
	"github.com/nftdance/dydxprotocol/x/subaccounts/keeper"
	"github.com/nftdance/dydxprotocol/x/subaccounts/types"
)

// InitGenesis initializes the subaccounts module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.InitializeForGenesis(ctx)

	// Set all the subaccounts
	for _, elem := range genState.Subaccounts {
		k.SetSubaccount(ctx, elem)
		k.GetIndexerEventManager().AddTxnEvent(
			ctx,
			indexerevents.SubtypeSubaccountUpdate,
			indexerevents.SubaccountUpdateEventVersion,
			indexer_manager.GetBytes(
				indexerevents.NewSubaccountUpdateEvent(
					elem.Id,
					elem.PerpetualPositions,
					elem.AssetPositions,
					nil,
				),
			),
		)
	}
}

// ExportGenesis returns the subaccounts module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.Subaccounts = k.GetAllSubaccount(ctx)

	return genesis
}
