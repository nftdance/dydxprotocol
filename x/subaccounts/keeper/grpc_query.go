package keeper

import (
	"github.com/nftdance/dydxprotocol/x/subaccounts/types"
)

var _ types.QueryServer = Keeper{}
