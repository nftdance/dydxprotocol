package keeper

import (
	"github.com/nftdance/dydxprotocol/x/epochs/types"
)

var _ types.QueryServer = Keeper{}
