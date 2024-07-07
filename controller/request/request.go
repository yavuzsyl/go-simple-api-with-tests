package request

import "go-product-app/service/model"

type AddProductRequest struct {
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Discount float32 `json:"discount"`
	Store    string  `json:"store"`
}

func (addProductRequest AddProductRequest) ToModel() model.CreateProduct {
	return model.CreateProduct{
		Name:     addProductRequest.Name,
		Price:    addProductRequest.Price,
		Discount: addProductRequest.Discount,
		Store:    addProductRequest.Store,
	}
}
