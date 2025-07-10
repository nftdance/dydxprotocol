package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/telemetry"
	"github.com/nftdance/dydxprotocol/daemons/liquidation/api"
	"github.com/nftdance/dydxprotocol/daemons/server/types"
	liquidationtypes "github.com/nftdance/dydxprotocol/daemons/server/types/liquidations"
	"github.com/nftdance/dydxprotocol/lib/metrics"
)

// LiquidationServer defines the fields required for liquidation updates.
type LiquidationServer struct {
	daemonLiquidationInfo *liquidationtypes.DaemonLiquidationInfo
}

// WithDaemonLiquidationInfo sets the `daemonLiquidationInfo` field.
// This is updated by the liquidation service with a list of potentially liquidatable
// subaccount ids to be processed by the `PerpetualLiquidationsKeeper`.
func (server *Server) WithDaemonLiquidationInfo(
	daemonLiquidationInfo *liquidationtypes.DaemonLiquidationInfo,
) *Server {
	server.daemonLiquidationInfo = daemonLiquidationInfo
	return server
}

// LiquidateSubaccounts stores the list of potentially liquidatable subaccount ids
// in a go-routine safe slice.
func (s *Server) LiquidateSubaccounts(
	ctx context.Context,
	req *api.LiquidateSubaccountsRequest,
) (
	response *api.LiquidateSubaccountsResponse,
	err error,
) {
	telemetry.ModuleSetGauge(
		metrics.LiquidationDaemon,
		float32(len(req.LiquidatableSubaccountIds)),
		metrics.LiquidatableSubaccountIds,
		metrics.Received,
		metrics.Count,
	)

	s.daemonLiquidationInfo.Update(
		req.BlockHeight,
		req.LiquidatableSubaccountIds,
		req.NegativeTncSubaccountIds,
		req.SubaccountOpenPositionInfo,
	)

	// Capture valid responses in metrics.
	s.reportValidResponse(types.LiquidationsDaemonServiceName)

	return &api.LiquidateSubaccountsResponse{}, nil
}
