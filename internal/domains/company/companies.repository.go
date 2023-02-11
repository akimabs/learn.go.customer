package companies

import (
	"context"
	"customer/internal/interfaces"
	"customer/internal/lib"
	"customer/internal/models"
	"fmt"
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
		Find(&data).
		Error

	return
}

func (r *companyRepositoryImpl) PostCompanies(ctx context.Context, bodies *models.Company) (err error) {
	fmt.Println(bodies)
	// err = r.DB.Table("companies").
	// 	Select("name", "description").
	// 	Create(&bodies).
	// 	Error

	return
}
