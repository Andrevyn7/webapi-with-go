package migrations

import (
	"github.com/Andrevyn7/webapi-with-go.git/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
