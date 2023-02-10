package companies

import (
	"customer/internal/interfaces"
	"customer/internal/models"
	"fmt"
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

	fmt.Println(err, data, ctx)

	// if accountNumber, err = strconv.ParseInt(c.Param("account_number"), 10, 64); err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	// }

	if data, err = ctl.service.GetCompanies(ctx); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, models.GenericRes{
		Code:    http.StatusOK,
		Message: "SUCCEESS",
		Data:    data,
	})
}
