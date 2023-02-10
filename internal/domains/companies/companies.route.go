package companies

import (
	"customer/internal/interfaces"
	"customer/internal/lib"
)

type CompaniesRouteImpl struct {
	echo lib.EchoHandler
	rest interfaces.CompaniesRest
}

func NewCompanyRoute(echo lib.EchoHandler, rest interfaces.CompaniesRest) interfaces.CompanyRoute {
	return &CompaniesRouteImpl{
		echo: echo,
		rest: rest,
	}
}

func (r *CompaniesRouteImpl) Setup() {
	companies := r.echo.Echo.Group("/companies")
	{
		companies.GET("/", r.rest.GetCompanies)
	}
}
