package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/jyablonski/go_praq/internal/database"
	_ "github.com/lib/pq" // include this code even if we dont use it. need the driver
)

func main() {
	godotenv.Load("../.env")

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("port is not found in env vars")
	}
	fmt.Println(port)

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("db_url not found in env vars")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("can't connect to database:", err)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server Starting on Port %v", port)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
