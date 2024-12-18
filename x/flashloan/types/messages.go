package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgRequestLoan struct {
	Borrower   sdk.AccAddress `json:"borrower"`
	Amount     sdk.Coin       `json:"amount"`
	Callback   string         `json:"callback"`
}

func (msg MsgRequestLoan) Route() string { return RouterKey }
func (msg MsgRequestLoan) Type() string  { return "RequestLoan" }
func (msg MsgRequestLoan) ValidateBasic() error {
	if msg.Amount.IsZero() || msg.Amount.IsNegative() {
		return fmt.Errorf("amount must be positive")
	}
	if len(msg.Callback) == 0 {
		return fmt.Errorf("callback function required")
	}
	return nil
}

func (msg MsgRequestLoan) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Borrower}
}

