package keeper

import (
	"github.com/kaushik-i-b/flashloanchain/x/flashloan/types"
)

var _ types.QueryServer = Keeper{}
