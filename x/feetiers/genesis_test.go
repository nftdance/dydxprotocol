package feetiers_test

import (
	"testing"

	testapp "github.com/nftdance/dydxprotocol/testutil/app"
	feetiers "github.com/nftdance/dydxprotocol/x/feetiers"
	"github.com/nftdance/dydxprotocol/x/feetiers/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	tApp := testapp.NewTestAppBuilder(t).Build()
	ctx := tApp.InitChain()
	got := feetiers.ExportGenesis(ctx, *tApp.App.FeeTiersKeeper)
	require.NotNil(t, got)
	require.Equal(t, types.DefaultGenesis(), got)
}
