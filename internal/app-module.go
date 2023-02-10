package main

import (
	"context"
	"customer/internal/domains/companies"
	"customer/internal/interfaces"
	"customer/internal/lib"
	"customer/internal/middlewares"
	"fmt"
	"os"
	"time"

	"github.com/go-playground/validator"
	"go.uber.org/fx"
)

var Module = fx.Options(
	companies.Module,
	lib.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	echoHandler lib.EchoHandler,
	database lib.Database,
	company interfaces.CompanyRoute,
) {
	conn, _ := database.DB.DB()
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				conn.SetMaxIdleConns(10)
				conn.SetMaxOpenConns(100)
				conn.SetConnMaxLifetime(time.Hour)

				go func() {
					port, found := os.LookupEnv("PORT")

					if !found {
						port = "1323"
					}

					echoHandler.Echo.Validator = middlewares.CustomHandler(validator.New())
					echoHandler.Echo.HTTPErrorHandler = middlewares.ErrorHandler

					company.Setup()

					echoHandler.Echo.Logger.Fatal(
						echoHandler.Echo.Start(fmt.Sprintf(":%s", port)),
					)
				}()
				return nil
			},
			OnStop: func(context.Context) error {
				return conn.Close()
			},
		},
	)
}
