package delaymsg

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/nftdance/dydxprotocol/x/delaymsg/keeper"
	"github.com/nftdance/dydxprotocol/x/delaymsg/types"
)

// EndBlocker executes all ABCI EndBlock logic respective to the delaymsg module.
func EndBlocker(ctx sdk.Context, k types.DelayMsgKeeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)
	keeper.DispatchMessagesForBlock(k, ctx)
}
