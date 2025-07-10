package kraken_test

import (
	"testing"

	"github.com/nftdance/dydxprotocol/daemons/pricefeed/client/price_function/kraken"
	"github.com/stretchr/testify/require"
)

func TestKrakenUrl(t *testing.T) {
	require.Equal(t, "https://api.kraken.com/0/public/Ticker", kraken.KrakenDetails.Url)
}

func TestKrakenIsMultiMarket(t *testing.T) {
	require.True(t, kraken.KrakenDetails.IsMultiMarket)
}
