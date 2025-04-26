package cosmwasm_test

import (
	"testing"

	keepertest "github.com/olimdzhon/achilles/testutil/keeper"
	"github.com/olimdzhon/achilles/testutil/nullify"
	cosmwasm "github.com/olimdzhon/achilles/x/cosmwasm/module"
	"github.com/olimdzhon/achilles/x/cosmwasm/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CosmwasmKeeper(t)
	cosmwasm.InitGenesis(ctx, k, genesisState)
	got := cosmwasm.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
