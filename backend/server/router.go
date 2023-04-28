package server

import (
	"github.com/VolunteerOne/volunteer-one-app/backend/controllers"
	"github.com/VolunteerOne/volunteer-one-app/backend/database"
	"github.com/VolunteerOne/volunteer-one-app/backend/middleware"
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

	loginRepository := repository.NewLoginRepository(database.GetDatabase())
	usersRepository := repository.NewUsersRepository(database.GetDatabase())
	friendRepository := repository.NewFriendRepository(database.GetDatabase())
	organizationRepository := repository.NewOrganizationRepository(database.GetDatabase())
	eventRepository := repository.NewEventRepository(database.GetDatabase())

	// *********************************************************
	// INITIALIZE SERVICES HERE
	// *********************************************************

	loginService := service.NewLoginService(loginRepository)
	usersService := service.NewUsersService(usersRepository)
	friendService := service.NewFriendService(friendRepository)
	organizationService := service.NewOrganizationService(organizationRepository)
	eventService := service.NewEventService(eventRepository)

	// *********************************************************
	// INITIALIZE CONTROLLERS HERE
	// *********************************************************

	loginController := controllers.NewLoginController(loginService)
	usersController := controllers.NewUsersController(usersService)
	friendController := controllers.NewFriendController(friendService)
	organizationController := controllers.NewOrganizationController(organizationService)
	eventController := controllers.NewEventController(eventService)

	userGroup := router.Group("user")

	// userGroup := new(controllers.UsersController)
	userGroup.POST("/", usersController.Create)
	userGroup.GET("/:id", middleware.BasicAuth, usersController.One)
	userGroup.DELETE("/:id", usersController.Delete)
	userGroup.PUT("/:id", usersController.Update)

	loginGroup := router.Group("login")

	//Simple login, checks database against users inputted email and password to login
	loginGroup.GET("/:email/:password", loginController.Login)
	//Get the users email, sends a forgotten password code to them
	loginGroup.POST("/:email", loginController.SendEmailForPassReset)
	//Get the secret code from the users email, if matches reset password
	loginGroup.PUT("/:email/:resetcode/:newpassword", loginController.PasswordReset)
	//Check valid access token
	loginGroup.POST("/verify", middleware.BasicAuth, loginController.VerifyAccessToken)
	//Get refresh token
	loginGroup.POST("/refresh", loginController.RefreshToken)

	// organizationGroup := router.Group("organization")
	// organizationGroup.POST("/", organizationController.Create)
	// organizationGroup.GET("/", organizationController.All)
	// organizationGroup.GET("/:id", organizationController.One)
	// organizationGroup.DELETE("/:id", organizationController.Delete)
	// organizationGroup.PUT("/:id", organizationController.Update)

	eventGroup := router.Group("event")
	eventGroup.POST("/", eventController.Create)
	eventGroup.GET("/", eventController.All)
	eventGroup.GET("/:id", eventController.One)
	eventGroup.DELETE("/:id", eventController.Delete)
	eventGroup.PUT("/:id", eventController.Update)

	orgUsersGroup := router.Group("orgUsers")
	{
		orgUsers := new(controllers.OrgUsersController)
		orgUsersGroup.POST("/", orgUsers.CreateOrgUser)
		orgUsersGroup.GET("/", orgUsers.ListAllOrgUsers)
		orgUsersGroup.GET("/:id", orgUsers.FindOrgUser)
		orgUsersGroup.PUT("/:id", orgUsers.UpdateOrgUser)
		orgUsersGroup.DELETE("/:id", orgUsers.DeleteOrgUser)
	}

	organizationGroup := router.Group("organization")
	organizationGroup.POST("/", organizationController.Create)
	organizationGroup.GET("/", organizationController.All)
	organizationGroup.GET("/:id", organizationController.One)
	organizationGroup.DELETE("/:id", organizationController.Delete)
	organizationGroup.PUT("/:id", organizationController.Update)

	friendGroup := router.Group("friend")
	friendGroup.POST("/", friendController.Create)
	friendGroup.GET("/", friendController.All)
	friendGroup.GET("/:id", friendController.One)
	friendGroup.DELETE("/:id", friendController.Reject)
	friendGroup.PUT("/:id", friendController.Accept)

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
