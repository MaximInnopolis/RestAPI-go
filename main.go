package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"log"
	"net/http"

	"RestAPI-go/models"
	"RestAPI-go/storage"
)

func createUser(c echo.Context) error {
	user := models.User{}
	db := storage.GetDBInstance()

	err := c.Bind(&user)
	if err != nil {
		log.Printf("Failed processing createUser request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("User added: %v", user)

	db.Create(&user)

	return c.JSONPretty(http.StatusOK, user, "  ")
}

func readUser(c echo.Context) error {
	user := models.User{}
	id := c.Param("id")
	db := storage.GetDBInstance()

	db.First(&user, id)

	log.Printf("User with id %v is readable: %v ", id, user)

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

	log.Printf("User with id %v deleted: %v ", id, user)

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

	log.Printf("User with id %v saved: %v ", id, user)

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

	log.Printf("Key created and saved: %v ", key)

	return c.JSONPretty(http.StatusOK, key, "  ")
}

func supremeAccess(key string, c echo.Context) (bool, error) {
	k := models.Key{}
	db := storage.GetDBInstance()

	db.Where("key = ?", key).Last(&k)

	log.Printf("key %v", k)

	if k.Key == "" {

		log.Println("Key not found")

		return false, c.String(http.StatusNotFound, "Not found")
	}

	user := models.User{}
	db.First(&user, k.UserID)

	log.Printf("user %v", user)

	if user.Status == "BLOCKED" || user.Role == "READER" {
		log.Println("Unable to process operation")
		return false, c.String(http.StatusForbidden, "Forbidden")
	}

	return true, nil
}

func main() {

	storage.DBConn()

	// Echo instance
	e := echo.New()

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes

	e.POST("/login", login)

	e.GET("/user/:id", readUser, middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		k := models.Key{}
		db := storage.GetDBInstance()

		db.Where("key = ?", key).Last(&k)

		log.Printf("key %v", k)

		if k.Key == "" {

			log.Println("Key not found")

			return false, c.String(http.StatusNotFound, "Not found")
		}

		user := models.User{}
		db.First(&user, k.UserID)

		log.Printf("user %v", user)

		if user.Status == "BLOCKED" {
			log.Println("Unable to read")
			return false, c.String(http.StatusForbidden, "Forbidden")
		}

		return true, nil
	}))

	e.POST("/user", createUser, middleware.KeyAuth(supremeAccess))
	e.DELETE("/user/:id", deleteUser, middleware.KeyAuth(supremeAccess))
	e.PUT("/user/:id", updateUser, middleware.KeyAuth(supremeAccess))

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
