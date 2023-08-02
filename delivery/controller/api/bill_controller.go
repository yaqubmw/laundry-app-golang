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

type BillController struct {
	router *gin.Engine
	billUC usecase.BillUseCase
}

func (b *BillController) createHandler(c *gin.Context) {
	var bill model.Bill
	if err := c.ShouldBindJSON(&bill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	bill.Id = common.GenerateID()
	if err := b.billUC.RegisterNewBill(bill); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bill)
}
func (b *BillController) listHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	paginationParam := dto.PaginationParam{
		Page:  page,
		Limit: limit,
	}
	bills, paging, err := b.billUC.FindAllBill(paginationParam)
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
		"data":   bills,
		"paging": paging,
	})
}
func (b *BillController) getHandler(c *gin.Context) {
	id := c.Param("id")
	bill, err := b.billUC.FindByIdBill(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	status := map[string]any{
		"code":        200,
		"description": "Get By Id Data Successfully",
	}
	c.JSON(200, gin.H{
		"status": status,
		"data":   bill,
	})
}

func NewBillController(r *gin.Engine, usecase usecase.BillUseCase) *BillController {
	controller := BillController{
		router: r,
		billUC: usecase,
	}

	rg := r.Group("/api/v1")
	rg.POST("/bills", controller.createHandler)
	rg.GET("/bills", controller.listHandler)
	rg.GET("/bills/:id", controller.getHandler)
	return &controller
}
