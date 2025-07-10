package keeper

import (
	"context"

	"github.com/nftdance/dydxprotocol/lib"
	"github.com/nftdance/dydxprotocol/x/clob/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) BlockRateLimitConfiguration(
	c context.Context,
	req *types.QueryBlockRateLimitConfigurationRequest,
) (*types.QueryBlockRateLimitConfigurationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := lib.UnwrapSDKContext(c, types.ModuleName)
	blockRateLimitConfig := k.GetBlockRateLimitConfiguration(ctx)

	return &types.QueryBlockRateLimitConfigurationResponse{
		BlockRateLimitConfig: blockRateLimitConfig,
	}, nil
}
