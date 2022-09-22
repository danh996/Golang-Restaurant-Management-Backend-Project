package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/danh996/go-destiny/book"
	"github.com/danh996/go-destiny/database"
	"github.com/gofiber/fiber"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	_ "github.com/lib/pq"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"
	_, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connect to Postgres Database Success")
	// database.DBConn.AutoMigrate(&book.Book{})
	// fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()

	setupRoutes(app)
	app.Listen(3000)

	defer database.DBConn.Close()
}
