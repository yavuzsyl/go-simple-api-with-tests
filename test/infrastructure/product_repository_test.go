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

func TestAdd(t *testing.T) {
	expectedProducts := []domain.Product{
		{Id: 1, Name: "laptop", Price: 50000.0, Discount: 10.0, Store: "ABC TECH"},
	}

	product := domain.Product{
		Name:     "laptop",
		Price:    50000.0,
		Discount: 10.0,
		Store:    "ABC TECH",
	}

	t.Run("Add Product", func(t *testing.T) {
		productRepository.Add(product)
		actualProducts, _ := productRepository.GetAll()
		assert.Equal(t, 1, len(actualProducts))
		assert.Equal(t, expectedProducts, actualProducts)
	})

	clearSetup(ctx, dbPool)
}

func TestGetById(t *testing.T) {
	setup(ctx, dbPool)

	expectedProduct := domain.Product{
		Id: 1, Name: "air", Price: 3000.0, Discount: 22.0, Store: "ABC TECH",
	}

	t.Run("GetById", func(t *testing.T) {
		actualProduct, _ := productRepository.GetById(1)
		assert.Equal(t, expectedProduct, actualProduct)
	})

	clearSetup(ctx, dbPool)
}

func TestGeyByIdNotFound(t *testing.T) {
	setup(ctx, dbPool)

	t.Run("GetByIdNotFound", func(t *testing.T) {
		_, err := productRepository.GetById(100)
		assert.NotNil(t, err)
		assert.Equal(t, fmt.Sprintf("Product with id %d not found", 100), err.Error())
	})

	clearSetup(ctx, dbPool)
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

func TestGetAllByStore(t *testing.T) {
	setup(ctx, dbPool)

	expectedProducts := []domain.Product{
		{Id: 1, Name: "air", Price: 3000.0, Discount: 22.0, Store: "ABC TECH"},
		{Id: 2, Name: "iron", Price: 1500.0, Discount: 10.0, Store: "ABC TECH"},
		{Id: 3, Name: "fax", Price: 10000.0, Discount: 15.0, Store: "ABC TECH"},
	}

	t.Run("GetAllByStore", func(t *testing.T) {
		actualProducts, _ := productRepository.GetAllByStore("ABC TECH")
		assert.Equal(t, expectedProducts, actualProducts)
	})

	clearSetup(ctx, dbPool)
}

func TestDeleteById(t *testing.T) {
	setup(ctx, dbPool)

	t.Run("DeleteById", func(t *testing.T) {
		productRepository.DeleteById(1)
		products, _ := productRepository.GetAll()
		assert.Equal(t, 3, len(products))
	})

	clearSetup(ctx, dbPool)
}

func TestDeleteByIdNotFound(t *testing.T) {
	setup(ctx, dbPool)

	t.Run("DeleteByIdNotFound", func(t *testing.T) {
		err := productRepository.DeleteById(100)
		assert.NotNil(t, err)
		assert.Equal(t, fmt.Sprintf("Product with id %d not found", 100), err.Error())
	})

	clearSetup(ctx, dbPool)
}

func TestUpdateProductPrice(t *testing.T) {
	setup(ctx, dbPool)

	t.Run("UpdateProductPrice", func(t *testing.T) {
		product, _ := productRepository.GetById(1)
		productRepository.UpdateProductPrice(product.Id, 4000.0)
		updatedProduct, _ := productRepository.GetById(product.Id)
		assert.Equal(t, float32(4000.0), updatedProduct.Price)
	})

	clearSetup(ctx, dbPool)
}

func TestUpdateProductPriceNotFound(t *testing.T) {
	setup(ctx, dbPool)

	t.Run("UpdateProductPriceNotFound", func(t *testing.T) {
		err := productRepository.UpdateProductPrice(100, 4000.0)
		assert.NotNil(t, err)
		assert.Equal(t, fmt.Sprintf("Product with id %d not found", 100), err.Error())
	})

	clearSetup(ctx, dbPool)
}
