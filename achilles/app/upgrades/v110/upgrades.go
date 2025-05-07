package v110

import (
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/olimdzhon/achilles/app/upgrades"
)

const UpgradeName = "v1.1.0-rc1"

func CreateUpgradeHandler(
	_ module.Configurator,
	_ *module.SimulationManager,
) upgradetypes.UpgradeHandler {
	return upgrades.EmptyUpgradeHandler()
}