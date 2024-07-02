package persistence

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
	"go-product-app/domain"
)

type IProductRepository interface {
	GetAll() ([]domain.Product, error)
}

type ProductRepository struct {
	dbPool *pgxpool.Pool
}

func NewProductRepository(dbPool *pgxpool.Pool) IProductRepository {
	return &ProductRepository{dbPool: dbPool}
}

func (productRepository *ProductRepository) GetAll() ([]domain.Product, error) {
	ctx := context.Background()
	productRows, err := productRepository.dbPool.Query(ctx, "SELECT * FROM products")
	if err != nil {
		log.Error("Error while fetching products: %v\n", err)
		return []domain.Product{}, err
	}

	var products []domain.Product

	for productRows.Next() {

		var product domain.Product
		err = productRows.Scan(&product.Id, &product.Name, &product.Price, &product.Discount, &product.Price)
		if err != nil {
			log.Error("Error while scanning product rows: %v\n", err)
			return []domain.Product{}, err
		}

		products = append(products, product)
	}

	return products, nil
}
