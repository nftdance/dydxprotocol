package keeper_test

import (
	"testing"

	"github.com/nftdance/dydxprotocol/lib"
	testapp "github.com/nftdance/dydxprotocol/testutil/app"
	"github.com/nftdance/dydxprotocol/testutil/constants"
	"github.com/nftdance/dydxprotocol/x/revshare/keeper"
	"github.com/nftdance/dydxprotocol/x/revshare/types"
	"github.com/stretchr/testify/require"
)

func TestUpdateUnconditionalRevShareConfig(t *testing.T) {
	tests := map[string]struct {
		msg         *types.MsgUpdateUnconditionalRevShareConfig
		expectedErr error
	}{
		"Success": {
			msg: &types.MsgUpdateUnconditionalRevShareConfig{
				Authority: lib.GovModuleAddress.String(),
				Config: types.UnconditionalRevShareConfig{
					Configs: []types.UnconditionalRevShareConfig_RecipientConfig{
						{
							Address:  constants.AliceAccAddress.String(),
							SharePpm: 100_000, // 10%
						},
						{
							Address:  constants.BobAccAddress.String(),
							SharePpm: 100_000, // 10%
						},
					},
				},
			},
		},
		"Failure when sum of shares > 100%": {
			msg: &types.MsgUpdateUnconditionalRevShareConfig{
				Authority: lib.GovModuleAddress.String(),
				Config: types.UnconditionalRevShareConfig{
					Configs: []types.UnconditionalRevShareConfig_RecipientConfig{
						{
							Address:  constants.AliceAccAddress.String(),
							SharePpm: 100_000, // 10%
						},
						{
							Address:  constants.BobAccAddress.String(),
							SharePpm: 910_000, // 91%
						},
					},
				},
			},
			expectedErr: types.ErrInvalidRevShareConfig,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tApp := testapp.NewTestAppBuilder(t).Build()
			ctx := tApp.InitChain()
			k := tApp.App.RevShareKeeper
			msgServer := keeper.NewMsgServerImpl(k)

			_, err := msgServer.UpdateUnconditionalRevShareConfig(ctx, tc.msg)

			if tc.expectedErr == nil {
				require.NoError(t, err)
				storedConfig, err := k.GetUnconditionalRevShareConfigParams(ctx)
				require.NoError(t, err)
				require.Equal(t, tc.msg.Config, storedConfig)
			} else {
				require.ErrorIs(t, err, tc.expectedErr)
			}
		})
	}
}
