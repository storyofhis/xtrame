package config

import (
	"fmt"
	"log"
	"os"

	"github.com/storyofhis/xtrame/httpserver/repository/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	host := os.Getenv("PGHOST")
	port := os.Getenv("PGPORT")
	user := os.Getenv("PGUSER")
	pass := os.Getenv("PGPASSWORD")
	pgdb := os.Getenv("PGDATABASE")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, pgdb,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println(err)
	}

	// don't forget to automigrate
	db.Debug().AutoMigrate(models.Users{}, models.Category{}, models.Tickets{})

	return db, nil
}
