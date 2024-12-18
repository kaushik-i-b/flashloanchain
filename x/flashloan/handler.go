package flashloan

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/kaushik-i-b/x/flashloan/keeper"
	"github.com/kaushik-i-b/x/flashloan/types"
)

func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case *types.MsgRequestLoan:
			err := k.RequestLoan(ctx, *msg)
			if err != nil {
				return nil, err
			}
			return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
		default:
			return nil, sdk.ErrUnknownRequest("unknown flashloan message type")
		}
	}
}

