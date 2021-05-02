package routes

import (
	"bitbucket.com/Divyanshu_B/go_fiber_rest/controllers"
	"github.com/gofiber/fiber/v2"
)

func RootRoute(route fiber.Router) {
	route.Get("/ping", controllers.Pong)
}

func ProductRoutes(route fiber.Router) {
	route.Post("/", controllers.AddProduct)
}
