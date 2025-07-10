package types_test

import (
	"testing"

	"github.com/nftdance/dydxprotocol/x/sending/types"
	"github.com/stretchr/testify/require"
)

func TestModuleKeys(t *testing.T) {
	require.Equal(t, "sending", types.ModuleName)
	require.Equal(t, "sending", types.StoreKey)
}
