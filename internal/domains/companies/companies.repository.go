package companies

import (
	"context"
	"customer/internal/interfaces"
	"customer/internal/lib"
	"customer/internal/models"
)

type companyRepositoryImpl struct {
	interfaces.RepositoryImpl
}

func NewCompanyRepository(db lib.Database) interfaces.CompanyRepository {
	companyRepository := &companyRepositoryImpl{}
	companyRepository.Database = db

	return companyRepository
}

func (r *companyRepositoryImpl) GetCompanies(ctx context.Context) (data models.Company, err error) {
	err = r.DB.Table("companies").
		Select("name").
		WithContext(ctx).
		Find(&data).
		Error
	return
}
