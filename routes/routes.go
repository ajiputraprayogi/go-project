package routes

import (
	"go-project/helper"
	"go-project/controllers"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"go-project/middleware"

)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	}))

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.POST("/logout", controllers.Logout)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		r.POST("/users", controllers.CreateUser)
		r.GET("/users", controllers.GetUsers)
		r.GET("/users/:id", controllers.GetUser)
		r.PUT("/users/:id", controllers.UpdateUser)
		r.DELETE("/users/:id", controllers.DeleteUser)
	
		r.POST("/posts", controllers.CreatePost)
		r.GET("/posts", controllers.GetPosts)
		r.GET("/posts/:id", controllers.GetPost)
		r.PUT("/posts/:id", controllers.UpdatePost)
		r.DELETE("/posts/:id", controllers.DeletePost)
	}

	global := r.Group("/global")
	// global.Use(middleware.AuthMiddleware())
	{
		global.GET("/users-options", helper.GetUserOptions)
	}

	return r
}
