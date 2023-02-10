package companies

import (
	"context"
	"customer/internal/interfaces"
	"customer/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CompanyServiceImpl struct {
	companyRepository interfaces.CompanyRepository
}

func NewCompanyService(companyRepository interfaces.CompanyRepository) interfaces.CompanyService {
	return &CompanyServiceImpl{
		companyRepository: companyRepository,
	}
}

func (s *CompanyServiceImpl) GetCompanies(ctx context.Context) (data models.Companies, err error) {
	data, err = s.companyRepository.GetCompanies(ctx)

	if data.Name == "" {
		err = echo.NewHTTPError(http.StatusNotFound, "Asdasd")
		return
	}
	return
}
