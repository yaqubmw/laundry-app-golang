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

type EmployeeController struct {
	router     *gin.Engine
	employeeUC usecase.EmployeeUseCase
}

func (p *EmployeeController) createHandler(c *gin.Context) {
	var employee model.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	employee.Id = common.GenerateID()
	if err := p.employeeUC.RegisterNewEmployee(employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, employee)
}

func (p *EmployeeController) listHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	paginationParam := dto.PaginationParam{
		Page:  page,
		Limit: limit,
	}
	employees, paging, err := p.employeeUC.FindAllEmployee(paginationParam)
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
		"data":   employees,
		"paging": paging,
	})
}

func (p *EmployeeController) getHandler(c *gin.Context) {
	id := c.Param("id")
	uom, err := p.employeeUC.FindByIdEmployee(id)
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

func (p *EmployeeController) updateHandler(c *gin.Context) {
	var employee model.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if err := p.employeeUC.UpdateEmployee(employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, employee)
}

func (p *EmployeeController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	err := p.employeeUC.DeleteEmployee(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.String(http.StatusNoContent, "")
}

func NewEmployeeController(r *gin.Engine, usecase usecase.EmployeeUseCase) *EmployeeController {
	controller := EmployeeController{
		router:     r,
		employeeUC: usecase,
	}

	rg := r.Group("/api/v1")
	rg.POST("/employees", controller.createHandler)
	rg.GET("/employees", controller.listHandler)
	rg.GET("/employees/:id", controller.getHandler)
	rg.PUT("/employees", controller.updateHandler)
	rg.DELETE("/employees/:id", controller.deleteHandler)
	return &controller
}
