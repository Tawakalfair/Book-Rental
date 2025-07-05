package main

import (
	"log"

	"github.com/Tawakalfair/Book-Rental/book"
	"github.com/Tawakalfair/Book-Rental/database"
	"github.com/Tawakalfair/Book-Rental/router"
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

	// Serve static files from the "public" directory
	app.Static("/public", "./public")

	// Setup the routes
	router.SetupRoutes(app)

	app.Listen(":3000")
}
