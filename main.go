package main

import (
	"github.com/gofiber/fiber/v2"
	controller "github.com/abiyyu03/database-book-api/controllers"
	config "github.com/abiyyu03/database-book-api/config"
	"github.com/spf13/viper"
)

func main(){
	app := fiber.New()

	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	config.Connect()
	bookController := new(controller.BookController)
	book := app.Group("/book")

	book.Get("/", bookController.GetAll)
	book.Get("/:id", bookController.GetById) 
	book.Delete("/:id", bookController.Delete)
	book.Post("/", bookController.CreateData)
	book.Put("/:id", bookController.Update)

	app.Listen(viper.Get("PORT").(string))

}