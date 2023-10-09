package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pusher/pusher-http-go/v5"
)

func main() {

	app := fiber.New()

	app.Use(cors.New())

	pusherClient := pusher.Client{
		AppID:   "1683338",
		Key:     "5bfed9f71dd24705a3e1",
		Secret:  "4b224cba6339d25d5549",
		Cluster: "eu",
		Secure:  true,
	}

	app.Post("/api/messages", func(c *fiber.Ctx) error {
		var data map[string]string
		err := c.BodyParser(&data)
		if err != nil {
			return err
		}

		pusherClient.Trigger("chat", "message", data)

		return c.JSON([]string{})

	})

	app.Listen("localhost:3001")

}
