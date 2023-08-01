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
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if err := p.productUC.RegisterNewProduct(product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, product)
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

func (p *ProductController) getHandler(c *gin.Context) {}

func (p *ProductController) updateHandler(c *gin.Context) {}

func (p *ProductController) deleteHandler(c *gin.Context) {}

func NewProductController(r *gin.Engine, usecase usecase.ProductUseCase) *ProductController {
	controller := ProductController{
		router:    r,
		productUC: usecase,
	}

	rg := r.Group("/api/v1")
	rg.POST("/products", controller.createHandler)
	rg.GET("/products", controller.listHandler)
	rg.GET("/products/:id", controller.getHandler)
	rg.PUT("/products", controller.updateHandler)
	rg.DELETE("/products/:id", controller.deleteHandler)
	return &controller
}
