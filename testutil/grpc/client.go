package grpc

import (
	bridgetypes "github.com/nftdance/dydxprotocol/daemons/bridge/api"
	liquidationtypes "github.com/nftdance/dydxprotocol/daemons/liquidation/api"
	pricefeedtypes "github.com/nftdance/dydxprotocol/daemons/pricefeed/api"
	blocktimetypes "github.com/nftdance/dydxprotocol/x/blocktime/types"
	clobtypes "github.com/nftdance/dydxprotocol/x/clob/types"
	perptypes "github.com/nftdance/dydxprotocol/x/perpetuals/types"
	pricetypes "github.com/nftdance/dydxprotocol/x/prices/types"
	satypes "github.com/nftdance/dydxprotocol/x/subaccounts/types"
)

// QueryClient combines all the query clients used in testing into a single mock interface for testing convenience.
type QueryClient interface {
	blocktimetypes.QueryClient
	satypes.QueryClient
	clobtypes.QueryClient
	perptypes.QueryClient
	pricetypes.QueryClient
	bridgetypes.BridgeServiceClient
	liquidationtypes.LiquidationServiceClient
	pricefeedtypes.PriceFeedServiceClient
}
