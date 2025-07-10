package keeper

import (
	"github.com/nftdance/dydxprotocol/x/sending/types"
)

type msgServer struct {
	Keeper types.SendingKeeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper types.SendingKeeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}
