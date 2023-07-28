package controllers

import (
	// "strconv"
	"github.com/abiyyu03/database-book-api/models"
	config "github.com/abiyyu03/database-book-api/config"
	"github.com/gofiber/fiber/v2"
)

type BookController struct{}

func (b *BookController) GetAll(c *fiber.Ctx) error {
	var books []models.Book

	config.Database.Find(&books)
	return c.Status(200).JSON(books)
}

func (b *BookController) GetById(c *fiber.Ctx) error {
	return c.SendString("get Books number "+c.Params("id"))
}

func (b *BookController) Delete(c *fiber.Ctx) error {
	return c.SendString("Delete Books number "+c.Params("id"))
}

func (b *BookController) CreateData(c *fiber.Ctx) error {
	return c.SendString("Add Books number")
}

func (b *BookController) Update(c *fiber.Ctx) error {
	return c.SendString("Add Books number")
}