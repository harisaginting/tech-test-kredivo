package db

import (
	"log"
	"gorm.io/gorm"
	"github.com/harisaginting/tech-test-kredivo/db/table"
)

func Migration(db *gorm.DB) {
	log.Println("migration db ",db.Migrator().CurrentDatabase())
	table.MigrateUser(db)
	log.Println("migration success")
}
