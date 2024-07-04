package infrastructure

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
	"go-product-app/common/postgresql"
	"go-product-app/domain"
	"go-product-app/persistence"
	"os"
	"testing"
)

//go test -v - all test will run under current directory

var productRepository persistence.IProductRepository
var dbPool *pgxpool.Pool
var ctx context.Context

func TestMain(m *testing.M) {
	ctx = context.Background()

	dbPool = postgresql.GetConnectionPool(ctx, postgresql.Config{
		Host:                  "localhost",
		Port:                  "6432",
		Database:              "productapp",
		User:                  "postgres",
		Password:              "postgres",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "30s",
	})

	productRepository = persistence.NewProductRepository(dbPool)
	fmt.Println("before all tests")
	exitCode := m.Run()
	fmt.Println("after all tests")
	os.Exit(exitCode)
}

func setup(ctx context.Context, dbPool *pgxpool.Pool) {
	TestDataInitialize(ctx, dbPool)
}

func clearSetup(ctx context.Context, dbPool *pgxpool.Pool) {
	TruncateTestData(ctx, dbPool)
}

func TestGetAll(t *testing.T) {
	setup(ctx, dbPool)

	expectedProducts := []domain.Product{
		{Id: 1, Name: "air", Price: 3000.0, Discount: 22.0, Store: "ABC TECH"},
		{Id: 2, Name: "iron", Price: 1500.0, Discount: 10.0, Store: "ABC TECH"},
		{Id: 3, Name: "fax", Price: 10000.0, Discount: 15.0, Store: "ABC TECH"},
		{Id: 4, Name: "phone", Price: 2000.0, Discount: 0.0, Store: "x brand"},
	}

	t.Run("GetAll", func(t *testing.T) {
		actualProducts, _ := productRepository.GetAll()
		assert.Equal(t, 4, len(actualProducts))
		assert.Equal(t, expectedProducts, actualProducts)

	})
	clearSetup(ctx, dbPool)

}
