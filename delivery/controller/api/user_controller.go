package api

import (
	"enigma-laundry-apps/model"
	"enigma-laundry-apps/usecase"
	"enigma-laundry-apps/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	router *gin.Engine
	userUC usecase.UserUseCase
}

func (b *UserController) createHandler(c *gin.Context) {
	var user model.UserCredential
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	user.Id = common.GenerateID()
	if err := b.userUC.RegisterNewUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	userResponse := map[string]any{
		"id":       user.Id,
		"username": user.Username,
		"isActive": user.IsActive,
	}


	c.JSON(http.StatusOK, userResponse)
}

func (b *UserController) listHandler(c *gin.Context) {

}

func NewUserController(r *gin.Engine, usecase usecase.UserUseCase) *UserController {
	controller := UserController{
		router: r,
		userUC: usecase,
	}

	rg := r.Group("/api/v1")
	rg.POST("/users", controller.createHandler)
	rg.GET("/users", controller.listHandler)
	return &controller
}
