package persist

import (
	"errors"
	"fmt"
	"log"

	"github.com/el-Mike/gochat/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB - singleton variable for storing DB connection
var DB *gorm.DB

// Init initializes database driver
func InitDatabase(username, password, dbname, host, port string) (*gorm.DB, error) {
	if DB != nil {
		return DB, nil
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, username, password, dbname)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	DB = conn

	log.Println("Database connection estabilished")

	return DB, nil
}

// AutoMigrate - fires the migrations for DB schemas
func AutoMigrate() error {
	if DB == nil {
		return errors.New("Database has not been initialized")
	}

	err := DB.AutoMigrate(
		&models.UserModel{},
	)

	if err != nil {
		return err
	}

	return nil
}
