package router

import (
	"github.com/Tawakalfair/Book-Rental/auth"
	"github.com/Tawakalfair/Book-Rental/book"
	"github.com/Tawakalfair/Book-Rental/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	// Public routes for authentication pages and actions
	app.Get("/register", auth.ShowRegisterPage)
	app.Post("/register", auth.RegisterUser)
	app.Get("/login", auth.ShowLoginPage)
	app.Post("/login", auth.LoginUser)
	app.Post("/logout", auth.LogoutUser)

	// Protected routes
	app.Get("/dashboard", middleware.Protected, auth.ShowDashboard)
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
