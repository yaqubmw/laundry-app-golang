package api

import (
	"enigma-laundry-apps/model"
	"enigma-laundry-apps/model/dto"
	"enigma-laundry-apps/usecase"
	"enigma-laundry-apps/utils/common"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	router     *gin.Engine
	customerUC usecase.CustomerUseCase
}

func (p *CustomerController) createHandler(c *gin.Context) {
	var customer model.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	customer.Id = common.GenerateID()
	if err := p.customerUC.RegisterNewCustomer(customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, customer)
}

func (p *CustomerController) listHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	paginationParam := dto.PaginationParam{
		Page:  page,
		Limit: limit,
	}
	customers, paging, err := p.customerUC.FindAllCustomer(paginationParam)
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
		"data":   customers,
		"paging": paging,
	})
}

func (p *CustomerController) getHandler(c *gin.Context) {
	id := c.Param("id")
	uom, err := p.customerUC.FindByIdCustomer(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	status := map[string]any{
		"code":        200,
		"description": "Get by Id Data Successfully",
	}

	c.JSON(200, gin.H{
		"status": status,
		"data":   uom,
	})
}

func (p *CustomerController) updateHandler(c *gin.Context) {
	var customer model.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if err := p.customerUC.UpdateCustomer(customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, customer)
}

func (p *CustomerController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	err := p.customerUC.DeleteCustomer(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.String(http.StatusNoContent, "")
}

func NewCustomerController(r *gin.Engine, usecase usecase.CustomerUseCase) *CustomerController {
	controller := CustomerController{
		router:     r,
		customerUC: usecase,
	}

	rg := r.Group("/api/v1")
	rg.POST("/customers", controller.createHandler)
	rg.GET("/customers", controller.listHandler)
	rg.GET("/customers/:id", controller.getHandler)
	rg.PUT("/customers", controller.updateHandler)
	rg.DELETE("/customers/:id", controller.deleteHandler)
	return &controller
}
