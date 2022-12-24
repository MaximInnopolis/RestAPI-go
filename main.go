package main

import (
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

	e.POST("/login", login)

	e.GET("/user/:id", readUser, middleware.KeyAuth(basicAccess))
	e.POST("/user", createUser, middleware.KeyAuth(supremeAccess))
	e.DELETE("/user/:id", deleteUser, middleware.KeyAuth(supremeAccess))
	e.PUT("/user/:id", updateUser, middleware.KeyAuth(supremeAccess))

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
