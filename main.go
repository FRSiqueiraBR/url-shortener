package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/FRSiqueiraBR/url-shortener/internal/infra/database"
	usecaseUrlShortener "github.com/FRSiqueiraBR/url-shortener/internal/usecase/url-shortener"
	"github.com/joho/godotenv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	goDotEnvVariable("HOST")

	db, err := sql.Open("sqlite3", "infra/database/UrlShortener.db")
	if err != nil {
		panic(err)
	}

	defer db.Close() //espera tudo rodar depois executa o close

	urlRespository := database.NewUrlRepository(db)
	ucSave := usecaseUrlShortener.NewSaveUrlShort(urlRespository)
	ucFindAll := usecaseUrlShortener.NewFindAllShortUrls(urlRespository)

	_, err = ucSave.Save("https://google.com.br", "192.168.0.1", time.Date(2099, time.December, 23, 59, 59, 0, 0, time.UTC))
	if err != nil {
		panic(err)
	}

	entities, err := ucFindAll.FindAll()
	if err != nil {
		panic(err)
	}

	fmt.Println(entities)
}

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}

	return os.Getenv(key)
}
