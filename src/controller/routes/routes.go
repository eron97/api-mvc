package routes

import (
	"github.com/eron97/api-mvc.git/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userId", controller.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)
}

/*

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userId", func(c *gin.Context) {})
	r.GET("/getUserByEmail/:userEmail", func(c *gin.Context) {})
	r.POST("/createUser", func(c *gin.Context) {})
	r.PUT("/updateUser/:userId", func(c *gin.Context) {}, func(c *gin.Context) {})
	r.DELETE("/deleteUser/:userId", func(c *gin.Context) {})
}

*/
