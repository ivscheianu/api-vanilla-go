package repository

import (
	"API/data"
	customError "API/error"
	"API/persistence/dao"
)

type ProductRepository interface {
	Init(productDAO *dao.ProductDAOImpl)
	SaveProduct(product *data.ProductDO) (data.ProductDO, *customError.PersistenceError)
	GetProduct(barcode string) (data.ProductDO, *customError.PersistenceError)
	GetAll() ([]*data.ProductDO, *customError.PersistenceError)
	UpdateProduct(product *data.ProductDO) (*data.ProductDO, *customError.PersistenceError)
	DeleteProduct(barcode string) *customError.PersistenceError
}
