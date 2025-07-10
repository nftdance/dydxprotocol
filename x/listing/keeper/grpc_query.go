package keeper

import (
	"github.com/nftdance/dydxprotocol/x/listing/types"
)

var _ types.QueryServer = Keeper{}
