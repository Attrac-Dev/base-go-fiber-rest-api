package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/skyler-saville/base-api-fiber/config"
	model "github.com/skyler-saville/base-api-fiber/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Declare the variable for the DB
var DB *gorm.DB

// Connect to the database
func ConnectDB() {
	var err error 
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Idiot")
	}

	// Connection URL to connect to the Postgresql Database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	// Connect to the DB and initialize the DB variable
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("Failed to connect to database")
	}
	if port != 5432 {
		// connected to the docker instance of Postgresql
		fmt.Println("Connection to Docker Container DB successful")
	} else {
		// connected to the default port (which is the local install of Postgresql)
		fmt.Println("Connection to local DB successful")
	}
	

	// Migrate the database
	DB.AutoMigrate(&model.Note{})
	fmt.Println("Database has been Migrated")
}