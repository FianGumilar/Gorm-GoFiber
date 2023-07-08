package controllers

import (
	"gorm-fiber/database"
	"gorm-fiber/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllTag(c *fiber.Ctx) error {
	var tags []*models.TagResponseWithPost

	database.DB.Preload("Posts").Find(&tags)

	return c.Status(200).JSON(fiber.Map{
		"message": "success get all tags",
		"tags":    tags,
	})
}

func CreateTag(c *fiber.Ctx) error {
	tag := new(models.Tag)

	if err := c.BodyParser(tag); err != nil {
		return c.Status(503).JSON(err.Error())
	}

	if tag.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Name is required",
		})
	}

	database.DB.Debug().Create(&tag)

	return c.Status(200).JSON(fiber.Map{
		"message": "sucess created tag",
	})
}
