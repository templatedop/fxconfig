package fxconfig

import (
	"go.uber.org/fx"
)

var FxConfigModule = fx.Module("config",
	fx.Provide(
		NewDefaultConfigFactory,
		NewFxConfig,
	),
)

type FxConfigParam struct {
	fx.In
	Factory ConfigFactory
}

func NewFxConfig(p FxConfigParam) (Econfig, error) {
	c, err := p.Factory.Create(
		WithFileName("config"),
		WithType("yaml"),
		WithPath("./configs"),
	)
	
	return c, err
}
