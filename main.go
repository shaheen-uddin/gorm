package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/shaheen-uddin/gorm/book"
	"github.com/shaheen-uddin/gorm/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/book", book.GetBooks)
	app.Get("/api/book/:id", book.GetBook)
	app.Post("/app/book", book.NewBook)
	app.Delete("app/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	database.DbCon, err = gorm.Open(sqlite.Open("book.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")

}

func main() {
	app := fiber.New()
	app.Use(cors.New)

	initDatabase()
	defer database.DbCon.Close()
	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
