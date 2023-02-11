package companies

import (
	"context"
	"customer/internal/interfaces"
	"customer/internal/models"
)

type CompanyServiceImpl struct {
	companyRepository interfaces.CompanyRepository
}

func NewCompanyService(companyRepository interfaces.CompanyRepository) interfaces.CompanyService {
	return &CompanyServiceImpl{
		companyRepository: companyRepository,
	}
}

func (s *CompanyServiceImpl) GetCompanies(ctx context.Context) (data models.Company, err error) {
	// var (
	// 	companies []models.Company
	// )

	if data, err = s.companyRepository.GetCompanies(ctx); err != nil {
		return
	}

	// if len(companies) == 0 {
	// 	err = echo.NewHTTPError(http.StatusNotFound, "ACCOUNT_NOT_FOUND")
	// 	return
	// }

	return
}

func (s *CompanyServiceImpl) PostCompanies(ctx context.Context, body *models.Company) (err error) {
	return s.PostCompanies(ctx, &models.Company{
		Name:        body.Name,
		Description: body.Description,
	})
}
