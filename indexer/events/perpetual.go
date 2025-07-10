package events

import (
	v1 "github.com/nftdance/dydxprotocol/indexer/protocol/v1"
	perptypes "github.com/nftdance/dydxprotocol/x/perpetuals/types"
)

// NewUpdatePerpetualEvent creates a UpdatePerpetualEventV3 representing
// update of a perpetual.
func NewUpdatePerpetualEvent(
	id uint32,
	ticker string,
	marketId uint32,
	atomicResolution int32,
	liquidityTier uint32,
	marketType perptypes.PerpetualMarketType,
	defaultFundingPpm int32,
) *UpdatePerpetualEventV3 {
	return &UpdatePerpetualEventV3{
		Id:                   id,
		Ticker:               ticker,
		MarketId:             marketId,
		AtomicResolution:     atomicResolution,
		LiquidityTier:        liquidityTier,
		MarketType:           v1.ConvertToPerpetualMarketType(marketType),
		DefaultFunding8HrPpm: defaultFundingPpm,
	}
}
