package controller

import (
	"fmt"
	"gin-gonic-crud/models/dao"
	"gin-gonic-crud/models/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create ...
func Create(c *gin.Context) {
	fmt.Println(c)
	fmt.Println("----")
	var user model.User
	fmt.Println(&user)
	c.BindJSON(&user)
	err := dao.CreateUser(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, user)
}

// GetAll ...
func GetAll(c *gin.Context) {
	var user []model.User
	err := dao.GetAllUser(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, user)
}

// GetByID ...
func GetByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user model.User
	err := dao.GetByIDUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, user)

}

// Update ...
func Update(c *gin.Context) {
	var user model.User
	id := c.Params.ByName("id")
	err := dao.GetByIDUser(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = dao.UpdateUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, user)
}

// Delete ...
func Delete(c *gin.Context) {
	var user model.User
	id := c.Params.ByName("id")
	err := dao.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
}
