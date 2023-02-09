package orm

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	database := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	port := os.Getenv("DB_PORT")

	connection := fmt.Sprintf("host=%s database=%s user=%s password=%s port=%s", host, database, user, password, port)

	var gormDb, err = gorm.Open(postgres.Open(connection), &gorm.Config{})
		if err != nil {
			panic("db connection failed")
		}

		return gormDb, nil

}		

