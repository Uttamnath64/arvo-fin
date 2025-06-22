package routes

import "github.com/Uttamnath64/arvo-fin/fin-api/internal/handlers"

func (routes *Routes) UserRoutes() {

	userHandler := handlers.NewUserHandler(routes.container)
	userGroup := routes.server.Group("/user")
	{
		userGroup.GET("/", userHandler.GetUsersHandler)
		userGroup.GET("/:id", userHandler.GetUserByIDHandler)

		userGroup.POST("/", userHandler.CreateUserHandler)
		userGroup.PUT("/:id", userHandler.UpdateUserHandler)
		userGroup.DELETE("/:id", userHandler.DeleteUserHandler)
	}
}
