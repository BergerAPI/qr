package main

import (
	"context"
	"github.com/axiomhq/axiom-go/axiom"
	"github.com/axiomhq/axiom-go/axiom/ingest"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/template/html/v2"
	_ "github.com/joho/godotenv/autoload"
	"github.com/skip2/go-qrcode"
	"os"
	"time"
)

func main() {
	ctx := context.Background()
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Axiom for Data-analysis
	client, err := axiom.NewClient(
		axiom.SetPersonalTokenConfig(os.Getenv("AXIOM_TOKEN"), os.Getenv("AXIOM_ORG_ID")),
	)

	if err != nil {
		log.Fatal("Something went very wrong with Axiom, please check the environment.")
	}

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

		if _, err = client.IngestEvents(ctx, "qr", []axiom.Event{
			{ingest.TimestampField: time.Now(), "data": queries["data"]},
		}); err != nil {
			log.Fatal(err)
		}

		return c.Send(png)
	})

	err = app.Listen("0.0.0.0:" + os.Getenv("PORT"))

	if err != nil {
		log.Fatal("Something went very wrong.")
	}
}
