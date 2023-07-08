package controllers

import (
	"gorm-fiber/database"
	"gorm-fiber/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllPost(c *fiber.Ctx) error {
	posts := []*models.Post{}

	database.DB.Preload("User").Find(&posts)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success get all posts",
		"posts":   posts,
	})
}

func CreatePost(c *fiber.Ctx) error {
	posts := new(models.Post)

	// Parsing Body request ke object struct
	if err := c.BodyParser(posts); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if posts.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Title is required",
		})
	}

	if posts.Body == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Body is required",
		})
	}

	if len(posts.TagID) == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Tag ID id required",
		})
	}

	database.DB.Create(&posts)

	if len(posts.TagID) > 0 {
		for _, tagID := range posts.TagID {
			postTag := new(models.PostTag)
			postTag.PostID = posts.ID
			postTag.TagID = tagID
			database.DB.Create(&postTag)
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success created post",
		"post":    posts,
	})
}
