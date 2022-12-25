package posts

import (
	"RestAPI-go/models"
	"RestAPI-go/storage"
	"github.com/labstack/echo/v4"

	"log"
	"net/http"
)

func CreateUser(c echo.Context) error {
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

func ReadUser(c echo.Context) error {
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

func UpdateUser(c echo.Context) error {
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

func DeleteUser(c echo.Context) error {
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

func Login(c echo.Context) error {
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
