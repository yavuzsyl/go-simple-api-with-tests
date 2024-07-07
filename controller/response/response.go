package response

import "go-product-app/domain"

type ErrorResponse struct {
	Description string `json:"description"`
}

type ProductResponse struct {
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Discount float32 `json:"discount"`
	Store    string  `json:"store"`
}

func ToProductResponse(product domain.Product) ProductResponse {
	return ProductResponse{
		Name:     product.Name,
		Price:    product.Price,
		Discount: product.Discount,
		Store:    product.Store,
	}
}

func ToProductResponseList(products []domain.Product) []ProductResponse {
	productResponseList := make([]ProductResponse, 0)
	for _, product := range products {
		productResponseList = append(productResponseList, ToProductResponse(product))
	}
	return productResponseList
}
