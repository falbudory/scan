package controllers

import (
	"github.com/gofiber/fiber/v2"
	"serverWeb/models"
)

func GetCodeUser(c *fiber.Ctx) string {
	return GetSessionUser(c).CodeUser
}

func IsChecked(currentID int, checkedID []models.RolePermission) bool {
	found := false

	for _, check := range checkedID {
		if check.PermissionID == currentID {
			found = true
			break
		}
	}

	return found
}

func IsSelected(currentID, selectedID int) bool {
	return currentID == selectedID
}
