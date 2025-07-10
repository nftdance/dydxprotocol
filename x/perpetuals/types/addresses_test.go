package types_test

import (
	"testing"

	"github.com/nftdance/dydxprotocol/x/perpetuals/types"
	"github.com/stretchr/testify/require"
)

func TestInsuranceFundModuleAddress(t *testing.T) {
	require.Equal(t, "dydx1c7ptc87hkd54e3r7zjy92q29xkq7t79w64slrq", types.InsuranceFundModuleAddress.String())
}
