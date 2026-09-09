package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"go-api/internal/handlers"
	"go-api/internal/repositories"
	"go-api/internal/services"

	_ "github.com/go-sql-driver/mysql"
)

type Message struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func helloAPI(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	response := Message{
		Status:  "success",
		Message: "Hello from your first Go API!",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// http.HandleFunc("/api/hello", helloAPI)
	// http.ListenAndServe(":8080", nil)

	db, err := sql.Open("mysql", "asjrf517_ozorasoft_admin:Passw0rd1978@tcp(ozorasoft.ca:3306)/asjrf517_ozorasoft_marketing")
	if err != nil {
		log.Fatal(err)
	}

	repo := &repositories.ProspectRepository{DB: db}
	service := &services.ProspectService{Repo: repo}
	handler := &handlers.ProspectHandler{Service: service}

	http.HandleFunc("/api/prospects", handler.GetRowCount)

	log.Println("API running on :8080")
	http.ListenAndServe(":8080", nil)
}
