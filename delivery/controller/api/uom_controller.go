package api

import (
	"enigma-laundry-apps/model"
	"enigma-laundry-apps/usecase"

	"github.com/gin-gonic/gin"
)

type UomController struct {
	uomUC  usecase.UomUseCase
	router *gin.Engine
}

func (u *UomController) createHandler(c *gin.Context) {
	var uom model.Uom
	// cek error bind body JSON
	if err := c.ShouldBindJSON(&uom); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return // agar tidak diteruskan ke bawah
	}
	// cek error ketika sever tidak merespon/terjadi kesalahan pada server
	if err := u.uomUC.RegisterNewUom(uom); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return // agar tidak diteruskan ke bawah
	}

	c.JSON(201, uom)

}

func (u *UomController) listHandler(c *gin.Context) {
	uoms, err := u.uomUC.FindAllUom()
	if err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.JSON(200, uoms)
}

func (u *UomController) getHandler(c *gin.Context) {}

func (u *UomController) updateHandler(c *gin.Context) {}

func (u *UomController) deleteHandler(c *gin.Context) {}

func NewUomController(usecase usecase.UomUseCase, r *gin.Engine) *UomController {

	controller := UomController{
		router: r,
		uomUC:  usecase,
	}

	//  daftarkan semua url path disini
	// /uom -> GET, POST, PUT, DELETE
	r.POST("/uoms", controller.createHandler)
	r.GET("/uoms", controller.listHandler)
	r.GET("/uoms/:id", controller.getHandler)
	r.PUT("/uoms", controller.updateHandler)
	r.DELETE("/uoms", controller.deleteHandler)
	return &controller

}
