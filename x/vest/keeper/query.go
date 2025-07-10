package keeper

import (
	"github.com/nftdance/dydxprotocol/x/vest/types"
)

var _ types.QueryServer = Keeper{}
