package main

import (
	"github.com/gofiber/fiber/v2"
	controller "github.com/abiyyu03/database-book-api/controllers"
	config "github.com/abiyyu03/database-book-api/config"
)

func main(){
	app := fiber.New()

	config.Connect()
	bookController := new(controller.BookController)
	book := app.Group("/book")

	book.Get("/", bookController.GetAll)
	book.Get("/:id", bookController.GetById) 
	book.Delete("/:id", bookController.Delete)
	book.Post("/", bookController.CreateData)
	book.Put("/:id", bookController.Update)

	app.Listen(":8080")

}