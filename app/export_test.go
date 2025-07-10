package app_test

import (
	"github.com/nftdance/dydxprotocol/testutil/app"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExportAppStateAndValidators_Panics(t *testing.T) {
	dydxApp := app.DefaultTestApp(nil)
	require.Panics(t, func() { dydxApp.ExportAppStateAndValidators(false, nil, nil) }) // nolint:errcheck
}
