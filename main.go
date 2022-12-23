package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"

	"net/http"

	"RestAPI-go/model"
)

// load .env file
func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

// connect to the database and return it as an object
func dbConn() (db *sql.DB) {
	// pass the db credentials into variables
	host := goDotEnvVariable("DBHOST")
	port := goDotEnvVariable("DBPORT")
	dbUser := goDotEnvVariable("DBUSER")
	dbPass := goDotEnvVariable("DBPASS")
	dbname := goDotEnvVariable("DBNAME")
	// create a connection string
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, dbUser, dbPass, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}

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
	return c.JSONPretty(http.StatusOK, user, "  ")
}

// print Connected! if connection is successful
func main() {
	db := dbConn()
	var err error
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
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
