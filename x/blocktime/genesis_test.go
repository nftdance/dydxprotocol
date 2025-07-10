package blocktime_test

import (
	"testing"

	testapp "github.com/nftdance/dydxprotocol/testutil/app"
	"github.com/nftdance/dydxprotocol/x/blocktime"
	"github.com/nftdance/dydxprotocol/x/blocktime/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	tApp := testapp.NewTestAppBuilder(t).Build()
	ctx := tApp.InitChain()
	got := blocktime.ExportGenesis(ctx, tApp.App.BlockTimeKeeper)
	require.NotNil(t, got)
	require.Equal(t, types.DefaultGenesis(), got)
}
