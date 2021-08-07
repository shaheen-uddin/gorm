package book

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shaheen-uddin/gorm/database"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"name"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) error {
	db := database.DbCon
	var books []Book
	db.Find(&books)
	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	db := database.DbCon
	id := c.Params("id")
	var book Book
	db.Find(&book, id)
	return c.JSON(book)
}

func NewBook(c *fiber.Ctx) error {
	db := database.DbCon
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Create(&book)
	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DbCon

	var book Book
	db.First(&book, id)
	if book.Title == "" {
		return c.Status(500).SendString("No book found with ID")
	}
	db.Delete(&book)
	return c.SendString("Book Successfully deleted")
}
