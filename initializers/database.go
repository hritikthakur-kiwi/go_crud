package initializers

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := "host=localhost user=tutorial dbname=golang password=qwert123 port=5432"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("failed to connect to the database:", err)
	} else {
		log.Println("database connected:", dsn)
	}

}
