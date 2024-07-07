package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-product-app/common/app"
	"go-product-app/common/postgresql"
	"go-product-app/controller"
	"go-product-app/persistence"
	"go-product-app/service"
)

func main() {
	ctx := context.Background()
	e := echo.New()

	configurationManager := app.NewConfigurationManager()

	//db - repo - service - controller
	dbPool := postgresql.GetConnectionPool(ctx, configurationManager.PostgreSqlConfig)
	productRepository := persistence.NewProductRepository(dbPool)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	productController.RegisterRoutes(e)

	err := e.Start("localhost:8080")
	if err != nil {
		panic(err)
	}
}
