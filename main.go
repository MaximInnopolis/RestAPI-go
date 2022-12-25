package main

import (
	"RestAPI-go/posts"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"RestAPI-go/storage"
)

func main() {

	storage.DBConn()

	// Echo instance
	e := echo.New()

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/login", posts.Login)

	e.GET("/user/:id", posts.ReadUser, middleware.KeyAuth(posts.BasicAccess))
	e.POST("/user", posts.CreateUser, middleware.KeyAuth(posts.SupremeAccess))
	e.DELETE("/user/:id", posts.DeleteUser, middleware.KeyAuth(posts.SupremeAccess))
	e.PUT("/user/:id", posts.UpdateUser, middleware.KeyAuth(posts.SupremeAccess))

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
