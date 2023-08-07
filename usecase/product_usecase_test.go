package usecase

import (
	"enigma-laundry-apps/model"
	"enigma-laundry-apps/model/dto"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type repoMock struct {
	mock.Mock
}

func (r *repoMock) Create(payload model.Product) error {
	args := r.Called(payload)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (r *repoMock) List() ([]model.Product, error) {
	args := r.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Product), nil
}

func (r *repoMock) Get(id string) (model.Product, error) {
	args := r.Called(id)
	if args.Get(1) != nil {
		return model.Product{}, args.Error(1)
	}
	return args.Get(0).(model.Product), nil
}

func (r *repoMock) Update(payload model.Product) error {
	args := r.Called(payload)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}
func (r *repoMock) Delete(id string) error {
	args := r.Called(id)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}
func (r *repoMock) Paging(requestPaging dto.PaginationParam) ([]model.Product, dto.Paging, error) {
	args := r.Called(requestPaging)
	if args.Get(2) != nil {
		return nil, dto.Paging{}, args.Error(2)
	}
	return args.Get(0).([]model.Product), args.Get(1).(dto.Paging), nil
}

type usecaseMock struct {
	mock.Mock
}

// DeleteUom implements UomUseCase.
func (*usecaseMock) DeleteUom(id string) error {
	panic("unimplemented")
}

// FindAllUom implements UomUseCase.
func (*usecaseMock) FindAllUom() ([]model.Uom, error) {
	panic("unimplemented")
}

// RegisterNewUom implements UomUseCase.
func (*usecaseMock) RegisterNewUom(payload model.Uom) error {
	panic("unimplemented")
}

// UpdateUom implements UomUseCase.
func (*usecaseMock) UpdateUom(payload model.Uom) error {
	panic("unimplemented")
}

// Mock Usecase
// Karena yang dibutuhkan adalah get id uom di register
func (u *usecaseMock) FindByIdUom(id string) (model.Uom, error) {
	args := u.Called(id)
	if args.Get(1) != nil {
		return model.Uom{}, args.Error(1)
	}
	return args.Get(0).(model.Uom), nil
}

type ProductUseCaseTestSuite struct {
	suite.Suite
	repoMock    *repoMock
	usecaseMock *usecaseMock
}

func (suite *ProductUseCaseTestSuite) SetupTest() {
	suite.repoMock = new(repoMock)
	suite.usecaseMock = new(usecaseMock)
}

// Test Case
var productDummy = []model.Product{
	{
		Id:    "1",
		Name:  "Product A",
		Price: 10000,
		Uom:   model.Uom{Id: "1"},
	},
	{
		Id:    "2",
		Name:  "Product B",
		Price: 50000,
		Uom:   model.Uom{Id: "1"},
	},
	{
		Id:    "3",
		Name:  "Product C",
		Price: 5000,
		Uom:   model.Uom{Id: "1"},
	},
}

var uomDummy = model.Uom{
	Id:   "1",
	Name: "Pcs",
}

func (suite *ProductUseCaseTestSuite) TestRegisterNewProduct_Success() {
	dmProduct := productDummy[0]
	suite.usecaseMock.On("FindByIdUom", dmProduct.Uom.Id).Return(uomDummy, nil)
	dmProduct.Uom = uomDummy
	suite.repoMock.On("Create", dmProduct).Return(nil)
	usecase := NewProductUseCase(suite.repoMock, suite.usecaseMock)
	err := usecase.RegisterNewProduct(dmProduct)
	assert.Nil(suite.T(), err)
}

func (suite *ProductUseCaseTestSuite) TestRegisterNewProduct_EmptyField() {
	suite.repoMock.On("Create", model.Product{}).Return(fmt.Errorf("field requierd"))
	usecase := NewProductUseCase(suite.repoMock, suite.usecaseMock)
	err := usecase.RegisterNewProduct(model.Product{})
	assert.Error(suite.T(), err)
}

func (suite *ProductUseCaseTestSuite) TestRegisterNewProduct_InvalidUOM() {
	dummy := productDummy[0]
	suite.usecaseMock.On("FindByIdUom", "1xxx").Return(model.Uom{}, fmt.Errorf("uom not found"))
	usecase := NewProductUseCase(suite.repoMock, suite.usecaseMock)
	dummy.Uom.Id = "1xxx"
	err := usecase.RegisterNewProduct(dummy)
	assert.Error(suite.T(), err)
}

func (suite *ProductUseCaseTestSuite) TestRegisterNewProduct_Fail() {
	suite.usecaseMock.On("FindByIdUom", "1").Return(uomDummy, nil)
	productDummy[0].Uom = uomDummy
	suite.repoMock.On("Create", productDummy[0]).Return(fmt.Errorf("failed register"))
	usecase := NewProductUseCase(suite.repoMock, suite.usecaseMock)
	err := usecase.RegisterNewProduct(productDummy[0])
	assert.Error(suite.T(), err)
}

func TestProductUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(ProductUseCaseTestSuite))
}
