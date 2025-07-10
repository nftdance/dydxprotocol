package keeper

import (
	"github.com/nftdance/dydxprotocol/x/delaymsg/types"
)

var _ types.QueryServer = Keeper{}
