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
	// CompanyID = utils.GenerateId()
	return companyRepository
}

func (r *companyRepositoryImpl) GetCompanies(ctx context.Context) (data models.Company, err error) {
	err = r.DB.Table("companies").
		Find(&data).
		Error

	return
}

func (r *companyRepositoryImpl) PostCompanies(ctx context.Context, body *models.Company) (err error) {
	err = r.DB.Table("companies").
		Create(&body).
		Error

	return
}
