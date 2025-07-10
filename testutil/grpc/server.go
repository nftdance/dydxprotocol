package grpc

import pricetypes "github.com/nftdance/dydxprotocol/x/prices/types"

type QueryServer interface {
	pricetypes.QueryServer
}
