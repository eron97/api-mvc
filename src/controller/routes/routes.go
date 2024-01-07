package routes

import "github.com/gin-gonic/gin"

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userId", func(c *gin.Context) {})
	r.GET("/getUserByEmail/:userEmail", func(c *gin.Context) {})
	r.POST("/createUser", func(c *gin.Context) {})
	r.PUT("/updateUser/:userId", func(c *gin.Context) {}, func(c *gin.Context) {})
	r.DELETE("/deleteUser/:userId", func(c *gin.Context) {})
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
