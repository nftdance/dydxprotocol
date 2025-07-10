package keeper

import (
	"github.com/nftdance/dydxprotocol/x/perpetuals/types"
)

var _ types.QueryServer = Keeper{}
