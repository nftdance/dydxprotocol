package keeper_test

import (
	keepertest "github.com/nftdance/dydxprotocol/testutil/keeper"
	"github.com/nftdance/dydxprotocol/x/perpetuals/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestParams(t *testing.T) {
	pc := keepertest.PerpetualsKeepers(t)

	tests := map[string]struct {
		req         *types.QueryParamsRequest
		res         *types.QueryParamsResponse
		expectedErr error
	}{
		"nil request": {
			req:         nil,
			res:         nil,
			expectedErr: status.Error(codes.InvalidArgument, "invalid request"),
		},
		"valid request": {
			req: &types.QueryParamsRequest{},
			res: &types.QueryParamsResponse{
				Params: types.DefaultGenesis().Params,
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := pc.PerpetualsKeeper.Params(pc.Ctx, tc.req)
			if tc.expectedErr != nil {
				require.ErrorIs(t, err, tc.expectedErr)
			} else {
				require.Equal(t, tc.res, res)
			}
		})
	}
}
