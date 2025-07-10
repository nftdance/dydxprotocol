package accountplus

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nftdance/dydxprotocol/x/accountplus/keeper"
	"github.com/nftdance/dydxprotocol/x/accountplus/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) {
	err := k.SetGenesisState(ctx, data)
	if err != nil {
		panic(err)
	}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	accounts, err := k.GetAllAccountStates(ctx)
	if err != nil {
		panic(err)
	}

	params := k.GetParams(ctx)
	nextAuthenticatorId := k.InitializeOrGetNextAuthenticatorId(ctx)

	data, err := k.GetAllAuthenticatorData(ctx)
	if err != nil {
		panic(err)
	}

	return &types.GenesisState{
		Accounts:            accounts,
		Params:              params,
		NextAuthenticatorId: nextAuthenticatorId,
		AuthenticatorData:   data,
	}
}
