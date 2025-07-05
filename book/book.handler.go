package book

import (
	"example-app/database"
	"log"

	"github.com/gofiber/fiber/v2"
)

// get all books from DB
func GetBooks(c *fiber.Ctx) error {
	var books []Book
	if err := database.DB.Find(&books).Error; err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(books)
}

// get one book
func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book Book
	if err := database.DB.First(&book, id).Error; err != nil {
		return c.Status(404).SendString("Book not found")
	}
	return c.JSON(book)
}

// Create new book to db
func CreateBook(c *fiber.Ctx) error {
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	if err := database.DB.Create(&book).Error; err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(book)
}

// update book
func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book Book

	if err := database.DB.First(&book, id).Error; err != nil {
		return c.Status(404).SendString("Book Not found")
	}

	var updateData Book
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	database.DB.Model(&book).Updates(updateData)

	return c.JSON(book)

}

// delete book
func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book Book

	result := database.DB.Delete(&book, id)

	if result.RowsAffected == 0 {
		return c.Status(404).SendString("Book not found")
	}

	return c.SendStatus(fiber.StatusNoContent)

}

func RenderBooksPage(c *fiber.Ctx) error {
	var books []Book
	if err := database.DB.Find(&books).Error; err != nil {
		// Log the actual error for debugging purposes
		log.Printf("Database error: %v", err)

		// Render the error page for the user
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{
			"StatusCode": fiber.StatusInternalServerError,
			"Message":    "We couldn't retrieve the book list at this time. Please try again later.",
		})
	}

	return c.Render("index", fiber.Map{
		"Books": books,
	})
}
