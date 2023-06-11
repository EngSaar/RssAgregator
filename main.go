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

	portString := LoadEnviroment()
	serverStarter(portString)

}

func serverStarter(portString string) {
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
	v1Router.Get("/produtos", handlerGet)
	v1Router.Post("/produtos", handlerPost)
	v1Router.Put("/err", handlerErro)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Addr:    ":" + portString,
		Handler: router,
	}

	fmt.Printf("Server Starterd on port: %s \n ", portString)
	erro := srv.ListenAndServe()
	if erro != nil {
		log.Fatal(erro)
	}
}

func LoadEnviroment() string {
	err := godotenv.Load()
	var portString string

	if err != nil {
		log.Fatal("Porta não configurada...")
	} else {
		portString = os.Getenv("PORT")
	}

	fmt.Println("SUCESSO! \n PORT: " + portString)
	return portString
}
