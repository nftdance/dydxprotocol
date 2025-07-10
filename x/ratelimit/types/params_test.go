package types_test

import (
	"testing"
	"time"

	"github.com/nftdance/dydxprotocol/dtypes"
	"github.com/nftdance/dydxprotocol/x/ratelimit/types"
	"github.com/stretchr/testify/require"
)

func TestDefaultUsdcRateLimitParams(t *testing.T) {
	require.Equal(t,
		types.LimitParams{
			Denom: "ibc/8E27BA2D5493AF5636760E354E46004562C46AB7EC0CC4C1CA14E9E20E2545B5",
			Limiters: []types.Limiter{
				{
					Period:          3600 * time.Second,
					BaselineMinimum: dtypes.NewInt(1_000_000_000_000),
					BaselineTvlPpm:  10_000,
				},
				{
					Period:          24 * time.Hour,
					BaselineMinimum: dtypes.NewInt(10_000_000_000_000),
					BaselineTvlPpm:  100_000,
				},
			},
		},
		types.DefaultUsdcRateLimitParams(),
	)
}
