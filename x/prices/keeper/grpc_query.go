package keeper

import (
	"github.com/nftdance/dydxprotocol/x/prices/types"
)

var _ types.QueryServer = Keeper{}
