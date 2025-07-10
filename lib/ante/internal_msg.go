package ante

import (
	upgrade "cosmossdk.io/x/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	consensus "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisis "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distribution "github.com/cosmos/cosmos-sdk/x/distribution/types"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	slashing "github.com/cosmos/cosmos-sdk/x/slashing/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	icahosttypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/host/types"
	ibctransfer "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	ibcclient "github.com/cosmos/ibc-go/v8/modules/core/02-client/types" //nolint:staticcheck
	ibcconn "github.com/cosmos/ibc-go/v8/modules/core/03-connection/types"
	accountplus "github.com/nftdance/dydxprotocol/x/accountplus/types"
	affiliates "github.com/nftdance/dydxprotocol/x/affiliates/types"
	blocktime "github.com/nftdance/dydxprotocol/x/blocktime/types"
	bridge "github.com/nftdance/dydxprotocol/x/bridge/types"
	clob "github.com/nftdance/dydxprotocol/x/clob/types"
	delaymsg "github.com/nftdance/dydxprotocol/x/delaymsg/types"
	feetiers "github.com/nftdance/dydxprotocol/x/feetiers/types"
	govplus "github.com/nftdance/dydxprotocol/x/govplus/types"
	listing "github.com/nftdance/dydxprotocol/x/listing/types"
	perpetuals "github.com/nftdance/dydxprotocol/x/perpetuals/types"
	prices "github.com/nftdance/dydxprotocol/x/prices/types"
	ratelimit "github.com/nftdance/dydxprotocol/x/ratelimit/types"
	revshare "github.com/nftdance/dydxprotocol/x/revshare/types"
	rewards "github.com/nftdance/dydxprotocol/x/rewards/types"
	sending "github.com/nftdance/dydxprotocol/x/sending/types"
	stats "github.com/nftdance/dydxprotocol/x/stats/types"
	vault "github.com/nftdance/dydxprotocol/x/vault/types"
	vest "github.com/nftdance/dydxprotocol/x/vest/types"
)

// IsInternalMsg returns true if the given msg is an internal message.
func IsInternalMsg(msg sdk.Msg) bool {
	switch msg.(type) {
	case
		// ------- CosmosSDK default modules
		// auth
		*auth.MsgUpdateParams,

		// bank
		*bank.MsgSetSendEnabled,
		*bank.MsgUpdateParams,

		// consensus
		*consensus.MsgUpdateParams,

		// crisis
		*crisis.MsgUpdateParams,

		// distribution
		*distribution.MsgCommunityPoolSpend,
		*distribution.MsgUpdateParams,

		// gov
		*gov.MsgExecLegacyContent,
		*gov.MsgUpdateParams,

		// slashing
		*slashing.MsgUpdateParams,

		// staking
		*staking.MsgUpdateParams,

		// upgrade
		*upgrade.MsgCancelUpgrade,
		*upgrade.MsgSoftwareUpgrade,

		// ------- Custom modules
		// accountplus
		*accountplus.MsgSetActiveState,

		// blocktime
		*blocktime.MsgUpdateDowntimeParams,
		*blocktime.MsgUpdateSynchronyParams,
		// bridge
		*bridge.MsgCompleteBridge,
		*bridge.MsgUpdateEventParams,
		*bridge.MsgUpdateProposeParams,
		*bridge.MsgUpdateSafetyParams,

		// clob
		*clob.MsgCreateClobPair,
		*clob.MsgUpdateBlockRateLimitConfiguration,
		*clob.MsgUpdateClobPair,
		*clob.MsgUpdateEquityTierLimitConfiguration,
		*clob.MsgUpdateLiquidationsConfig,

		// delaymsg
		*delaymsg.MsgDelayMessage,

		// feetiers
		*feetiers.MsgUpdatePerpetualFeeParams,

		// govplus
		*govplus.MsgSlashValidator,

		// listing
		*listing.MsgSetMarketsHardCap,
		*listing.MsgSetListingVaultDepositParams,
		*listing.MsgUpgradeIsolatedPerpetualToCross,

		// perpetuals
		*perpetuals.MsgCreatePerpetual,
		*perpetuals.MsgSetLiquidityTier,
		*perpetuals.MsgUpdateParams,
		*perpetuals.MsgUpdatePerpetualParams,

		// prices
		*prices.MsgCreateOracleMarket,
		*prices.MsgUpdateMarketParam,

		// ratelimit
		*ratelimit.MsgSetLimitParams,
		*ratelimit.MsgSetLimitParamsResponse,

		// revshare
		*revshare.MsgSetMarketMapperRevenueShare,
		*revshare.MsgSetMarketMapperRevShareDetailsForMarket,
		*revshare.MsgUpdateUnconditionalRevShareConfig,

		// rewards
		*rewards.MsgUpdateParams,

		// sending
		*sending.MsgSendFromModuleToAccount,

		// stats
		*stats.MsgUpdateParams,

		// vault
		*vault.MsgUnlockShares,
		*vault.MsgUpdateDefaultQuotingParams,
		*vault.MsgUpdateOperatorParams,

		// vest
		*vest.MsgDeleteVestEntry,
		*vest.MsgSetVestEntry,

		// ibc
		*icahosttypes.MsgUpdateParams,
		*ibctransfer.MsgUpdateParams,
		*ibcclient.MsgUpdateParams,
		*ibcconn.MsgUpdateParams,

		// affiliates
		*affiliates.MsgUpdateAffiliateTiers,
		*affiliates.MsgUpdateAffiliateWhitelist:

		return true

	default:
		return false
	}
}
