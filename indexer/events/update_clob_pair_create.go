package events

import (
	v1 "github.com/nftdance/dydxprotocol/indexer/protocol/v1"
	"github.com/nftdance/dydxprotocol/x/clob/types"
	satypes "github.com/nftdance/dydxprotocol/x/subaccounts/types"
)

// NewUpdateClobPairEvent creates a UpdateClobPairEventV1
// representing an update of a clob pair.
func NewUpdateClobPairEvent(
	clobPairId types.ClobPairId,
	status types.ClobPair_Status,
	quantumConversionExponent int32,
	subticksPerTick types.SubticksPerTick,
	stepBaseQuantums satypes.BaseQuantums,
) *UpdateClobPairEventV1 {
	// ClobPair metadata is not included in the event because it should never change.
	// A change would imply either transitioning to a different perpetual market or transitioning
	// to an asset market and asset markets are not supported.
	return &UpdateClobPairEventV1{
		ClobPairId:                uint32(clobPairId),
		Status:                    v1.ConvertToClobPairStatus(status),
		QuantumConversionExponent: quantumConversionExponent,
		SubticksPerTick:           uint32(subticksPerTick),
		StepBaseQuantums:          uint64(stepBaseQuantums),
	}
}
