package controller

import (
	"github.com/labstack/echo/v4"
	"go-product-app/controller/request"
	"go-product-app/controller/response"
	"go-product-app/domain"
	"go-product-app/service"
	"net/http"
	"strconv"
)

type ProductController struct {
	productService service.IProductService
}

func NewProductController(productService service.IProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (productController *ProductController) RegisterRoutes(e *echo.Echo) {
	e.GET("/api/v1/products", productController.GetAll)
	e.GET("/api/v1/products/:id", productController.GetById)
	e.POST("/api/v1/products", productController.Add)
	e.PUT("/api/v1/products/:id", productController.UpdatePrice)
	e.DELETE("/api/v1/products/:id", productController.DeleteById)
}

func (productController *ProductController) GetAll(c echo.Context) error {
	store := c.QueryParam("store")
	var err error
	var products []domain.Product
	if len(store) > 0 {
		products, err = productController.productService.GetAllByStore(store)
	} else {
		products, err = productController.productService.GetAll()
	}

	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{
			Description: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.ToProductResponseList(products))
}

func (productController *ProductController) GetById(c echo.Context) error {
	idParam := c.Param("id")
	if len(idParam) == 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Description: "Id parameter is required",
		})
	}
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Description: err.Error(),
		})
	}
	product, err := productController.productService.GetById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{
			Description: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.ToProductResponse(product))
}

func (productController *ProductController) Add(c echo.Context) error {
	var addProductRequest request.AddProductRequest
	err := c.Bind(&addProductRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Description: err.Error(),
		})
	}

	product := addProductRequest.ToModel()
	err = productController.productService.Add(product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Description: err.Error(),
		})
	}
	return c.NoContent(http.StatusCreated)
}

func (productController *ProductController) UpdatePrice(c echo.Context) error {
	id := c.Param("id")
	price := c.QueryParam("price")
	if len(id) == 0 || len(price) == 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Description: "Id and price parameters are required",
		})
	}
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Description: err.Error(),
		})
	}
	priceFloat, err := strconv.ParseFloat(price, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Description: err.Error(),
		})
	}
	err = productController.productService.UpdatePrice(idInt, float32(priceFloat))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Description: err.Error(),
		})
	}
	return c.NoContent(http.StatusOK)
}

func (productController *ProductController) DeleteById(c echo.Context) error {
	id := c.Param("id")
	if len(id) == 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Description: "Id parameter is required",
		})
	}
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Description: err.Error(),
		})
	}
	err = productController.productService.DeleteById(idInt)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Description: err.Error(),
		})
	}
	return c.NoContent(http.StatusOK)

}
