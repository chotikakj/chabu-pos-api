package main

import (
	"log"
	GoCache "pos-api/cache/go-cache"
	"pos-api/config"
	"pos-api/database"

	"pos-api/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	database.InitialMongoDB()
	GoCache.InitCache()
	app.Use(cors.New(config.CorsConfigDefault))
	router.SetupRoutes(app)
	err := app.Listen(":8888")
	if err != nil {
		log.Fatal(err)
	}
}
