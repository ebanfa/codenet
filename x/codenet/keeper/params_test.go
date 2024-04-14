package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "codenet/testutil/keeper"
	"codenet/x/codenet/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.CodenetKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
