package controllers

import (
	"gorm-fiber/database"
	"gorm-fiber/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func LockerGetAll(c *fiber.Ctx) error {
	var lockers []models.Locker

	database.DB.Preload("User").Find(&lockers)

	return c.JSON(fiber.Map{
		"lockers": lockers,
	})
}

func CreateLocker(c *fiber.Ctx) error {
	locker := new(models.Locker)

	//Parse Body Request ke Object Struct
	if err := c.BodyParser(locker); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"err": "Cant handle this request",
		})
	}

	if locker.Code == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"err": "Code is empty",
		})
	}

	if locker.UserId == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"err": "UserID is required",
		})
	}

	database.DB.Create(&locker)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "successfully created locker",
		"locker":  locker,
	})
}

func LockerGetById(c *fiber.Ctx) error {
	lockers := []*models.Locker{}

	if err := database.DB.First(&lockers, c.Params("id")); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"lockers": lockers,
	})
}

func UpdateLocker(c *fiber.Ctx) error {
	locker := new(models.Locker)

	// Parsing body request data ke object struct
	if err := c.BodyParser(locker); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	id, _ := strconv.Atoi(c.Params("id"))

	database.DB.Model(&models.Locker{}).Where("id = ?", id).Update("code", locker.Code)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success to update locker",
		"locker":  locker,
	})
}

func DeleteLocker(c *fiber.Ctx) error {
	locker := new(models.Locker)

	id, _ := strconv.Atoi(c.Params("id"))

	database.DB.Where("id = ?", id).Delete(&locker)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success to delete locker",
	})
}
