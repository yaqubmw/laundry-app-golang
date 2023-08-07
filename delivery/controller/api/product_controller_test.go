package api

import (
	"bytes"
	"encoding/json"
	"enigma-laundry-apps/model"
	"enigma-laundry-apps/model/dto"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type productUseCaseMock struct {
	mock.Mock
}

func (p *productUseCaseMock) RegisterNewProduct(payload model.Product) error {
	args := p.Called(payload)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (p *productUseCaseMock) FindAllProduct(requesPaging dto.PaginationParam) ([]model.Product, dto.Paging, error) {
	args := p.Called(requesPaging)
	if args.Get(2) != nil {
		return nil, dto.Paging{}, args.Error(2)
	}
	return args.Get(0).([]model.Product), args.Get(1).(dto.Paging), nil
}

func (p *productUseCaseMock) FindByIdProduct(id string) (model.Product, error) {
	args := p.Called(id)
	if args.Get(1) != nil {
		return model.Product{}, args.Error(1)
	}
	return args.Get(0).(model.Product), nil
}

func (p *productUseCaseMock) UpdateProduct(payload model.Product) error {
	args := p.Called(payload)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (p *productUseCaseMock) DeleteProduct(id string) error {
	args := p.Called(id)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

type ProductControllerTestSuite struct {
	suite.Suite
	usecaseMock *productUseCaseMock
	router      *gin.Engine
}

func (suite *ProductControllerTestSuite) SetupTest() {
	suite.usecaseMock = new(productUseCaseMock)
	suite.router = gin.Default()

}

func (suite *ProductControllerTestSuite) TestCreateHandlerProduct_Success() {
	// pembuatan dummy untuk proses request ke http server
	dummyRequest := dto.ProductRequestDto{
		Name:  "Product A",
		Price: 1000,
		UomId: "1",
	}

	// pembuatan variable product, karena yg diterima UC Create adalah model.Product
	// Harus di assign ulang dari dummyRequest ke model.Product
	var newProduct model.Product
	dummyRequest.Id = "XA2123"
	newProduct.Id = dummyRequest.Id
	newProduct.Name = dummyRequest.Name
	newProduct.Uom.Id = dummyRequest.UomId
	newProduct.Price = dummyRequest.Price

	// pemanggilan usecase mock
	suite.usecaseMock.On("RegisterNewProduct", newProduct).Return(nil)
	NewProductController(suite.router, suite.usecaseMock)

	// pembuatan Recorder untuk merekam keseluruhan respons HTTP yg dikirim Client
	recorder := httptest.NewRecorder()

	// melakukan Marshal dari struct -> JSON. karena mengirim ke HTTP nenggunakan JSON
	payload, _ := json.Marshal(dummyRequest)

	// membuat simulasi request ke HTTP dengan method, path, dan body
	request, _ := http.NewRequest(http.MethodPost, "/api/v1/products", bytes.NewBuffer(payload))

	// simulasi menjalankan server dengan serveHTTP
	suite.router.ServeHTTP(recorder, request)

	// kita tangkap hasil response dari recorder yg dikirim oleh server
	response := recorder.Body.Bytes()

	// setelah response diterima dalam bentuk JSON, kembalikan lagi JSON -> Struct
	// masukan hasilnya ke dalam struct dummyRequest
	actualProduct := dto.ProductRequestDto{}
	json.Unmarshal(response, &actualProduct)

	// lakukan assertion response status yg dikirim
	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.Equal(suite.T(), dummyRequest, actualProduct)
}

func (suite *ProductControllerTestSuite) TestCreateHandlerProduct_BindingError() {
	NewProductController(suite.router, suite.usecaseMock)
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/api/v1/products", nil)
	suite.router.ServeHTTP(recorder, request)
	assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)

}

func (suite *ProductControllerTestSuite) TestCreateHandlerProduct_UsecaseError() {
	var newProduct model.Product
	suite.usecaseMock.On("RegisterNewProduct", newProduct).Return(errors.New("error"))
	NewProductController(suite.router, suite.usecaseMock)
	recorder := httptest.NewRecorder()
	payload, _ := json.Marshal(dto.ProductRequestDto{})
	request, _ := http.NewRequest(http.MethodPost, "/api/v1/products", bytes.NewBuffer(payload))
	suite.router.ServeHTTP(recorder, request)
	response := recorder.Body.Bytes()
	var actualError struct {
		Err string
	}
	json.Unmarshal(response, &actualError)
	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	fmt.Println("actualError:", actualError)
	assert.Equal(suite.T(), "error", actualError.Err)
}


func (suite *ProductControllerTestSuite) TestListHandlerProduct_Success() {
	expectedPaginationParam := dto.PaginationParam{
		Page:  1,
		Limit: 5,
	}

	expectedProduct := []model.Product{{
		Id:    "1",
		Name:  "Product A",
		Price: 1000,
		Uom: model.Uom{
			Id:   "1",
			Name: "Pcs",
		},
	}}

	expectedPaging := dto.Paging{
		Page:        1,
		RowsPerPage: 5,
		TotalRows:   1,
		TotalPages:  1,
	}

	suite.usecaseMock.On("FindAllProduct", expectedPaginationParam).Return(expectedProduct, expectedPaging, nil)
	NewProductController(suite.router, suite.usecaseMock)
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/api/v1/products?page=1&limit=5", nil)
	suite.router.ServeHTTP(recorder, request)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
}

func (suite *ProductControllerTestSuite) TestListHandlerProduct_Fail() {
	suite.usecaseMock.On("FindAllProduct", dto.PaginationParam{Page: 1, Limit: 5}).Return(nil, dto.Paging{}, errors.New("error"))
	NewProductController(suite.router, suite.usecaseMock)
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/api/v1/products?page=1&limit=5", nil)
	suite.router.ServeHTTP(recorder, request)
	response := recorder.Body.Bytes()
	var actualError struct {
		Err string
	}
	json.Unmarshal(response, &actualError)
	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), "error", actualError.Err)
}


func TestProductControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ProductControllerTestSuite))
}
