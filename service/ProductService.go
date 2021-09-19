package service

import (
	"API/data"
	customError "API/error"
	"API/persistence/repository"
)

type ProductService interface {
	Init(productRepository *repository.ProductRepositoryImpl)
	SaveProduct(product data.ProductDTO) (*data.ProductDTO, *customError.ServiceError)
	GetProduct(barcode string) (*data.ProductDTO, *customError.ServiceError)
	GetAllProducts() ([]*data.ProductDTO, *customError.ServiceError)
	UpdateProduct(product data.ProductDTO) (*data.ProductDTO, *customError.ServiceError)
	DeleteProduct(barcode string) *customError.ServiceError
}
