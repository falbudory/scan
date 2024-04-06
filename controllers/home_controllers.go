package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func GetHome(c *fiber.Ctx) error {
	return c.Render("pages/home/index", fiber.Map{
		"Ctx": c,
	}, "layouts/main")
}
