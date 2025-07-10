package keeper

import (
	"github.com/nftdance/dydxprotocol/x/sending/types"
)

var _ types.QueryServer = Keeper{}
