package companies

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewCompanyService),
	fx.Provide(NewCompanyRest),
	fx.Provide(NewCompanyRoute),
	fx.Provide(NewCompanyRepository),
)
