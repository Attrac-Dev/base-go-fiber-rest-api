package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/skyler-saville/base-api-fiber/database"
)


func main() {
    app := fiber.New()

    // attempt database connection
    database.ConnectDB()

    app.Get("/", func(c *fiber.Ctx) error {
        err := c.SendString("API is up!!!")
		return err
    })

	// listen on port:3000
    app.Listen(":3000")
}