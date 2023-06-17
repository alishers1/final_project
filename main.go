package main

import (
	"posts/auth"
	"posts/config"
	"posts/controller"
	"posts/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.New()
	godotenv.Load()
	config.ConnectToDB()

	v1 := r.Group("/auth")
	{
		v1.POST("/signup", auth.SignUp)
		v1.POST("/signin", auth.SignIn)
	}

	v1 = r.Group("/users", middleware.RequireAuth)
	{
		v1.GET("/", controller.GetUsersByFilter)
		v1.POST("/", controller.CreateUser)
		v1.PUT("/:id", controller.UpdateUser)
		v1.DELETE("/:id", controller.DeleteUser)
	}
	v1 = r.Group("/clients", middleware.RequireAuth)
	{
		v1.GET("/", controller.GetClientsByFilter)
		v1.GET("/stats", controller.ClientStats)
		v1.POST("/", controller.CreateClient)
		v1.PUT("/:id", controller.UpdateClient)
		v1.DELETE("/:id", controller.DeleteClient)
	}

	r.Run()
}
