package subaccounts_test

import (
	"math/big"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	keepertest "github.com/nftdance/dydxprotocol/testutil/keeper"
	"github.com/nftdance/dydxprotocol/testutil/nullify"
	testutil "github.com/nftdance/dydxprotocol/testutil/util"
	"github.com/nftdance/dydxprotocol/x/subaccounts"
	"github.com/nftdance/dydxprotocol/x/subaccounts/keeper"
	"github.com/nftdance/dydxprotocol/x/subaccounts/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Subaccounts: []types.Subaccount{
			{
				Id: &types.SubaccountId{
					Owner:  "foo",
					Number: uint32(0),
				},
				AssetPositions: testutil.CreateUsdcAssetPositions(big.NewInt(1_000)),
			},
			{
				Id: &types.SubaccountId{
					Owner:  "bar",
					Number: uint32(99),
				},
				AssetPositions: testutil.CreateUsdcAssetPositions(big.NewInt(1_000)),
			},
		},
	}

	ctx, k, _, _, _, _, _, _, _, _, _ := keepertest.SubaccountsKeepers(t, true)
	subaccounts.InitGenesis(ctx, *k, genesisState)
	assertSubaccountUpdateEventsInIndexerBlock(t, k, ctx, 2)
	got := subaccounts.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState) //nolint:staticcheck
	nullify.Fill(got)           //nolint:staticcheck

	require.ElementsMatch(t, genesisState.Subaccounts, got.Subaccounts)
}

// assertSubaccountUpdateEventsInIndexerBlock checks that the number of subaccount update events
// included in the Indexer block kafka message.
func assertSubaccountUpdateEventsInIndexerBlock(
	t *testing.T,
	k *keeper.Keeper,
	ctx sdk.Context,
	numSubaccounts int,
) {
	subaccountUpdates := keepertest.GetSubaccountUpdateEventsFromIndexerBlock(ctx, k)
	require.Len(t, subaccountUpdates, numSubaccounts)
}
