package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"

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

	_ = storage.GetDBInstance().Create(&user)

	return c.JSONPretty(http.StatusOK, user, "  ")
}

func readUser(c echo.Context) error {
	user := model.User{}
	id := c.Param("id")

	storage.GetDBInstance().First(&user, id)

	if user.ID == 0 {
		return c.String(http.StatusNotFound, "Not found")
	}

	return c.JSONPretty(http.StatusOK, user, "  ")
}

func deleteUser(c echo.Context) error {
	user := model.User{}
	id := c.Param("id")
	db := storage.GetDBInstance()

	db.First(&user, id)

	if user.ID == 0 {
		return c.String(http.StatusNotFound, "Not found")
	}

	db.Delete(&user)

	return c.String(http.StatusOK, "")
}

func updateUser(c echo.Context) error {
	user := model.User{}
	id := c.Param("id")
	db := storage.GetDBInstance()

	db.First(&user, id)

	if user.ID == 0 {
		return c.String(http.StatusNotFound, "Not found")
	}

	err := c.Bind(&user)
	if err != nil {
		log.Printf("Failed processing updateUser request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	db.Save(&user)

	return c.String(http.StatusOK, "user with id "+id+" successfully updated")
}

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
	g.GET("/user/:id", readUser)
	g.POST("/user", addUser)
	g.DELETE("/user/:id", deleteUser)
	g.PUT("/user/:id", updateUser)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
