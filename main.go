package main

import (
	"log"
	"os"

	"github.com/el-Mike/gochat/core/control"
	"github.com/el-Mike/gochat/core/control/rbac"
	"github.com/el-Mike/gochat/routing"

	"github.com/el-Mike/gochat/persist"

	"github.com/joho/godotenv"
)

func main() {
	// When app is run in debug mode, we want to use postgres host defined
	// in task's configuration, as debug mode is run outside of docker-compose network.
	debugDbHost := os.Getenv("POSTGRES_HOST")

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	pgUser := os.Getenv("POSTGRES_USER")
	pgPassword := os.Getenv("POSTGRES_PASSWORD")
	pgDBname := os.Getenv("POSTGRES_DB")
	pgPort := os.Getenv("POSTGRES_PORT")

	pgHost := debugDbHost

	if pgHost == "" {
		pgHost = os.Getenv("POSTGRES_HOST")
	}

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	_, err := persist.InitDatabase(pgUser, pgPassword, pgDBname, pgHost, pgPort)
	_ = persist.InitRedisClient(redisHost, redisPort, redisPassword)

	if err != nil {
		log.Fatal("Database connection failed")
	}

	err = persist.AutoMigrate()

	if err != nil {
		log.Fatal(err)
	}

	_, err = rbac.InitAccessManager(control.Policy)

	if err != nil {
		log.Fatal("RBAC initialization failed")
	}

	routing.InitRouting()
}
