package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"uno/src/services"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var userInput services.UserInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.userService.CreateUser(userInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (uc *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := uc.userService.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var userInput services.UserInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.userService.UpdateUser(id, userInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if err := uc.userService.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (uc *UserController) RegisterRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("/", uc.CreateUser)
		userGroup.GET("/:id", uc.GetUser)
		userGroup.PUT("/:id", uc.UpdateUser)
		userGroup.DELETE("/:id", uc.DeleteUser)
	}
}
