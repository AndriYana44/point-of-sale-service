package controllers

import "github.com/gin-gonic/gin"

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (controller *UserController) GetUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get users successfully",
	})
}

func (controller *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, gin.H{
		"message": "Get user successfully",
		"id":      id,
	})
}

func (controller *UserController) CreateUser(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "User created successfully",
	})
}
