package controllers

import (
	"gorm-fiber/database"
	"gorm-fiber/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func UserGetAll(c *fiber.Ctx) error {
	var users []*models.User

	database.DB.Preload("Locker").Preload("Posts").Find(&users)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "All users",
		"users":   users,
	})
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	// Parse Body Request ke Object Struct
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Can't handle request",
		})
	}

	if user.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name is required",
		})
	}

	database.DB.Create(&user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Created User successfully",
		"user":    user,
	})
}

func UserGetById(c *fiber.Ctx) error {
	users := []*models.User{}

	database.DB.First(&users, c.Params("id"))

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"users": users,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	id, _ := strconv.Atoi(c.Params("id"))

	database.DB.Model(&models.User{}).Where("id = ?", id).Update("name", user.Name)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success to update user data",
		"user":    user,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	book := new(models.User)

	id, _ := strconv.Atoi(c.Params("id"))

	database.DB.Where("id = ?", id).Delete(&book)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": `success delete book with id: ${id}`,
	})
}
