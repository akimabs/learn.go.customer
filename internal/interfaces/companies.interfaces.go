package interfaces

import (
	"context"
	"customer/internal/models"

	"github.com/labstack/echo/v4"
)

type CompaniesRest interface {
	GetCompanies(echo.Context) error
	PostCompanies(echo.Context) error
}

type CompanyService interface {
	GetCompanies(ctx context.Context) (data models.Company, err error)
	PostCompanies(ctx context.Context, bodies *models.Company) (err error)
}

type CompanyRoute interface {
	Setup()
}

type CompanyRepository interface {
	Repository
	GetCompanies(ctx context.Context) (data models.Company, err error)
	PostCompanies(ctx context.Context, bodies *models.Company) (err error)
}

type PostCompanyDTO struct {
	Name        string
	Description string
}
