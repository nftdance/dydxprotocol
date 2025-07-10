package rate_limit_test

import (
	testapp "github.com/nftdance/dydxprotocol/testutil/app"
	"github.com/nftdance/dydxprotocol/x/clob/rate_limit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPanicRateLimiter(t *testing.T) {
	tApp := testapp.NewTestAppBuilder(t).Build()
	ctx := tApp.InitChain()
	rl := rate_limit.NewPanicRateLimiter[int]()
	require.Panics(t, func() {
		//nolint:errcheck
		rl.RateLimit(ctx, 5)
	})
	require.Panics(t, func() {
		//nolint:errcheck
		rl.PruneRateLimits(ctx)
	})
}
