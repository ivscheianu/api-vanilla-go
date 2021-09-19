package dao

import (
	"API/data"
	customError "API/error"
)

type ProductDAO interface {
	Init()
	SaveProducts(products []*data.ProductDO) (savedProducts []data.ProductDO, persistenceError *customError.PersistenceError)
	SaveProduct(product *data.ProductDO) (savedProduct data.ProductDO, persistenceError *customError.PersistenceError)
	GetAllProducts() (products []*data.ProductDO, persistenceError *customError.PersistenceError)
	UpdateProduct(updatedProductFromPayload *data.ProductDO) (updatedProduct *data.ProductDO, persistenceError *customError.PersistenceError)
	DeleteProduct(barcode string) (persistenceError *customError.PersistenceError)
}
