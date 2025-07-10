package v_5_0_0

import (
	store "cosmossdk.io/store/types"
	"github.com/nftdance/dydxprotocol/app/upgrades"

	vaulttypes "github.com/nftdance/dydxprotocol/x/vault/types"
)

const (
	UpgradeName = "v5.0.0"
	// VEEnableHeightDelta is the number of blocks after the v5.0.0 upgrade to turn Vote Extensions on.
	VEEnableHeightDelta = int64(4)
)

var Upgrade = upgrades.Upgrade{
	UpgradeName: UpgradeName,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{
			vaulttypes.StoreKey,
		},
	},
}
