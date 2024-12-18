package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/kaushik-i-b/flashloanchain/testutil/keeper"
	"github.com/kaushik-i-b/flashloanchain/x/flashloanchain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.FlashloanchainKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
