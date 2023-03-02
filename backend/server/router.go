package server

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// router.Use(middlewares.AuthMiddleware())

	// root := new(controllers.RootController)
	// router.GET("/", root.Get)

	// example := new(controllers.ExampleController)
	// router.GET("/example", controllers.ExampleGet)

	// v1 := router.Group("v1")
	// {
	// 	userGroup := v1.Group("user")
	// 	{
	// 		user := new(controllers.UserController)
	// 		userGroup.GET("/:id", user.Retrieve)
	// 	}
	// }

	return router

}
