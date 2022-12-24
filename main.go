package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"

	"net/http"

	"RestAPI-go/models"
	"RestAPI-go/storage"
)

func mainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, admin! You are on the admin main page")
}

func addUser(c echo.Context) error {
	user := models.User{}

	err := c.Bind(&user)
	if err != nil {
		log.Printf("Failed processing addUser request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("User added: %#v", user)

	storage.GetDBInstance().Create(&user)

	return c.JSONPretty(http.StatusOK, user, "  ")
}

func readUser(c echo.Context) error {
	user := models.User{}
	id := c.Param("id")

	storage.GetDBInstance().First(&user, id)

	log.Printf("user %v", user)

	if user.ID == 0 {
		return c.String(http.StatusNotFound, "Not found")
	}

	return c.JSONPretty(http.StatusOK, user, "  ")
}

func deleteUser(c echo.Context) error {
	user := models.User{}
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
	user := models.User{}
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

func login(c echo.Context) error {
	db := storage.GetDBInstance()

	type RequestBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var body RequestBody

	if err := c.Bind(&body); err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	user := models.User{}

	if err := db.Where("login = ?", body.Username).First(&user).Error; err != nil {
		return c.NoContent(http.StatusForbidden)
	}

	if user.Password != body.Password {
		return c.NoContent(http.StatusForbidden)
	}

	key := models.GenerateKey(user.ID)
	db.Create(&key)
	return c.JSONPretty(http.StatusOK, key, "  ")
}

func main() {

	// Echo instance
	e := echo.New()

	//g := e.Group("/admin")

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return key == "valid-key", nil
	}))

	// Routes

	e.POST("/login", login)

	e.GET("/main", mainAdmin)

	e.GET("/user/:id", readUser)
	e.POST("/user", addUser)
	e.DELETE("/user/:id", deleteUser)
	e.PUT("/user/:id", updateUser)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
