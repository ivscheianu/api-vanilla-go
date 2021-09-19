package main

import (
	"API/controller"
	"API/errhandler"
	"API/persistence/dao"
	"API/persistence/repository"
	"API/service"
)

func main() {
	var errHandler = errhandler.ErrorHandler{}
	var productDAO = dao.ProductDAOImpl{}
	productDAO.Init(&errHandler)
	var productRepository = repository.ProductRepositoryImpl{}
	productRepository.Init(&productDAO)
	var productService = service.ProductServiceImpl{}
	productService.Init(&productRepository, &errHandler)
	var productController = controller.ProductController{}
	productController.Init(&productService)
}
