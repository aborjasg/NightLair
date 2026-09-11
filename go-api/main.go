package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"go-api/internal/handlers"
	"go-api/internal/repositories"
	"go-api/internal/services"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		requiredEnv("DB_USER"),
		requiredEnv("DB_PASSWORD"),
		requiredEnv("DB_HOST"),
		envOrDefault("DB_PORT", "3306"),
		requiredEnv("DB_NAME"),
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected!")

	repo := &repositories.ProspectRepository{DB: db}
	service := &services.ProspectService{Repo: repo}
	handler := &handlers.ProspectHandler{Service: service}

	http.HandleFunc("/api/hello", handler.HelloAPI)
	http.HandleFunc("/api/prospects", handler.GetRowCount)

	log.Println("API running on :8080")
	http.ListenAndServe(":8080", nil)
}

func requiredEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("required environment variable %s is not set", key)
	}
	return value
}

func envOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
