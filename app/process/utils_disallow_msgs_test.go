package process_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nftdance/dydxprotocol/app/msgs"
	"github.com/nftdance/dydxprotocol/app/process"
	"github.com/nftdance/dydxprotocol/lib"
	"github.com/nftdance/dydxprotocol/testutil/constants"
	clobtypes "github.com/nftdance/dydxprotocol/x/clob/types"
	"github.com/stretchr/testify/require"
)

func TestIsDisallowClobOrderMsgInOtherTxs_Empty(t *testing.T) {
	require.False(t, process.IsDisallowClobOrderMsgInOtherTxs(nil))
}

func TestIsDisallowClobOrderMsgInOtherTxs(t *testing.T) {
	allMsgSamples := lib.MergeAllMapsMustHaveDistinctKeys(
		msgs.AllowMsgs,
		msgs.DisallowMsgs,
	)

	for _, msg := range allMsgSamples {
		result := process.IsDisallowClobOrderMsgInOtherTxs(msg)
		switch msg.(type) {
		case *clobtypes.MsgCancelOrder, *clobtypes.MsgPlaceOrder, *clobtypes.MsgBatchCancel:
			// The sample msgs are short-term orders, so we expect these to be disallowed.
			require.True(t, result) // true -> disallow
		default:
			require.False(t, result) // false -> not disallow -> allow
		}
	}

	// Long-term orders are allowed.
	longTermOrders := []sdk.Msg{
		constants.Msg_PlaceOrder_LongTerm,
		constants.Msg_CancelOrder_LongTerm,
	}
	for _, msg := range longTermOrders {
		result := process.IsDisallowClobOrderMsgInOtherTxs(msg)
		require.False(t, result) // false -> not disallow -> allow
	}
}
