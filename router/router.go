package router

import (
	"example-app/book"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	// Route for Web interface
	app.Get("/", book.RenderBooksPage)

	// grouping routes under /api
	api := app.Group("/api")
	// Book Routes
	api.Get("/books", book.GetBooks)
	api.Get("/books/:id", book.GetBook)
	api.Post("/books", book.CreateBook)
	api.Put("/books/:id", book.UpdateBook)
	api.Delete("/books/:id", book.DeleteBook)

}
