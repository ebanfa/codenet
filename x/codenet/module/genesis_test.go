package codenet_test

import (
	"testing"

	keepertest "codenet/testutil/keeper"
	"codenet/testutil/nullify"
	codenet "codenet/x/codenet/module"
	"codenet/x/codenet/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CodenetKeeper(t)
	codenet.InitGenesis(ctx, k, genesisState)
	got := codenet.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
