package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"progress-tracker/internal/handlers"

	"progress-tracker/internal/services"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	dbURL := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	jobService := services.NewJobService(db)
	jobHandler := handlers.NewJobHandler(*jobService)

	// Создаём новый router
	r := mux.NewRouter()

	// Регистрируем роуты
	r.HandleFunc("/jobs", jobHandler.CreateJob).Methods("POST")
	r.HandleFunc("/jobs/{id}", jobHandler.GetJobByID).Methods("GET")

	// Запускаем сервер
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
