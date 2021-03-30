package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/el-Mike/gochat/auth"
	"github.com/el-Mike/gochat/core/control"
	"github.com/el-Mike/gochat/models"
	"github.com/el-Mike/gochat/persist"
	"github.com/el-Mike/gochat/services"
	"github.com/joho/godotenv"
)

func main() {
	_, b, _, _ := runtime.Caller(0)
	rootPath := filepath.Join(filepath.Dir(b), "../../")

	if err := godotenv.Load(rootPath + "/.env"); err != nil {
		log.Fatal(err)
	}

	pgUser := os.Getenv("POSTGRES_USER")
	pgPassword := os.Getenv("POSTGRES_PASSWORD")
	pgDBname := os.Getenv("POSTGRES_DB")
	pgPort := os.Getenv("POSTGRES_PORT")
	pgHost := "localhost"

	adminPassword := os.Getenv("GOCHAT_ADMIN_PASSWORD")
	adminEmail := os.Getenv("GOCHAT_ADMIN_EMAIL")

	_, err := persist.InitDatabase(pgUser, pgPassword, pgDBname, pgHost, pgPort)

	us := services.NewUserService()
	am := auth.NewAuthManager()

	if err != nil {
		log.Fatal("Database connection failed")
	}

	hashedPassword, err := am.HashAndSalt([]byte(adminPassword))

	if err != nil {
		log.Fatal(err)
	}

	adminUser := &models.UserModel{
		Password:  hashedPassword,
		Email:     adminEmail,
		FirstName: "John",
		LastName:  "Doe",
		Role:      control.SuperAdminRole,
	}

	err = us.SaveUser(adminUser)

	if err != nil {
		log.Fatal(err)
	}
}
