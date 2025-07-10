package v_8_1

import (
	store "cosmossdk.io/store/types"
	"github.com/nftdance/dydxprotocol/app/upgrades"
)

const (
	UpgradeName = "v8.1"
)

var Upgrade = upgrades.Upgrade{
	UpgradeName:   UpgradeName,
	StoreUpgrades: store.StoreUpgrades{},
}
