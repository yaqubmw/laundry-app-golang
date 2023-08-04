package api

import (
	"enigma-laundry-apps/model"
	"enigma-laundry-apps/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	router  *gin.Engine
	usecase usecase.AuthUseCase
}

func (a *AuthController) loginHandler(c *gin.Context) {
	var payload model.UserCredential
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := a.usecase.Login(payload.Username, payload.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"token": token})
}

func NewAuthController(r *gin.Engine, usecase usecase.AuthUseCase) *AuthController {
	controller := AuthController{
		router:  r,
		usecase: usecase,
	}
	rg := r.Group("/api/v1")
	rg.POST("/login", controller.loginHandler)
	return &controller
}
