package keeper_test

import (
	"context"
	"testing"

	testapp "github.com/nftdance/dydxprotocol/testutil/app"
	"github.com/nftdance/dydxprotocol/x/revshare/keeper"
	"github.com/nftdance/dydxprotocol/x/revshare/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t *testing.T) (keeper.Keeper, types.MsgServer, context.Context) {
	tApp := testapp.NewTestAppBuilder(t).Build()
	ctx := tApp.InitChain()
	k := tApp.App.RevShareKeeper

	return k, keeper.NewMsgServerImpl(k), ctx
}

func TestMsgServer(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	require.NotNil(t, k)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
}
