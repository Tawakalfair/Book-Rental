package main

import (
	"example-app/book"
	"example-app/database"
	"example-app/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Initialize the template engine
	engine := html.New("./views", ".html")

	// Initialize the database
	database.ConnectDB()
	log.Println("Running Migrations")
	// Run the migration here
	database.DB.AutoMigrate(&book.Book{})

	// Create a new Fiber app
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Setup the routes
	router.SetupRoutes(app)

	app.Listen(":3000")
}
