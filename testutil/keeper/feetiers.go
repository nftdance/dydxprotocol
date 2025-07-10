package keeper

import (
	dbm "github.com/cosmos/cosmos-db"
	"github.com/nftdance/dydxprotocol/lib"
	"github.com/nftdance/dydxprotocol/mocks"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	affiliateskeeper "github.com/nftdance/dydxprotocol/x/affiliates/keeper"
	delaymsgtypes "github.com/nftdance/dydxprotocol/x/delaymsg/types"
	"github.com/nftdance/dydxprotocol/x/feetiers/keeper"
	"github.com/nftdance/dydxprotocol/x/feetiers/types"
	statskeeper "github.com/nftdance/dydxprotocol/x/stats/keeper"
	vaultkeeper "github.com/nftdance/dydxprotocol/x/vault/keeper"
)

func createFeeTiersKeeper(
	stateStore storetypes.CommitMultiStore,
	statsKeeper *statskeeper.Keeper,
	vaultKeeper *vaultkeeper.Keeper,
	affiliatesKeeper *affiliateskeeper.Keeper,
	db *dbm.MemDB,
	cdc *codec.ProtoCodec,
) (*keeper.Keeper, storetypes.StoreKey) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)

	mockMsgSender := &mocks.IndexerMessageSender{}
	mockMsgSender.On("Enabled").Return(true)

	authorities := []string{
		delaymsgtypes.ModuleAddress.String(),
		lib.GovModuleAddress.String(),
	}
	k := keeper.NewKeeper(
		cdc,
		statsKeeper,
		affiliatesKeeper,
		storeKey,
		authorities,
	)
	k.SetVaultKeeper(vaultKeeper)

	return k, storeKey
}
