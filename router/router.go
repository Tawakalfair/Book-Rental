package router

import (
	"github.com/Tawakalfair/Book-Rental/auth"
	"github.com/Tawakalfair/Book-Rental/book"
	"github.com/Tawakalfair/Book-Rental/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	// Route for Web interface
	app.Get("/", book.RenderBooksPage)

	app.Post("/api/register", auth.Register)
	app.Post("/api/login", auth.Login)
	// grouping routes under /api
	api := app.Group("/api", middleware.Protected)
	// Book Routes
	api.Get("/books", book.GetBooks)
	api.Get("/books/:id", book.GetBook)
	api.Post("/books", book.CreateBook)
	api.Put("/books/:id", book.UpdateBook)
	api.Delete("/books/:id", book.DeleteBook)

}
