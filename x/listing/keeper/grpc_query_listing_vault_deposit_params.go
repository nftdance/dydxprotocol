package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nftdance/dydxprotocol/x/listing/types"
)

func (k Keeper) ListingVaultDepositParams(
	ctx context.Context,
	req *types.QueryListingVaultDepositParams,
) (*types.QueryListingVaultDepositParamsResponse, error) {
	params := k.GetListingVaultDepositParams(sdk.UnwrapSDKContext(ctx))
	return &types.QueryListingVaultDepositParamsResponse{
		Params: params,
	}, nil
}
