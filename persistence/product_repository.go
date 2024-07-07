package persistence

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
	"go-product-app/domain"
	"go-product-app/persistence/errorMessages"
)

type IProductRepository interface {
	GetAll() ([]domain.Product, error)
	GetAllByStore(store string) ([]domain.Product, error)
	Add(product domain.Product) error
	GetById(id int64) (domain.Product, error)
	DeleteById(id int64) error
	UpdateProductPrice(id int64, price float32) error
}

type ProductRepository struct {
	dbPool *pgxpool.Pool
}

func NewProductRepository(dbPool *pgxpool.Pool) IProductRepository {
	return &ProductRepository{dbPool: dbPool}
}

func (productRepository *ProductRepository) Add(product domain.Product) error {
	ctx := context.Background()

	sqlCommand := `INSERT INTO products(name, price, discount, store) VALUES($1, $2, $3, $4)`

	exec, err := productRepository.dbPool.Exec(ctx, sqlCommand, product.Name, product.Price, product.Discount, product.Store)
	if err != nil {
		log.Error("Error while inserting product: %v\n", err)
		return err
	}

	log.Info("Product added successfully: %v\n", exec)
	return nil
}

func (productRepository *ProductRepository) GetById(id int64) (domain.Product, error) {
	ctx := context.Background()

	sqlCommand := `SELECT * FROM products WHERE id = $1`

	var product domain.Product
	err := productRepository.dbPool.QueryRow(ctx, sqlCommand, id).Scan(&product.Id, &product.Name, &product.Price, &product.Discount, &product.Store)
	if err != nil && err.Error() == errorMessages.NOT_FOUND {
		return domain.Product{}, errors.New(fmt.Sprintf("Product with id %d not found", id))
	}

	if err != nil {
		log.Error("Error while fetching product with id: %d %v\n", id, err)
		return domain.Product{}, errors.New(fmt.Sprintf("Error while fetching product by id %d", id))
	}

	return product, nil
}

func (productRepository *ProductRepository) GetAll() ([]domain.Product, error) {
	ctx := context.Background()
	productRows, err := productRepository.dbPool.Query(ctx, "SELECT * FROM products")
	if err != nil {
		log.Error("Error while fetching products: %v\n", err)
		return []domain.Product{}, err
	}

	return extractProductsFromRows(productRows, err)
}

func (productRepository *ProductRepository) GetAllByStore(store string) ([]domain.Product, error) {
	context := context.Background()

	query := `SELECT *FROM products WHERE store= $1`

	rows, err := productRepository.dbPool.Query(context, query, store)
	if err != nil {
		log.Error("Error while fetching products: %v\n", err)
		return []domain.Product{}, err
	}

	return extractProductsFromRows(rows, err)
}

func extractProductsFromRows(productRows pgx.Rows, err error) ([]domain.Product, error) {
	var products []domain.Product

	for productRows.Next() {

		var product domain.Product
		err = productRows.Scan(&product.Id, &product.Name, &product.Price, &product.Discount, &product.Store)
		if err != nil {
			log.Error("Error while scanning product rows: %v\n", err)
			return []domain.Product{}, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (productRepository *ProductRepository) DeleteById(id int64) error {
	ctx := context.Background()

	_, err := productRepository.GetById(id)
	if err != nil {
		log.Error("product with id %d not found: %v\n", id, err)
		return errors.New(fmt.Sprintf("Product with id %d not found", id))
	}

	sqlCommand := `DELETE FROM products WHERE id = $1`

	exec, err := productRepository.dbPool.Exec(ctx, sqlCommand, id)
	if err != nil {
		log.Error("Error while deleting product with id:%d %v\n", id, err)
		return errors.New(fmt.Sprintf("Error while deleting product with id %d", id))
	}

	log.Info("Product deleted successfully: %v\n", exec)
	return nil
}

func (productRepository *ProductRepository) UpdateProductPrice(id int64, price float32) error {
	ctx := context.Background()

	_, err := productRepository.GetById(id)
	if err != nil {
		log.Error("product with id %d not found: %v\n", id, err)
		return err
	}

	sqlCommand := `UPDATE products SET price = $1 WHERE id = $2`

	exec, err := productRepository.dbPool.Exec(ctx, sqlCommand, price, id)

	if err != nil {
		log.Error("Error while updating product price with id:%d %v\n", id, err)
		return errors.New(fmt.Sprintf("Error while updating product price with id %d", id))
	}

	log.Info("Product price updated successfully: %v\n", exec)
	return nil
}
