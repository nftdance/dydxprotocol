package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	bridgetypes "github.com/nftdance/dydxprotocol/x/bridge/types"
	clobtypes "github.com/nftdance/dydxprotocol/x/clob/types"
	perpetualstypes "github.com/nftdance/dydxprotocol/x/perpetuals/types"
	pricestypes "github.com/nftdance/dydxprotocol/x/prices/types"
)

// IsSingleAppInjectedMsg returns true if the given list of msgs contains an "app-injected msg"
// and it's the only msg in the list. Otherwise, returns false.
func IsSingleAppInjectedMsg(msgs []sdk.Msg) bool {
	return len(msgs) == 1 && IsAppInjectedMsg(msgs[0])
}

// IsAppInjectedMsg returns true if the given msg is an "app-injected msg".
// Otherwise, returns false.
func IsAppInjectedMsg(msg sdk.Msg) bool {
	switch msg.(type) {
	case
		// ------- Custom modules
		// bridge
		*bridgetypes.MsgAcknowledgeBridges,

		// clob
		*clobtypes.MsgProposedOperations,

		// perpetuals
		*perpetualstypes.MsgAddPremiumVotes,

		// prices
		*pricestypes.MsgUpdateMarketPrices:

		return true
	}
	return false
}
