package keeper

import (
	"context"
	"errors"
	"fmt"

	"github.com/nftdance/dydxprotocol/lib"
	perptypes "github.com/nftdance/dydxprotocol/x/perpetuals/types"
	"github.com/nftdance/dydxprotocol/x/subaccounts/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CollateralPoolAddress(
	c context.Context,
	req *types.QueryCollateralPoolAddressRequest,
) (*types.QueryCollateralPoolAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := lib.UnwrapSDKContext(c, types.ModuleName)

	collateralPool, err := k.GetCollateralPoolFromPerpetualId(
		ctx,
		req.PerpetualId,
	)
	if err != nil {
		if errors.Is(err, perptypes.ErrPerpetualDoesNotExist) {
			return nil,
				status.Error(
					codes.NotFound,
					fmt.Sprintf(
						"Perpetual id %+v not found.",
						req.PerpetualId,
					),
				)
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryCollateralPoolAddressResponse{
		CollateralPoolAddress: collateralPool.String(),
	}, nil
}
