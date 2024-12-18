package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/kaushik-i-b/x/flashloan/types"
)

func (k Keeper) RequestLoan(ctx sdk.Context, msg types.MsgRequestLoan) error {
	// Ensure the module has sufficient liquidity
	moduleAcc := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	if !k.bankKeeper.HasBalance(ctx, moduleAcc, msg.Amount) {
		return fmt.Errorf("insufficient liquidity in module account")
	}

	// Transfer loan amount to borrower
	err := k.bankKeeper.SendCoins(ctx, moduleAcc, msg.Borrower, sdk.Coins{msg.Amount})
	if err != nil {
		return err
	}

	// Calculate the fee: Fee = Amount * FeeRate
	fee := msg.Amount.Amount.ToDec().Mul(types.DefaultParams().FeeRate)
	feeAmount := sdk.NewCoin(msg.Amount.Denom, fee.TruncateInt())

	// The expected repayment amount (Loan + Fee)
	expectedRepayment := msg.Amount.Add(feeAmount)

	// Execute the callback (borrower logic)
	callbackErr := ctx.EventManager().EmitTypedEvent(sdk.NewEvent(
		"flashloan",
		sdk.NewAttribute("borrower", msg.Borrower.String()),
		sdk.NewAttribute("loan_amount", msg.Amount.String()),
		sdk.NewAttribute("fee", feeAmount.String()),
	))
	if callbackErr != nil {
		return callbackErr
	}

	// Ensure borrower has enough to repay loan + fee
	if !k.bankKeeper.HasBalance(ctx, msg.Borrower, sdk.Coins{expectedRepayment}) {
		return fmt.Errorf("repayment failed; borrower does not have enough funds")
	}

	// Repay the loan (borrower returns loan + fee)
	err = k.bankKeeper.SendCoins(ctx, msg.Borrower, moduleAcc, sdk.Coins{expectedRepayment})
	if err != nil {
		return err
	}

	// Transaction successfully completed
	return nil
}

