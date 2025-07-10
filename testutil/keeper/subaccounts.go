package keeper

import (
	"testing"

	"github.com/nftdance/dydxprotocol/streaming"

	affiliateskeeper "github.com/nftdance/dydxprotocol/x/affiliates/keeper"
	revsharekeeper "github.com/nftdance/dydxprotocol/x/revshare/keeper"

	"github.com/cosmos/gogoproto/proto"

	dbm "github.com/cosmos/cosmos-db"

	"github.com/nftdance/dydxprotocol/testutil/constants"

	indexerevents "github.com/nftdance/dydxprotocol/indexer/events"
	"github.com/nftdance/dydxprotocol/indexer/indexer_manager"
	"github.com/nftdance/dydxprotocol/mocks"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	asskeeper "github.com/nftdance/dydxprotocol/x/assets/keeper"
	blocktimekeeper "github.com/nftdance/dydxprotocol/x/blocktime/keeper"
	perpskeeper "github.com/nftdance/dydxprotocol/x/perpetuals/keeper"
	priceskeeper "github.com/nftdance/dydxprotocol/x/prices/keeper"
	"github.com/nftdance/dydxprotocol/x/subaccounts/keeper"
	"github.com/nftdance/dydxprotocol/x/subaccounts/types"
)

func SubaccountsKeepers(t testing.TB, msgSenderEnabled bool) (
	ctx sdk.Context,
	keeper *keeper.Keeper,
	pricesKeeper *priceskeeper.Keeper,
	perpetualsKeeper *perpskeeper.Keeper,
	accountKeeper *authkeeper.AccountKeeper,
	bankKeeper *bankkeeper.BaseKeeper,
	assetsKeeper *asskeeper.Keeper,
	blocktimeKeeper *blocktimekeeper.Keeper,
	revShareKeeper *revsharekeeper.Keeper,
	affiliatesKeeper *affiliateskeeper.Keeper,
	storeKey storetypes.StoreKey,
) {
	var mockTimeProvider *mocks.TimeProvider
	ctx = initKeepers(t, func(
		db *dbm.MemDB,
		registry codectypes.InterfaceRegistry,
		cdc *codec.ProtoCodec,
		stateStore storetypes.CommitMultiStore,
		transientStoreKey storetypes.StoreKey,
	) []GenesisInitializer {
		// Define necessary keepers here for unit tests
		epochsKeeper, _ := createEpochsKeeper(stateStore, db, cdc)

		accountKeeper, _ = createAccountKeeper(
			stateStore,
			db,
			cdc,
			registry)
		bankKeeper, _ = createBankKeeper(stateStore, db, cdc, accountKeeper)
		stakingKeeper, _ := createStakingKeeper(
			stateStore,
			db,
			cdc,
			accountKeeper,
			bankKeeper,
		)
		statsKeeper, _ := createStatsKeeper(
			stateStore,
			epochsKeeper,
			db,
			cdc,
			stakingKeeper,
		)
		affiliatesKeeper, _ = createAffiliatesKeeper(stateStore, db, cdc, statsKeeper, transientStoreKey, true)
		vaultKeeper, _ := createVaultKeeper(stateStore, db, cdc, transientStoreKey)
		feetiersKeeper, _ := createFeeTiersKeeper(stateStore, statsKeeper, vaultKeeper, affiliatesKeeper, db, cdc)
		revShareKeeper, _, _ = createRevShareKeeper(stateStore, db, cdc, affiliatesKeeper, feetiersKeeper)
		marketMapKeeper, _ := createMarketMapKeeper(stateStore, db, cdc)
		pricesKeeper, _, _, mockTimeProvider = createPricesKeeper(
			stateStore,
			db,
			cdc,
			transientStoreKey,
			revShareKeeper,
			marketMapKeeper,
		)
		perpetualsKeeper, _ = createPerpetualsKeeper(stateStore, db, cdc, pricesKeeper, epochsKeeper, transientStoreKey)
		assetsKeeper, _ = createAssetsKeeper(stateStore, db, cdc, pricesKeeper, transientStoreKey, msgSenderEnabled)
		blocktimeKeeper, _ = createBlockTimeKeeper(stateStore, db, cdc)

		keeper, storeKey = createSubaccountsKeeper(
			stateStore,
			db,
			cdc,
			assetsKeeper,
			bankKeeper,
			perpetualsKeeper,
			blocktimeKeeper,
			transientStoreKey,
			msgSenderEnabled,
		)

		return []GenesisInitializer{pricesKeeper, perpetualsKeeper, assetsKeeper, revShareKeeper, affiliatesKeeper, keeper}
	})

	// Mock time provider response for market creation.
	mockTimeProvider.On("Now").Return(constants.TimeT)

	return ctx,
		keeper,
		pricesKeeper,
		perpetualsKeeper,
		accountKeeper,
		bankKeeper,
		assetsKeeper,
		blocktimeKeeper,
		revShareKeeper,
		affiliatesKeeper,
		storeKey
}

func createSubaccountsKeeper(
	stateStore storetypes.CommitMultiStore,
	db *dbm.MemDB,
	cdc *codec.ProtoCodec,
	ak *asskeeper.Keeper,
	bk types.BankKeeper,
	pk *perpskeeper.Keeper,
	btk *blocktimekeeper.Keeper,
	transientStoreKey storetypes.StoreKey,
	msgSenderEnabled bool,
) (*keeper.Keeper, storetypes.StoreKey) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)

	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)

	mockMsgSender := &mocks.IndexerMessageSender{}
	mockMsgSender.On("Enabled").Return(msgSenderEnabled)
	mockIndexerEventsManager := indexer_manager.NewIndexerEventManager(mockMsgSender, transientStoreKey, true)

	k := keeper.NewKeeper(
		cdc,
		storeKey,
		ak,
		bk,
		pk,
		btk,
		mockIndexerEventsManager,
		streaming.NewNoopGrpcStreamingManager(),
	)

	return k, storeKey
}

// GetSubaccountUpdateEventsFromIndexerBlock returns the subaccount update events in the
// Indexer Block event Kafka message.
func GetSubaccountUpdateEventsFromIndexerBlock(
	ctx sdk.Context,
	keeper *keeper.Keeper,
) []*indexerevents.SubaccountUpdateEventV1 {
	var subaccountUpdates []*indexerevents.SubaccountUpdateEventV1
	block := keeper.GetIndexerEventManager().ProduceBlock(ctx)
	if block == nil {
		return subaccountUpdates
	}
	for _, event := range block.Events {
		if event.Subtype != indexerevents.SubtypeSubaccountUpdate {
			continue
		}
		var subaccountUpdate indexerevents.SubaccountUpdateEventV1
		err := proto.Unmarshal(event.DataBytes, &subaccountUpdate)
		if err != nil {
			panic(err)
		}
		subaccountUpdates = append(subaccountUpdates, &subaccountUpdate)
	}
	return subaccountUpdates
}
