package constants

import (
	"math/big"

	clobtypes "github.com/nftdance/dydxprotocol/x/clob/types"
	satypes "github.com/nftdance/dydxprotocol/x/subaccounts/types"
)

var (
	// Get state position functions.
	GetStatePosition_ZeroPositionSize = func(
		subaccountId satypes.SubaccountId,
		clobPairId clobtypes.ClobPairId,
	) (
		statePositionSize *big.Int,
	) {
		return big.NewInt(0)
	}
)
