package repository

import (
	"database/sql"
	"enigma-laundry-apps/model"
	"enigma-laundry-apps/model/dto"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

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

type ProductRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mocksql sqlmock.Sqlmock
	repo    ProductRepository
}

func (suite *ProductRepositoryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Printf("an error %s was not expected when opening a stub database connection", err)
	}
	suite.mockDb = db
	suite.mocksql = mock
	suite.repo = NewProductRepository(db)
}

func (suite *ProductRepositoryTestSuite) TearDownTest() {
	suite.mockDb.Close()
}

func (suite *ProductRepositoryTestSuite) TestCreateNewProduct_Success() {
	dummy := productDummy[0]
	suite.mocksql.ExpectExec("INSERT INTO product (.+)").WithArgs(dummy.Id, dummy.Name, dummy.Price, dummy.Uom.Id).WillReturnResult(sqlmock.NewResult(1, 1))
	actualError := suite.repo.Create(dummy)
	assert.Nil(suite.T(), actualError)
	assert.NoError(suite.T(), actualError)
}

func (suite *ProductRepositoryTestSuite) TestCreateNewProduct_Fail() {
	dummy := productDummy[0]
	suite.mocksql.ExpectExec("INSERT INTO product (.+)").WithArgs(dummy.Id, dummy.Name, dummy.Price, dummy.Uom.Id).WillReturnError(fmt.Errorf("error"))
	actualError := suite.repo.Create(dummy)
	assert.Error(suite.T(), actualError)
}

func (suite *ProductRepositoryTestSuite) TestListProductProduct_Success() {
	rows := sqlmock.NewRows([]string{"id", "name", "price", "uom_id", "uom_name"})
	for _, product := range productDummy {
		rows.AddRow(product.Id, product.Name, product.Price, product.Uom.Id, product.Uom.Name)
	}
	suite.mocksql.ExpectQuery("SELECT (.+) FROM product (.+)").WillReturnRows(rows)
	products, err := suite.repo.List()
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), products, 3)
	assert.Equal(suite.T(), productDummy[0], products[0])
	assert.Equal(suite.T(), productDummy[1], products[1])
	assert.Equal(suite.T(), productDummy[2], products[2])
}

func (suite *ProductRepositoryTestSuite) TestListProductProduct_Fail() {
	suite.mocksql.ExpectQuery("SELECT (.+) FROM product (.+)").WillReturnError(fmt.Errorf("error"))
	products, err := suite.repo.List()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), products)
}

func (suite *ProductRepositoryTestSuite) TestGetProduct_Success() {
	expectedProduct := productDummy[0]
	rows := sqlmock.NewRows([]string{"id", "name", "price", "uom_id", "uom_name"})
	rows.AddRow(expectedProduct.Id, expectedProduct.Name, expectedProduct.Price, expectedProduct.Uom.Id, expectedProduct.Uom.Name)
	suite.mocksql.ExpectQuery("SELECT (.+) FROM product (.+) WHERE p.id = ?").WithArgs(expectedProduct.Id).WillReturnRows(rows)
	actualProduct, actualError := suite.repo.Get(expectedProduct.Id)
	assert.NoError(suite.T(), actualError)
	assert.Nil(suite.T(), actualError)
	assert.Equal(suite.T(), expectedProduct, actualProduct)
}

func (suite *ProductRepositoryTestSuite) TestGetProduct_fail() {
	suite.mocksql.ExpectQuery("SELECT (.+) FROM product (.+) WHERE p.id = ?").WithArgs("1sss").WillReturnError(fmt.Errorf("error"))
	actualProduct, actualError := suite.repo.Get("1sss")
	assert.Error(suite.T(), actualError)
	assert.Equal(suite.T(), model.Product{}, actualProduct)
}

func (suite *ProductRepositoryTestSuite) TestDeleteProduct_Success() {
	suite.mocksql.ExpectExec("DELETE FROM product WHERE id = ?").WithArgs(productDummy[0].Id).WillReturnResult(sqlmock.NewResult(1, 1))
	actualError := suite.repo.Delete(productDummy[0].Id)
	assert.Nil(suite.T(), actualError)
}
func (suite *ProductRepositoryTestSuite) TestDeleteProduct_Fail() {
	suite.mocksql.ExpectExec("DELETE FROM product WHERE id = ?").WithArgs(productDummy[0].Id).WillReturnError(fmt.Errorf("error"))
	actualError := suite.repo.Delete(productDummy[0].Id)
	assert.Error(suite.T(), actualError)
}

func (suite *ProductRepositoryTestSuite) TestPagingProduct_Success() {
	requestPaging := dto.PaginationParam{
		Page:  1,
		Limit: 5,
	}
	paginationQuery := dto.PaginationQuery{
		Take: 5,
		Skip: 0,
	}
	rows := sqlmock.NewRows([]string{"id", "name", "price", "uom_id", "uom_name"})
	for _, product := range productDummy {
		rows.AddRow(product.Id, product.Name, product.Price, product.Uom.Id, product.Uom.Name)
	}
	suite.mocksql.ExpectQuery("SELECT (.+) FROM product (.+) ").WithArgs(paginationQuery.Take, paginationQuery.Skip).WillReturnRows(rows)

	// COUNT
	rowCount := sqlmock.NewRows([]string{"count"})
	rowCount.AddRow(3)
	suite.mocksql.ExpectQuery(regexp.QuoteMeta("SELECT COUNT(*) FROM product")).WillReturnRows(rowCount)

	actualProduct, actualPaging, actualError := suite.repo.Paging(requestPaging)
	assert.Nil(suite.T(), actualError)
	assert.NotNil(suite.T(), actualProduct)
	assert.Equal(suite.T(), actualPaging.TotalRows, 3)
}

func (suite *ProductRepositoryTestSuite) TestPagingProduct_QueryPagingError() {
	suite.mocksql.ExpectQuery("SELECT (.+) FROM product (.+) ").WillReturnError(fmt.Errorf("error"))
	actualProduct, actualPaging, actualError := suite.repo.Paging(dto.PaginationParam{})
	assert.Error(suite.T(), actualError)
	assert.Nil(suite.T(), actualProduct)
	assert.Equal(suite.T(), actualPaging.TotalRows, 0)
}

func (suite *ProductRepositoryTestSuite) TestPagingProduct_QueryCountError() {
	requestPaging := dto.PaginationParam{
		Page:  1,
		Limit: 5,
	}
	paginationQuery := dto.PaginationQuery{
		Take: 5,
		Skip: 0,
	}
	rows := sqlmock.NewRows([]string{"id", "name", "price", "uom_id", "uom_name"})
	for _, product := range productDummy {
		rows.AddRow(product.Id, product.Name, product.Price, product.Uom.Id, product.Uom.Name)
	}
	suite.mocksql.ExpectQuery("SELECT (.+) FROM product (.+) ").WithArgs(paginationQuery.Take, paginationQuery.Skip).WillReturnRows(rows)

	// COUNT
	rowCount := sqlmock.NewRows([]string{"count"})
	rowCount.AddRow(3)
	suite.mocksql.ExpectQuery(regexp.QuoteMeta("SELECT COUNT(*) FROM product")).WillReturnError(fmt.Errorf("error"))
	_, actualPaging, actualError := suite.repo.Paging(requestPaging)
	assert.Error(suite.T(), actualError)
	assert.Equal(suite.T(), dto.Paging{}, actualPaging)
}

func TestProductRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ProductRepositoryTestSuite))
}
