package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName = "flashloan"
	StoreKey   = ModuleName
	RouterKey  = ModuleName
)

type FlashLoanParams struct {
	FeeRate sdk.Dec // Fee percentage for flash loans
}

func DefaultParams() FlashLoanParams {
	return FlashLoanParams{FeeRate: sdk.NewDecWithPrec(1, 2)} // 1% fee
}

