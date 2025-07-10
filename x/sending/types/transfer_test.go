package types_test

import (
	"math/big"
	"testing"

	"github.com/nftdance/dydxprotocol/testutil/constants"
	testutil "github.com/nftdance/dydxprotocol/testutil/util"
	"github.com/nftdance/dydxprotocol/x/sending/types"
	satypes "github.com/nftdance/dydxprotocol/x/subaccounts/types"

	"github.com/stretchr/testify/require"
)

func TestGetBigQuoteQuantums(t *testing.T) {
	transfer := constants.Transfer_Carl_Num0_Dave_Num0_Quote_500
	bigQuoteQuantums := transfer.GetBigQuantums()
	require.Equal(t, new(big.Int).SetUint64(500_000_000), bigQuoteQuantums)
}

func TestGetSubaccountUpdates(t *testing.T) {
	tests := map[string]struct {
		transfer types.Transfer
		expected []satypes.Update
	}{
		"Test subaccount updates": {
			transfer: constants.Transfer_Carl_Num0_Dave_Num0_Quote_500,
			expected: []satypes.Update{
				{
					SubaccountId: constants.Carl_Num0,
					AssetUpdates: testutil.CreateUsdcAssetUpdates(big.NewInt(-500_000_000)),
				},
				{
					SubaccountId: constants.Dave_Num0,
					AssetUpdates: testutil.CreateUsdcAssetUpdates(big.NewInt(500_000_000)),
				},
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result := []satypes.Update{
				tc.transfer.GetSenderSubaccountUpdate(),
				tc.transfer.GetRecipientSubaccountUpdate(),
			}
			require.Equal(t, tc.expected, result)
		})
	}
}
