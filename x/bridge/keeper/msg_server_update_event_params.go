package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/nftdance/dydxprotocol/lib"
	"github.com/nftdance/dydxprotocol/x/bridge/types"
)

// UpdateEventParams updates the EventParams in state.
func (k msgServer) UpdateEventParams(
	goCtx context.Context,
	msg *types.MsgUpdateEventParams,
) (*types.MsgUpdateEventParamsResponse, error) {
	if !k.Keeper.HasAuthority(msg.GetAuthority()) {
		return nil, errorsmod.Wrapf(
			types.ErrInvalidAuthority,
			"message authority %s is not valid for sending update event params messages",
			msg.GetAuthority(),
		)
	}

	ctx := lib.UnwrapSDKContext(goCtx, types.ModuleName)

	if err := k.Keeper.UpdateEventParams(ctx, msg.Params); err != nil {
		return nil, err
	}

	return &types.MsgUpdateEventParamsResponse{}, nil
}
