package service

import (
	"errors"
	"go-product-app/domain"
	"go-product-app/persistence"
	"go-product-app/service/model"
)

type IProductService interface {
	Add(product model.CreateProduct) error
	UpdatePrice(id int64, price float32) error
	DeleteById(id int64) error
	GetById(id int64) (domain.Product, error)
	GetAll() ([]domain.Product, error)
	GetAllByStore(store string) ([]domain.Product, error)
}

type ProductService struct {
	productRepository persistence.IProductRepository
}

func NewProductService(productRepository persistence.IProductRepository) IProductService {
	return &ProductService{productRepository: productRepository}
}

func (productService *ProductService) Add(product model.CreateProduct) error {

	validationErr := validateProduct(product)
	if validationErr != nil {
		return validationErr
	}

	productEntity := domain.Product{Name: product.Name, Price: product.Price, Discount: product.Discount, Store: product.Store}
	err := productService.productRepository.Add(productEntity)
	if err != nil {
		return err
	}

	return nil
}

func (productService *ProductService) UpdatePrice(id int64, price float32) error {
	return productService.productRepository.UpdateProductPrice(id, price)
}

func (productService *ProductService) DeleteById(id int64) error {
	return productService.productRepository.DeleteById(id)
}

func (productService *ProductService) GetById(id int64) (domain.Product, error) {
	return productService.productRepository.GetById(id)
}

func (productService *ProductService) GetAll() ([]domain.Product, error) {
	return productService.productRepository.GetAll()
}

func (productService *ProductService) GetAllByStore(store string) ([]domain.Product, error) {
	return productService.productRepository.GetAllByStore(store)
}

func validateProduct(product model.CreateProduct) error {
	if product.Discount > 70 || product.Discount < 0 {
		return errors.New("Discount should be between 0 and 70")
	}
	return nil
}
