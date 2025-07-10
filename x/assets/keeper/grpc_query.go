package keeper

import (
	"github.com/nftdance/dydxprotocol/x/assets/types"
)

var _ types.QueryServer = Keeper{}
