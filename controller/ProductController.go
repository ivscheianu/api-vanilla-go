package controller

import (
	"API/data"
	customError "API/error"
	"API/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

const (
	post = "POST"
	get  = "GET"
	put  = "PUT"
	del  = "DELETE"
	barcode = "barcode"
	port = ":8080"
)

type ProductController struct {
	productService *service.ProductServiceImpl
	router         *mux.Router
	waitGroup      sync.WaitGroup
}

func (this *ProductController) Init(productService *service.ProductServiceImpl) {
	this.productService = productService
	this.handleRequests()
}

func (this *ProductController) saveProduct(responseWriter http.ResponseWriter, request *http.Request) {
	reqBody, _ := ioutil.ReadAll(request.Body)
	var product data.ProductDTO
	_ = json.Unmarshal(reqBody, &product)
	var response *data.ProductDTO
	var err *customError.ServiceError
	this.waitGroup.Add(1)
	go func() {
		defer this.waitGroup.Done()
		response, err = this.productService.SaveProduct(&product)
	}()
	this.waitGroup.Wait()
	this.sendResponse(responseWriter, err, response, nil)
}

func (this *ProductController) getProduct(responseWriter http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	barcode := params[barcode]
	this.waitGroup.Add(1)
	var response *data.ProductDTO
	var err *customError.ServiceError
	go func() {
		defer this.waitGroup.Done()
		response, err = this.productService.GetProduct(barcode)
	}()
	this.waitGroup.Wait()
	this.sendResponse(responseWriter, err, response, nil)
}

func (this *ProductController) getAllProducts(responseWriter http.ResponseWriter, request *http.Request) {
	this.waitGroup.Add(1)
	var response []*data.ProductDTO
	var err *customError.ServiceError
	go func() {
		defer this.waitGroup.Done()
		response, err = this.productService.GetAllProducts()
	}()
	this.waitGroup.Wait()
	this.sendResponse(responseWriter, err, response, nil)
}

func (this *ProductController) updateProduct(responseWriter http.ResponseWriter, request *http.Request) {
	reqBody, _ := ioutil.ReadAll(request.Body)
	var product data.ProductDTO
	_ = json.Unmarshal(reqBody, &product)
	this.waitGroup.Add(1)
	var response *data.ProductDTO
	var err *customError.ServiceError
	go func() {
		defer this.waitGroup.Done()
		response, err = this.productService.UpdateProduct(&product)
	}()
	this.waitGroup.Wait()
	this.sendResponse(responseWriter, err, response, nil)
}

func (this *ProductController) deleteProduct(responseWriter http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	barcode := params["barcode"]
	var err *customError.ServiceError
	var message = "Successfully deleted"
	this.waitGroup.Add(1)
	go func() {
		defer this.waitGroup.Done()
		err = this.productService.DeleteProduct(barcode)
	}()
	this.waitGroup.Wait()
	this.sendResponse(responseWriter, err, nil, &message)
}

func (this *ProductController) sendResponse(responseWriter http.ResponseWriter, err *customError.ServiceError, response interface{}, message *string) {
	if err != nil {
		_, _ = responseWriter.Write([]byte(err.GetMessage()))
		responseWriter.WriteHeader(err.GetStatusCode())
	} else {
		if message != nil {
			_, _ = responseWriter.Write([]byte(*message))
		}
		if response != nil {
			_ = json.NewEncoder(responseWriter).Encode(response)
		}
		responseWriter.WriteHeader(http.StatusOK)
	}
}

func (this *ProductController) handleRequests() {
	this.router = mux.NewRouter().StrictSlash(true)
	this.router.HandleFunc("/", this.saveProduct).Methods(post)
	this.router.HandleFunc("/", this.getAllProducts).Methods(get)
	this.router.HandleFunc("/{barcode}", this.getProduct).Methods(get)
	this.router.HandleFunc("/{barcode}", this.deleteProduct).Methods(del)
	this.router.HandleFunc("/", this.updateProduct).Methods(put)
	log.Fatal(http.ListenAndServe(port, this.router))
}
