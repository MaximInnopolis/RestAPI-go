package posts

import (
	"RestAPI-go/models"
	"RestAPI-go/storage"
	"github.com/labstack/echo/v4"

	"log"
	"net/http"
)

func BasicAccess(key string, c echo.Context) (bool, error) {
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

	if user.Status != "ACTIVE" {
		log.Println("Unable to read")
		return false, c.String(http.StatusForbidden, "Forbidden")
	}

	return true, nil
}

func SupremeAccess(key string, c echo.Context) (bool, error) {
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

	if user.Status != "ACTIVE" && user.Role != "ADMIN" {
		log.Println("Unable to process operation")
		return false, c.String(http.StatusForbidden, "Forbidden")
	}

	return true, nil
}
