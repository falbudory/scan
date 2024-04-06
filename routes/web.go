package routes

import (
	"github.com/gofiber/fiber/v2"
	"serverWeb/controllers"
)

func RouteInit(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		if c.Method() == "POST" {
			method := c.FormValue("_method")

			switch method {
			case "DELETE":
				c.Method("DELETE")
			case "PUT":
				c.Method("PUT")
			}
		}
		return c.Next()
	})

	app.Get("/", controllers.GetLogin)
	app.Get("/login", controllers.GetLogin)
	app.Get("/logout", controllers.GetLogout)
	app.Post("/login", controllers.PostLogin)

	home := app.Group("/home", IsAuthenticated, CheckSession, CheckPerHome)
	home.Get("", controllers.GetHome)

}
