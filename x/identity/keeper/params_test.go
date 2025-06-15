package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "Nexelra/testutil/keeper"
	"Nexelra/x/identity/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.IdentityKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
