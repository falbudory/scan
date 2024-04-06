package routes

import (
	"github.com/gofiber/fiber/v2"
	"serverWeb/controllers"
)

func CheckAuthentication(c *fiber.Ctx) bool {
	sess, _ := controllers.SessAuth.Get(c)
	username := sess.Get("login_success")
	if username != nil {
		return true
	}
	return false
}

func IsAuthenticated(c *fiber.Ctx) error {
	if !CheckAuthentication(c) {
		return c.Redirect("/login")
	}
	return c.Next()
}

func CheckSession(c *fiber.Ctx) error {
	userLogin := controllers.GetSessionUser(c)
	sess, _ := controllers.SessAuth.Get(c)

	if userLogin.Session != sess.Get("sessionId") {
		return c.Redirect("/logout")
	}
	return c.Next()
}

func CheckPerHome(c *fiber.Ctx) error {
	if !controllers.CheckPermission("home", c) {
		return c.Redirect("/errors/403")
	}
	return c.Next()
}
