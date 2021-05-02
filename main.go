package main

import (
	"log"

	"bitbucket.com/Divyanshu_B/go_fiber_rest/config"
	"bitbucket.com/Divyanshu_B/go_fiber_rest/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv" // new
)

func setUpRoutes(app *fiber.App) {
	api_root := app.Group("/api")
	routes.RootRoute(api_root)

}

func main() {

	config.InitDB()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()
	app.Use(logger.New())

	setUpRoutes(app)

	err = app.Listen(":8000")
	if err != nil {
		panic(err)
	}
}
