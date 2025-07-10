package util

import (
	"math/big"

	"github.com/nftdance/dydxprotocol/lib"
	"github.com/nftdance/dydxprotocol/x/ratelimit/types"
)

// GetBaseline returns the current capacity baseline for the given limiter.
// `baseline` formula:
//
//	baseline = max(baseline_minimum, baseline_tvl_ppm * current_tvl)
func GetBaseline(
	currentTvl *big.Int,
	limiter types.Limiter,
) *big.Int {
	return lib.BigMax(
		limiter.BaselineMinimum.BigInt(),
		lib.BigIntMulPpm(
			currentTvl,
			limiter.BaselineTvlPpm,
		),
	)
}
