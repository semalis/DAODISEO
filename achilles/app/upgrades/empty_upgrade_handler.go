package upgrades

import (
    "context"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

// EmptyUpgradeHandler - should be used to create upgrade handler for gov software upgrades that require no migration
func EmptyUpgradeHandler() upgradetypes.UpgradeHandler {
    return func(ctx context.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
        // your logic here
        return vm, nil
    }
}
