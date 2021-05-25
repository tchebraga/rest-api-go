package database

import (
	"github.com/jinzhu/gorm"
	"github.com/tchebraga/rest-api-go/internal/comment"
)

func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&comment.Comment{}); result.Error != nil {
		return result.Error
	}
	return nil
}
