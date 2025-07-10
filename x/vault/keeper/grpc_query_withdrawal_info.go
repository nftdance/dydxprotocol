package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/nftdance/dydxprotocol/dtypes"
	"github.com/nftdance/dydxprotocol/lib"
	"github.com/nftdance/dydxprotocol/x/vault/types"
)

func (k Keeper) MegavaultWithdrawalInfo(
	goCtx context.Context,
	req *types.QueryMegavaultWithdrawalInfoRequest,
) (*types.QueryMegavaultWithdrawalInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := lib.UnwrapSDKContext(goCtx, types.ModuleName)

	redeemedQuoteQuantums, megavaultEquity, totalShares, err := k.RedeemFromMainAndSubVaults(
		ctx,
		req.SharesToWithdraw.NumShares.BigInt(),
		true, // simulate the redemption (even though queries do execute on branched contexts).
	)
	if err != nil {
		return nil, err
	}

	return &types.QueryMegavaultWithdrawalInfoResponse{
		SharesToWithdraw:      req.SharesToWithdraw,
		ExpectedQuoteQuantums: dtypes.NewIntFromBigInt(redeemedQuoteQuantums),
		MegavaultEquity:       dtypes.NewIntFromBigInt(megavaultEquity),
		TotalShares: types.NumShares{
			NumShares: dtypes.NewIntFromBigInt(totalShares),
		},
	}, nil
}
