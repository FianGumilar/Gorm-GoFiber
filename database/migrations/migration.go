package migrations

import (
	"fmt"
	"gorm-fiber/database"
	"gorm-fiber/models"
)

func Migration() {
	err := database.DB.AutoMigrate(
		&models.User{},
		&models.Locker{},
		&models.Post{},
		&models.Tag{},
	)

	if err != nil {
		fmt.Println("Can't migrate...")
	}

	fmt.Println("Migrating...")
}
