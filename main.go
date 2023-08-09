package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/FRSiqueiraBR/url-shortener/internal/controller"
	"github.com/FRSiqueiraBR/url-shortener/internal/infra/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	loadEnvVariables()

	db, err := sql.Open("sqlite3", "infra/database/UrlShortener.db")
	if err != nil {
		panic(err)
	}

	defer db.Close() //espera tudo rodar depois executa o close

	urlRespository := database.NewUrlRepository(db)
	surlController := controller.NewSurlController(urlRespository)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/surl", surlController.FindAll)
	r.Post("/surl", surlController.Create)

	http.ListenAndServe(os.Getenv("HOST"), r)
}

func loadEnvVariables() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}
}
