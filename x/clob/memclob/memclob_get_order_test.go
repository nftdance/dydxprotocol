package memclob

import (
	"testing"

	"github.com/nftdance/dydxprotocol/testutil/constants"
	"github.com/nftdance/dydxprotocol/x/clob/types"
	satypes "github.com/nftdance/dydxprotocol/x/subaccounts/types"
	"github.com/stretchr/testify/require"
)

func TestGetOrder_Success(t *testing.T) {
	memclob := NewMemClobPriceTimePriority(false)

	orderId := types.OrderId{
		SubaccountId: satypes.SubaccountId{
			Owner: "testGetOrder",
		},
	}
	order := types.Order{OrderId: orderId}

	memclob.CreateOrderbook(constants.ClobPair_Btc)
	memclob.orderbooks[order.GetClobPairId()].orderIdToLevelOrder[orderId] = &types.LevelOrder{
		Value: types.ClobOrder{
			Order: order,
		},
	}

	foundOrder, found := memclob.GetOrder(orderId)
	require.True(t, found)
	require.Equal(t, order, foundOrder)
}

func TestGetOrder_ErrDoesNotExist(t *testing.T) {
	memclob := NewMemClobPriceTimePriority(false)

	memclob.CreateOrderbook(constants.ClobPair_Btc)
	_, found := memclob.GetOrder(types.OrderId{})
	require.False(t, found)
}
