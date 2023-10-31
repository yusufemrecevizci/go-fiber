package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Could not connect to the database! \n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to the database!")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Migrating the schema...")

	Database = DbInstance{Db: db}

}
