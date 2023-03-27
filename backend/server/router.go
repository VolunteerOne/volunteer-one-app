package server

import (
	"github.com/VolunteerOne/volunteer-one-app/backend/controllers"
	"github.com/VolunteerOne/volunteer-one-app/backend/database"
	"github.com/VolunteerOne/volunteer-one-app/backend/repository"
	"github.com/VolunteerOne/volunteer-one-app/backend/service"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

    // *********************************************************
    // INITIALIZE REPOSITORIES HERE -> DB migration is handled in main.go
    // *********************************************************

	userGroup := router.Group("user")
	{
		user := new(controllers.UsersController)
		userGroup.POST("/", user.Create)
		userGroup.GET("/", user.All)
		userGroup.GET("/:id", user.One)
		userGroup.DELETE("/:id", user.Delete)
		userGroup.PUT("/:id", user.Update)
	}

    loginRepository := repository.NewLoginRepository(database.GetDatabase())

    // *********************************************************
    // INITIALIZE SERVICES HERE
    // *********************************************************

    loginService := service.NewLoginService(loginRepository)

    // *********************************************************
    // INITIALIZE CONTROLLERS HERE
    // *********************************************************

    loginController := controllers.NewLoginController(loginService)

    router.POST("/signup", loginController.Signup)
    
    loginGroup := router.Group("login")  
	loginGroup.GET("/:email/:password", loginController.Login)

    

	// objectGroup := router.Group("object")
	// {
	// 	object := new(controllers.ObjectController)
	// 	objectGroup.POST("/", object.Create)
	// 	objectGroup.GET("/", object.All)
	// 	objectGroup.GET("/:id", object.One)
	// 	objectGroup.DELETE("/:id", object.Delete)
	// 	objectGroup.PUT("/:id", object.Update)
	// }

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
