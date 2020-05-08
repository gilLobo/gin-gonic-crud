package routes

import (
	controller "gin-gonic-crud/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter ..
func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.POST("user", controller.Create)
		v1.GET("user", controller.GetAll)
		v1.GET("user/:id", controller.GetByID)
		v1.PUT("user/:id", controller.Update)
		v1.DELETE("user/:id", controller.Delete)
	}
	return r
}
