package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/template/html/v2"
	"github.com/skip2/go-qrcode"
	"os"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	app.Get("/api/generate", func(c *fiber.Ctx) error {
		queries := c.Queries()

		if queries["data"] == "" {
			return c.Status(400).SendString("Bad Request. Data is missing.")
		}

		var png []byte
		png, err := qrcode.Encode(queries["data"], qrcode.Medium, c.QueryInt("size", 256))

		if err != nil {
			log.Info("Failed generating QR-Code.")

			return c.Status(500).SendString("Failed generating the QR Code.")
		}

		log.Info("Successfully generated QR-Code.")

		return c.Send(png)
	})

	err := app.Listen("0.0.0.0:" + os.Getenv("PORT"))

	if err != nil {
		log.Fatal("Something went very wrong.")
	}
}
