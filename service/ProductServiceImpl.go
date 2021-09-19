package service

import (
	"API/data"
	"API/errhandler"
	customError "API/error"
	"API/persistence/repository"
	"net/http"
)

type ProductServiceImpl struct {
	productRepository *repository.ProductRepositoryImpl
	errHandler        *errhandler.ErrorHandler
}

func (this *ProductServiceImpl) Init(productRepository *repository.ProductRepositoryImpl, errHandler *errhandler.ErrorHandler) {
	if this.productRepository == nil {
		this.productRepository = productRepository
	}
	if this.errHandler == nil {
		this.errHandler = errHandler
	}
}

func (this *ProductServiceImpl) SaveProduct(product *data.ProductDTO) (*data.ProductDTO, *customError.ServiceError) {
	productToBeSaved := this.convertToDO(product)
	savedProduct, err := this.productRepository.SaveProduct(productToBeSaved)
	if err != nil {
		return nil, customError.CreateServiceError(err.GetMessage(), err.GetCause(), http.StatusInternalServerError)
	}
	return this.convertToDTO(savedProduct), nil
}

func (this *ProductServiceImpl) GetProduct(barcode string) (*data.ProductDTO, *customError.ServiceError) {
	product, err := this.productRepository.GetProduct(barcode)
	if err != nil {
		return nil, customError.CreateServiceError(err.GetMessage(), err.GetCause(), http.StatusInternalServerError)
	}
	return this.convertToDTO(product), nil
}

func (this *ProductServiceImpl) GetAllProducts() ([]*data.ProductDTO, *customError.ServiceError) {
	productsFromDatabase, err := this.productRepository.GetAll()
	if err != nil {
		return nil, customError.CreateServiceError(err.GetMessage(), err.GetCause(), http.StatusInternalServerError)
	}
	var products []*data.ProductDTO
	for _, productFromDatabase := range productsFromDatabase {
		products = append(products, this.convertToDTO(productFromDatabase))
	}
	return products, nil
}

func (this *ProductServiceImpl) UpdateProduct(product *data.ProductDTO) (*data.ProductDTO, *customError.ServiceError) {
	payloadAsDO := this.convertToDO(product)
	updatedProduct, err := this.productRepository.UpdateProduct(payloadAsDO)
	if err != nil {
		return nil, customError.CreateServiceError(err.GetMessage(), err.GetCause(), http.StatusInternalServerError)
	}
	return this.convertToDTO(updatedProduct), nil
}

func (this *ProductServiceImpl) DeleteProduct(barcode string) *customError.ServiceError {
	err := this.productRepository.DeleteProduct(barcode)
	if err != nil {
		return customError.CreateServiceError(err.GetMessage(), err.GetCause(), http.StatusInternalServerError)
	}
	return nil
}

func (this *ProductServiceImpl) convertToDO(productDTO *data.ProductDTO) *data.ProductDO {
	productDO := data.ProductDO{}
	productDO.SetBarcode(productDTO.GetBarcode())
	productDO.SetName(productDTO.GetName())
	productDO.SetUnit(productDTO.GetUnit())
	productDO.SetAmount(productDTO.GetAmount())
	return &productDO
}

func (this *ProductServiceImpl) convertToDTO(productDO *data.ProductDO) *data.ProductDTO {
	return data.CreateProductDTO(productDO.GetBarcode(), productDO.GetName(), productDO.GetUnit(), productDO.GetAmount())
}
