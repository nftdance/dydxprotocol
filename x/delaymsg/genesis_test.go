package delaymsg_test

import (
	"testing"

	testutildelaymsg "github.com/nftdance/dydxprotocol/testutil/delaymsg"
	"github.com/nftdance/dydxprotocol/testutil/keeper"
	"github.com/nftdance/dydxprotocol/x/delaymsg"
	"github.com/nftdance/dydxprotocol/x/delaymsg/types"
	"github.com/stretchr/testify/require"
)

func TestInitGenesis(t *testing.T) {
	tests := map[string]struct {
		genesisState *types.GenesisState
	}{
		"default genesis": {
			genesisState: types.DefaultGenesis(),
		},
		"non-default genesis (e.g. network restart)": {
			genesisState: &types.GenesisState{
				DelayedMessages: []*types.DelayedMessage{
					{
						Id:          3,
						Msg:         testutildelaymsg.CreateTestAnyMsg(t),
						BlockHeight: 10,
					},
					{
						Id:          7,
						Msg:         testutildelaymsg.CreateTestAnyMsg(t),
						BlockHeight: 15,
					},
					{
						Id:          11,
						Msg:         testutildelaymsg.CreateTestAnyMsg(t),
						BlockHeight: 10,
					},
				},
				NextDelayedMessageId: 20,
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			ctx, delaymsgKeeper, _, _, _, _ := keeper.DelayMsgKeepers(t)
			delaymsgKeeper.InitializeForGenesis(ctx)
			delaymsg.InitGenesis(ctx, *delaymsgKeeper, *tc.genesisState)
			got := delaymsg.ExportGenesis(ctx, *delaymsgKeeper)
			require.NotNil(t, got)
			require.Equal(t, tc.genesisState, got)
		})
	}
}
