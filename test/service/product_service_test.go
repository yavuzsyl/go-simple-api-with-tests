package service

import (
	"github.com/stretchr/testify/assert"
	"go-product-app/domain"
	"go-product-app/service"
	"go-product-app/service/model"
	"os"
	"testing"
)

var productService service.IProductService

func TestMain(m *testing.M) {
	initialProducts := []domain.Product{
		{Id: 1, Name: "air", Price: 3000.0, Discount: 22.0, Store: "ABC TECH"},
		{Id: 2, Name: "iron", Price: 1500.0, Discount: 10.0, Store: "ABC TECH"},
		{Id: 3, Name: "fax", Price: 10000.0, Discount: 15.0, Store: "ABC TECH"},
		{Id: 4, Name: "phone", Price: 2000.0, Discount: 0.0, Store: "x brand"},
	}
	productRepositoryMock := NewProductRepositoryMock(initialProducts)
	productService = service.NewProductService(productRepositoryMock)

	exitCode := m.Run()
	os.Exit(exitCode)
}

func Test_GetAll_ShouldReturnAllProducts(t *testing.T) {
	t.Run("GetAll", func(t *testing.T) {
		products, _ := productService.GetAll()
		assert.Equal(t, 4, len(products))
	})
}

func Test_Add_ShouldAddProduct_WhenProductIsValid(t *testing.T) {
	t.Run("Add", func(t *testing.T) {
		product := model.CreateProduct{Name: "tv", Price: 5000.0, Discount: 10.0, Store: "ABC TECH"}
		err := productService.Add(product)
		allProducts, _ := productService.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, 5, len(allProducts))
		assert.Equal(t, domain.Product{
			Id: 5, Name: "tv", Price: 5000.0, Discount: 10.0, Store: "ABC TECH",
		}, allProducts[4])
	})
}

func Test_Add_ShouldReturnError_WhenDiscountIsInvalid(t *testing.T) {
	t.Run("Add", func(t *testing.T) {
		product := model.CreateProduct{Name: "tv", Price: 5000.0, Discount: 80.0, Store: "AVV"}
		err := productService.Add(product)
		assert.NotNil(t, err)
		assert.Equal(t, "Discount should be between 0 and 70", err.Error())
	})
}

func Test_UpdatePrice_ShouldUpdatePrice_WhenProductExists(t *testing.T) {
	t.Run("UpdatePrice", func(t *testing.T) {
		err := productService.UpdatePrice(1, 4000.0)
		product, _ := productService.GetById(1)
		assert.Equal(t, float32(4000.0), product.Price)
		assert.Nil(t, err)
	})
}

func Test_UpdatePrice_ShouldReturnError_WhenProductDoesNotExist(t *testing.T) {
	t.Run("UpdatePrice", func(t *testing.T) {
		err := productService.UpdatePrice(100, 4000.0)
		assert.NotNil(t, err)
		assert.Equal(t, "Product with id 100 not found", err.Error())
	})
}

func Test_GetById_ShouldReturnProduct_WhenProductExists(t *testing.T) {
	t.Run("GetById", func(t *testing.T) {
		product, _ := productService.GetById(1)
		assert.NotNil(t, product)
	})
}

func Test_GetAllByStore(t *testing.T) {
	t.Run("GetAllByStore", func(t *testing.T) {
		products, _ := productService.GetAllByStore("ABC TECH")
		assert.Equal(t, 3, len(products))
	})
}

func Test_DeleteById_ShouldDeleteProduct_WhenProductExists(t *testing.T) {
	t.Run("DeleteById", func(t *testing.T) {
		err := productService.DeleteById(1)
		products, _ := productService.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, 3, len(products))
	})
}
