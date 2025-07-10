package huobi_test

import (
	"testing"

	"github.com/nftdance/dydxprotocol/daemons/pricefeed/client/price_function/huobi"
	"github.com/stretchr/testify/require"
)

func TestHuobiUrl(t *testing.T) {
	require.Equal(t, "https://api.huobi.pro/market/tickers", huobi.HuobiDetails.Url)
}

func TestHuobiIsMultiMarket(t *testing.T) {
	require.True(t, huobi.HuobiDetails.IsMultiMarket)
}
