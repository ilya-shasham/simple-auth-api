package dbutils

import (
	"auth-api/globals"
	"auth-api/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDatabase() {
	var err error
	globals.Db, err = gorm.Open(postgres.Open(globals.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	globals.Db.AutoMigrate(models.All...)
}
