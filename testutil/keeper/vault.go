package keeper

import (
	"testing"

	dbm "github.com/cosmos/cosmos-db"
	"github.com/nftdance/dydxprotocol/indexer/indexer_manager"
	"github.com/nftdance/dydxprotocol/lib"
	delaymsgtypes "github.com/nftdance/dydxprotocol/x/delaymsg/types"
	"github.com/stretchr/testify/mock"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nftdance/dydxprotocol/mocks"
	"github.com/nftdance/dydxprotocol/x/vault/keeper"
	"github.com/nftdance/dydxprotocol/x/vault/types"
)

func VaultKeepers(
	t testing.TB,
) (
	ctx sdk.Context,
	keeper *keeper.Keeper,
	storeKey storetypes.StoreKey,
) {
	ctx = initKeepers(t, func(
		db *dbm.MemDB,
		registry codectypes.InterfaceRegistry,
		cdc *codec.ProtoCodec,
		stateStore storetypes.CommitMultiStore,
		transientStoreKey storetypes.StoreKey,
	) []GenesisInitializer {
		// Define necessary keepers here for unit tests
		keeper, storeKey = createVaultKeeper(stateStore, db, cdc, transientStoreKey)
		return []GenesisInitializer{keeper}
	})

	return ctx, keeper, storeKey
}

func createVaultKeeper(
	stateStore storetypes.CommitMultiStore,
	db *dbm.MemDB,
	cdc *codec.ProtoCodec,
	transientStoreKey storetypes.StoreKey,
) (
	*keeper.Keeper,
	storetypes.StoreKey,
) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)

	mockMsgSender := &mocks.IndexerMessageSender{}
	mockMsgSender.On("Enabled").Return(true)
	mockMsgSender.On("SendOnchainData", mock.Anything).Return()
	mockMsgSender.On("SendOffchainData", mock.Anything).Return()

	mockIndexerEventsManager := indexer_manager.NewIndexerEventManager(mockMsgSender, transientStoreKey, true)

	k := keeper.NewKeeper(
		cdc,
		storeKey,
		&mocks.AssetsKeeper{},
		&mocks.BankKeeper{},
		&mocks.ClobKeeper{},
		&mocks.DelayMsgKeeper{},
		&mocks.PerpetualsKeeper{},
		&mocks.PricesKeeper{},
		&mocks.SendingKeeper{},
		&mocks.SubaccountsKeeper{},
		mockIndexerEventsManager,
		[]string{
			lib.GovModuleAddress.String(),
			delaymsgtypes.ModuleAddress.String(),
		},
	)

	return k, storeKey
}
