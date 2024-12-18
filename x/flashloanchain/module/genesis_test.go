package flashloanchain_test

import (
	"testing"

	keepertest "github.com/kaushik-i-b/flashloanchain/testutil/keeper"
	"github.com/kaushik-i-b/flashloanchain/testutil/nullify"
	flashloanchain "github.com/kaushik-i-b/flashloanchain/x/flashloanchain/module"
	"github.com/kaushik-i-b/flashloanchain/x/flashloanchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FlashloanchainKeeper(t)
	flashloanchain.InitGenesis(ctx, k, genesisState)
	got := flashloanchain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
