package flashloan

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/module"
)

type AppModule struct{}

func NewAppModule() AppModule {
	return AppModule{}
}

func (am AppModule) RegisterServices(cfg module.Configurator) {}

func (am AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
	return nil
}

func (am AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
	return nil
}

func (am AppModule) Route() sdk.Route {
	return sdk.NewRoute("flashloan", NewHandler())
}

