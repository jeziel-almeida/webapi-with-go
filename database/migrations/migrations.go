package migrations

import (
	"gorm.io/gorm"
	"github.com/jeziel-almeida/webapi-with-go/models"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}