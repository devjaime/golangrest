package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/devjaime/golangrest/database"
	"github.com/devjaime/golangrest/product"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/joho/godotenv"
)

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

var (
	server       = goDotEnvVariable("server")
	port, errint = strconv.Atoi(goDotEnvVariable("port"))
	user         = goDotEnvVariable("user")
	password     = goDotEnvVariable("password")
	databaseBD   = goDotEnvVariable("databaseBD")
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/products", product.GetProducts)
	app.Get("/api/v1/product/:item_number", product.GetProduct)
	app.Post("/api/v1/product", product.NewProduct)

}

func initDatabase() {
	var err error
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		server, user, password, port, databaseBD)
	database.DBConn, err = gorm.Open("mssql", connectionString)

	if err != nil {
		panic("Failed to connect to database")
	}
	gorm.DefaultCallback.Create().Remove("mssql:set_identity_insert")

	fmt.Println("Database connection successfully opened")
	database.DBConn.AutoMigrate(&product.Product{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()

	initDatabase()
	defer database.DBConn.Close()
	setupRoutes(app)

	app.Listen(4000)
}
