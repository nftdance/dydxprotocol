package keeper

import (
	dbm "github.com/cosmos/cosmos-db"
	"github.com/nftdance/dydxprotocol/lib"
	"github.com/nftdance/dydxprotocol/mocks"
	delaymsgtypes "github.com/nftdance/dydxprotocol/x/delaymsg/types"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	epochskeeper "github.com/nftdance/dydxprotocol/x/epochs/keeper"
	"github.com/nftdance/dydxprotocol/x/stats/keeper"
	"github.com/nftdance/dydxprotocol/x/stats/types"
)

func createStatsKeeper(
	stateStore storetypes.CommitMultiStore,
	epochsKeeper *epochskeeper.Keeper,
	db *dbm.MemDB,
	cdc *codec.ProtoCodec,
	stakingKeeper *stakingkeeper.Keeper,
) (*keeper.Keeper, storetypes.StoreKey) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)
	transientStoreKey := storetypes.NewTransientStoreKey(types.TransientStoreKey)

	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(transientStoreKey, storetypes.StoreTypeTransient, db)

	mockMsgSender := &mocks.IndexerMessageSender{}
	mockMsgSender.On("Enabled").Return(true)

	authorities := []string{
		delaymsgtypes.ModuleAddress.String(),
		lib.GovModuleAddress.String(),
	}
	k := keeper.NewKeeper(
		cdc,
		epochsKeeper,
		storeKey,
		transientStoreKey,
		authorities,
		stakingKeeper,
	)

	return k, storeKey
}
