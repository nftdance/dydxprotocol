package delaymsg_test

import (
	"testing"
	"time"

	"github.com/nftdance/dydxprotocol/mocks"
	sdktest "github.com/nftdance/dydxprotocol/testutil/sdk"
	"github.com/nftdance/dydxprotocol/x/delaymsg"
	"github.com/nftdance/dydxprotocol/x/delaymsg/types"
)

func TestEndBlocker(t *testing.T) {
	k := &mocks.DelayMsgKeeper{}
	ctx := sdktest.NewContextWithBlockHeightAndTime(0, time.Now())
	// When DispatchMessagesForBlock is called, GetBlockMessageIds will be called, so we expect this call.
	// Return an empty list of message IDs and a false value to indicate that there are no messages.
	// In this case, the method should immediately return.
	k.On("GetBlockMessageIds", ctx, uint32(0)).Return(types.BlockMessageIds{}, false).Once()
	delaymsg.EndBlocker(ctx, k)
	k.AssertExpectations(t)
}
