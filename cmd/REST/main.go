package main

import (
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"progress-tracker/internal/config"
	"progress-tracker/internal/handlers"
	"progress-tracker/internal/middlewares"

	gohandlers "github.com/gorilla/handlers"
	"progress-tracker/internal/services"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config.Configurate()

	dbURL := viper.GetString("database.url")

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {

		log.Fatal(err)
	}

	jobService := services.NewJobService(db)
	progressService := services.NewProgressService()
	progressService.StartQueueWorker()

	jobHandler := handlers.NewJobHandler(*jobService, *progressService)

	// Создаём новый router
	r := mux.NewRouter()
	headersOk := gohandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := gohandlers.AllowedOrigins([]string{"*"})
	methodsOk := gohandlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	r.Use(middlewares.CorsMiddleware)
	r.Use(middlewares.LoggingMiddleware)
	r.Use(middlewares.AuthMiddleware)

	// Регистрируем роуты
	r.HandleFunc("/jobs/create", jobHandler.CreateJob).Methods("POST")
	r.HandleFunc("/jobs/start", jobHandler.StartJob).Methods("POST")
	r.HandleFunc("/jobs/{id}", jobHandler.GetJobByID).Methods("GET")
	r.HandleFunc("/jobs", jobHandler.GetAllJob).Methods("GET")
	r.HandleFunc("/jobs?job_id", jobHandler.GetAllJob).Methods("GET")
	r.HandleFunc("/progress/{job_id}", jobHandler.GetProgress).Methods("GET")
	r.HandleFunc("/progress", jobHandler.SetJobProgress).Methods("POST")

	// Запускаем сервер
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", gohandlers.CORS(originsOk, headersOk, methodsOk)(r)); err != nil {
		log.Fatal(err)
	}
}
