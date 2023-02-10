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
	data, err = s.companyRepository.GetCompanies(ctx)

	if data, err = s.GetCompanies(ctx); err != nil {
		return
	}

	return s.companyRepository.GetCompanies(ctx)
}
