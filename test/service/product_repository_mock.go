package service

import (
	"errors"
	"fmt"
	"go-product-app/domain"
	"go-product-app/persistence"
)

type ProductRepositoryMock struct {
	products []domain.Product
}

func NewProductRepositoryMock(initialProducts []domain.Product) persistence.IProductRepository {
	return &ProductRepositoryMock{products: initialProducts}
}

func (productRepository *ProductRepositoryMock) Add(product domain.Product) error {
	product.Id = int64(len(productRepository.products) + 1)
	productRepository.products = append(productRepository.products, product)
	return nil
}

func (productRepository *ProductRepositoryMock) GetById(id int64) (domain.Product, error) {
	for _, product := range productRepository.products {
		if product.Id == id {
			return product, nil
		}
	}

	return domain.Product{}, errors.New(fmt.Sprintf("Product with id %d not found", id))
}

func (productRepository *ProductRepositoryMock) GetAll() ([]domain.Product, error) {
	return productRepository.products, nil
}

func (productRepository *ProductRepositoryMock) GetAllByStore(store string) ([]domain.Product, error) {
	var products []domain.Product
	for _, product := range productRepository.products {
		if product.Store == store {
			products = append(products, product)
		}
	}

	return products, nil
}

func (productRepository *ProductRepositoryMock) DeleteById(id int64) error {
	for i, product := range productRepository.products {
		if product.Id == id {
			productRepository.products = append(productRepository.products[:i], productRepository.products[i+1:]...)
			return nil
		}
	}

	return errors.New(fmt.Sprintf("Product with id %d not found", id))
}

func (productRepository *ProductRepositoryMock) UpdateProductPrice(id int64, price float32) error {
	for i, product := range productRepository.products {
		if product.Id == id {
			productRepository.products[i].Price = price
			return nil
		}
	}

	return errors.New(fmt.Sprintf("Product with id %d not found", id))
}
