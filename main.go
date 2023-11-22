package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello World")

	//load the .env file
	godotenv.Load(".env")

	posrtString := os.Getenv("PORT")
	if posrtString == "" {
		log.Fatal("PORT NOT FOUND IN THE .env")
	}
	// fmt.Printf("PORT: %v", posrtString)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerError)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + posrtString,
	}

	log.Printf("Server starting on port %v", posrtString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
