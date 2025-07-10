package ratelimit_test

import (
	"testing"

	testapp "github.com/nftdance/dydxprotocol/testutil/app"
	"github.com/nftdance/dydxprotocol/x/ratelimit"
	"github.com/nftdance/dydxprotocol/x/ratelimit/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	tApp := testapp.NewTestAppBuilder(t).Build()
	ctx := tApp.InitChain()
	got := ratelimit.ExportGenesis(ctx, tApp.App.RatelimitKeeper)
	require.NotNil(t, got)
	require.Equal(t, types.DefaultGenesis(), got)
}
