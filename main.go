package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
)

type User struct {
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	BirthDate string `json:"birth_date"`
}

func mainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, admin! You are on the admin main page")
}

func addUser(c echo.Context) error {
	user := User{}

	err := c.Bind(&user)
	if err != nil {
		log.Printf("Failed processing addUser request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("User added: %#v", user)
	return c.String(http.StatusOK, "User added")
}

func main() {
	// Echo instance
	e := echo.New()

	g := e.Group("/admin")

	// Middleware
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
