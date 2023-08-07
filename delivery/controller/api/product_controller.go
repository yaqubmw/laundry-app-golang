package api

import (
	"enigma-laundry-apps/model"
	"enigma-laundry-apps/model/dto"
	"enigma-laundry-apps/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	router    *gin.Engine
	productUC usecase.ProductUseCase
}

func (p *ProductController) createHandler(c *gin.Context) {
	var productRequest dto.ProductRequestDto
	if err := c.ShouldBindJSON(&productRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	var newProduct model.Product
	// productRequest.Id = common.GenerateID()
	newProduct.Id = productRequest.Id
	newProduct.Name = productRequest.Name
	newProduct.Uom.Id = productRequest.UomId
	newProduct.Price = productRequest.Price
	if err := p.productUC.RegisterNewProduct(newProduct); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, productRequest)
}
func (p *ProductController) listHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	paginationParam := dto.PaginationParam{
		Page:  page,
		Limit: limit,
	}
	products, paging, err := p.productUC.FindAllProduct(paginationParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	status := map[string]any{
		"code":        200,
		"description": "Get All Data Successfully",
	}
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"data":   products,
		"paging": paging,
	})
}

func (p *ProductController) getHandler(c *gin.Context) {
	id := c.Param("id")
	product, err := p.productUC.FindByIdProduct(id)
	if err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	status := map[string]any{
		"code":        200,
		"description": "Get By Id Data Successfully",
	}
	c.JSON(200, gin.H{
		"status": status,
		"data":   product,
	})
}
func (p *ProductController) updateHandler(c *gin.Context) {
	var productRequest dto.ProductRequestDto
	if err := c.ShouldBindJSON(&productRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	var newProduct model.Product
	newProduct.Id = productRequest.Id
	newProduct.Name = productRequest.Name
	newProduct.Uom.Id = productRequest.UomId
	newProduct.Price = productRequest.Price
	if err := p.productUC.UpdateProduct(newProduct); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, productRequest)
}
func (p *ProductController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := p.productUC.DeleteProduct(id); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.String(204, "")
}

func NewProductController(r *gin.Engine, usecase usecase.ProductUseCase) *ProductController {
	controller := ProductController{
		router:    r,
		productUC: usecase,
	}
	// /api/v1/products
	rg := r.Group("/api/v1")
	rg.POST("/products", controller.createHandler)
	rg.GET("/products", controller.listHandler)
	rg.GET("/products/:id", controller.getHandler)
	rg.PUT("/products", controller.updateHandler)
	rg.DELETE("/products/:id", controller.deleteHandler)
	return &controller
}
