package controllers

import (
	"log"

	"bitbucket.com/Divyanshu_B/go_fiber_rest/models"
	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
)

func Pong(c *fiber.Ctx) error {

	var book *models.Book
	newBook := &models.Book{
		Name:  "3hornies in a boat",
		Pages: 3,
	}

	err := mgm.Coll(book).Create(newBook)
	if err != nil {
		log.Fatal("Creation unsuccessful")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Hello world; from fiber port 8000",
	})
}
