package repository

import (
	"API/data"
	customError "API/error"
	"API/persistence/dao"
)

type ProductRepositoryImpl struct {
	productDAO *dao.ProductDAOImpl
}

func (this *ProductRepositoryImpl) Init(productDAO *dao.ProductDAOImpl) {
	if this.productDAO == nil {
		this.productDAO = productDAO
	}
}

func (this *ProductRepositoryImpl) SaveProduct(product *data.ProductDO) (*data.ProductDO, *customError.PersistenceError) {
	return this.productDAO.SaveProduct(product)
}

func (this *ProductRepositoryImpl) GetProduct(barcode string) (*data.ProductDO, *customError.PersistenceError) {
	return this.productDAO.GetProduct(barcode)
}

func (this *ProductRepositoryImpl) GetAll() ([]*data.ProductDO, *customError.PersistenceError) {
	return this.productDAO.GetAllProducts()
}

func (this *ProductRepositoryImpl) UpdateProduct(product *data.ProductDO) (*data.ProductDO, *customError.PersistenceError) {
	return this.productDAO.UpdateProduct(product)
}

func (this *ProductRepositoryImpl) DeleteProduct(barcode string) *customError.PersistenceError {
	return this.productDAO.DeleteProduct(barcode)
}
