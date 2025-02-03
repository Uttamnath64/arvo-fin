package routes

import "github.com/Uttamnath64/arvo-fin/fin-api/internal/handlers"

func (routes *Routes) UserRoutes() {

	userHandler := handlers.NewUserHandler(routes.container)
	userGroup := routes.server.Group("/users")
	{
		userGroup.GET("/", userHandler.GetUsersHandler)         // List all users
		userGroup.POST("/", userHandler.CreateUserHandler)      // Create a new user
		userGroup.GET("/:id", userHandler.GetUserByIDHandler)   // Get user by ID
		userGroup.PUT("/:id", userHandler.UpdateUserHandler)    // Update user
		userGroup.DELETE("/:id", userHandler.DeleteUserHandler) // Delete user
	}
}
