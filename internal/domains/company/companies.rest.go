package companies

import (
	"customer/internal/interfaces"
	"customer/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CompanyRestImpl struct {
	service interfaces.CompanyService
}

func NewCompanyRest(service interfaces.CompanyService) interfaces.CompaniesRest {
	return &CompanyRestImpl{
		service: service,
	}
}

func (ctl *CompanyRestImpl) GetCompanies(c echo.Context) error {

	var (
		err  error
		data models.Company
		ctx  = c.Request().Context()
	)

	if data, err = ctl.service.GetCompanies(ctx); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, models.GenericRes{
		Code:    http.StatusOK,
		Message: "SUCCESS",
		Data:    data,
	})
}

func (ctl *CompanyRestImpl) PostCompanies(c echo.Context) error {
	var (
		err error
		ctx = c.Request().Context()
	)

	bodies := new(models.Company)

	if err = c.Bind(bodies); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = ctl.service.PostCompanies(ctx, &models.Company{
		Name:        bodies.Name,
		Description: bodies.Description,
	}); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, models.GenericRes{
		Code:    http.StatusOK,
		Message: "SUCCESS",
		Data:    bodies,
	})
}
