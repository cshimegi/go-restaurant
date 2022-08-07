package dao

import (
	"log"

	"users/shared/domain"

	"gorm.io/gorm"
)

// MigrateSchema migrates schema
func MigrateSchema(db gorm.DB) {
	migrator := db.Migrator()
	createTableIfNotExist(migrator, domain.User{})
}

func createTableIfNotExist(migrator gorm.Migrator, table interface{}) {
	exist := migrator.HasTable(&table)
	if !exist {
		err := migrator.CreateTable(&table)
		if err != nil {
			log.Panicln(err)
		}
	}
}
