package dao

import (
	"API/data"
	"API/errhandler"
	customError "API/error"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql              = "mysql"
	dbUser             = "root"
	dbPassword         = "admin"
	dbAddress          = "localhost"
	dbName             = "products"
	saveFailed         = "Failed to save product"
	getAllFailed       = "Failed to retrieve all products"
	getByBarcodeFailed = "Failed to retrieve product by barcode"
	updateFailed       = "Failed to update product"
	deleteFailed       = "Failed to delete product"
	notFound           = "Product not found"
	nothingChanged     = "The payload is the same as the db object"
	getAllQuery        = "SELECT * FROM product"
	getByBarcodeQuery  = "SELECT * FROM product p WHERE p.barcode = ?"
	insertQuery        = "INSERT INTO product (barcode, name, unit, amount) VALUES (?, ?, ?, ?)"
	updateQuery        = "UPDATE product SET name = ?, unit = ?, amount = ? WHERE barcode = ?"
	deleteQuery        = "DELETE FROM product WHERE barcode = ?"
	isFatal            = true
	isRecoverable      = false
)

type ProductDAOImpl struct {
	database   *sql.DB
	errHandler *errhandler.ErrorHandler
}

func (this *ProductDAOImpl) Init(errHandler *errhandler.ErrorHandler) {
	if this.errHandler == nil {
		this.errHandler = errHandler
	}
	if this.database == nil {
		var err error
		this.database, err = sql.Open(mysql, fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPassword, dbAddress, dbName))
		this.errHandler.HandleError(err, isFatal)
	}
}

func (this *ProductDAOImpl) SaveProducts(products []*data.ProductDO) (savedProducts []*data.ProductDO, persistenceError *customError.PersistenceError) {
	var err error
	insert, err := this.database.Prepare(insertQuery)
	defer insert.Close()
	if err != nil {
		this.errHandler.HandleError(err, isRecoverable)
		return savedProducts, customError.CreatePersistenceError(saveFailed, err)
	}
	for _, product := range products {
		_, err = insert.Exec(product.GetBarcode(), product.GetName(), product.GetUnit(), product.GetAmount())
		if err != nil {
			this.errHandler.HandleError(err, isRecoverable)
			return savedProducts, customError.CreatePersistenceError(saveFailed, err)
		}
		savedProducts = append(savedProducts, product)
	}
	return savedProducts, nil
}

func (this *ProductDAOImpl) SaveProduct(product *data.ProductDO) (savedProduct *data.ProductDO, persistenceError *customError.PersistenceError) {
	savedProducts, err := this.SaveProducts([]*data.ProductDO{product})
	if err != nil {
		this.errHandler.HandleError(err, isRecoverable)
		return nil, err
	}
	return savedProducts[0], err
}

func (this *ProductDAOImpl) GetAllProducts() (products []*data.ProductDO, persistenceError *customError.PersistenceError) {
	var err error
	result, err := this.database.Query(getAllQuery)
	defer result.Close()
	if err != nil {
		this.errHandler.HandleError(err, isRecoverable)
		return products, customError.CreatePersistenceError(getAllFailed, err)
	}
	for result.Next() {
		var productDO *data.ProductDO
		var id, amount uint64
		var barcode, name, unit string
		err = result.Scan(&id, &barcode, &name, &unit, &amount)
		if err != nil {
			this.errHandler.HandleError(err, isRecoverable)
			return products, customError.CreatePersistenceError(getAllFailed, err)
		}
		productDO = data.CreateProductDO(id, barcode, name, unit, amount)
		products = append(products, productDO)
	}
	return products, nil
}

func (this *ProductDAOImpl) GetProduct(barcode string) (product *data.ProductDO, persistenceError *customError.PersistenceError) {
	var err error
	result, err := this.database.Query(getByBarcodeQuery, barcode)
	defer result.Close()
	if err != nil {
		this.errHandler.HandleError(err, isRecoverable)
		return nil, customError.CreatePersistenceError(getByBarcodeFailed, err)
	}
	for result.Next() {
		var id, amount uint64
		var barcode, name, unit string
		err = result.Scan(&id, &barcode, &name, &unit, &amount)
		if err != nil {
			this.errHandler.HandleError(err, isRecoverable)
			return nil, customError.CreatePersistenceError(getByBarcodeFailed, err)
		}
		product = data.CreateProductDO(id, barcode, name, unit, amount)
	}
	if product == nil {
		return nil, customError.CreatePersistenceError(notFound, nil)
	}
	return product, nil
}

func (this *ProductDAOImpl) UpdateProduct(updatedProductFromPayload *data.ProductDO) (updatedProduct *data.ProductDO, persistenceError *customError.PersistenceError) {
	var err error
	update, err := this.database.Prepare(updateQuery)
	defer update.Close()
	if err != nil {
		this.errHandler.HandleError(err, isRecoverable)
		return nil, customError.CreatePersistenceError(updateFailed, err)
	}
	var result sql.Result
	result, err = update.Exec(updatedProductFromPayload.GetName(), updatedProductFromPayload.GetUnit(), updatedProductFromPayload.GetAmount(), updatedProductFromPayload.GetBarcode())
	if err != nil {
		this.errHandler.HandleError(err, isRecoverable)
		return nil, customError.CreatePersistenceError(updateFailed, err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		_, notFoundErr := this.GetProduct(updatedProductFromPayload.GetBarcode())
		if notFoundErr != nil && notFoundErr.GetMessage() == notFound {
			return nil, customError.CreatePersistenceError(notFound, nil)
		}
		return nil, customError.CreatePersistenceError(nothingChanged, nil)
	}
	return updatedProductFromPayload, nil
}

func (this *ProductDAOImpl) DeleteProduct(barcode string) (persistenceError *customError.PersistenceError) {
	drop, err := this.database.Prepare(deleteQuery)
	defer drop.Close()
	if err != nil {
		this.errHandler.HandleError(err, isRecoverable)
		return customError.CreatePersistenceError(deleteFailed, err)
	}
	var result sql.Result
	result, err = drop.Exec(barcode)
	if err != nil {
		this.errHandler.HandleError(err, isRecoverable)
		return customError.CreatePersistenceError(deleteFailed, err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return customError.CreatePersistenceError(notFound, nil)
	}
	return nil
}
