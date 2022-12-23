package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"net/http"

	"RestAPI-go/model"
	"RestAPI-go/storage"
)

func mainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, admin! You are on the admin main page")
}

func addUser(c echo.Context) error {
	user := model.User{}

	err := c.Bind(&user)
	if err != nil {
		log.Printf("Failed processing addUser request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("User added: %#v", user)

	log.Printf("%v", storage.GetDBInstance())
	_ = storage.GetDBInstance().Create(&user)

	return c.JSONPretty(http.StatusOK, user, "  ")
}

// print Connected! if connection is successful
func main() {
	_ = storage.DBConn()

	fmt.Println("Connected to db!")

	// Echo instance
	e := echo.New()

	g := e.Group("/admin")

	// Middlewares
	g.Use(middleware.Logger())
	g.Use(middleware.Recover())

	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "admin" && password == "admin" {
			return true, nil
		}

		return false, nil
	}))

	// Routes
	g.GET("/main", mainAdmin)
	g.POST("/users", addUser)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
