package main

import (
	"fmt"
	"log"

	"github.com/RamiroCyber/gateway-go/internal/config"
	"github.com/RamiroCyber/gateway-go/internal/repository"
	"github.com/RamiroCyber/gateway-go/internal/service"
	"github.com/RamiroCyber/gateway-go/internal/web/server"
	_ "github.com/lib/pq"
)

func main() {
	envConfig := config.NewEnvConfig()
	if err := envConfig.Load(); err != nil {
		log.Fatal("Error loading config:", err)
	}

	dbConfig := config.NewSQLDatabase()
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		envConfig.Get("DB_HOST"),
		envConfig.Get("DB_PORT"),
		envConfig.Get("DB_USER"),
		envConfig.Get("DB_PASSWORD"),
		envConfig.Get("DB_NAME"),
		envConfig.Get("DB_SSLMODE"),
	)

	db, err := dbConfig.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	defer dbConfig.Close()

	repo := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(repo)

	server := server.NewServer(accountService, envConfig.Get("PORT"))
	if err := server.Start(); err != nil {
		log.Fatal("Error starting server:", err)
	}
	log.Println("Server started on port", envConfig.Get("PORT"))
}
