package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/FRSiqueiraBR/url-shortener/internal/infra/database"
	urlShortener "github.com/FRSiqueiraBR/url-shortener/internal/usecase/url-shortener"
)

type SurlControllerInterface interface {
	FindAll(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

type SurlController struct {
	FindAllShortUrls urlShortener.FindAllShortUrls
	SaveUrlShort urlShortener.SaveUrlShort 
}

func NewSurlController(urlRepository database.UrlRepositoryInterface) *SurlController {
	return &SurlController{
		FindAllShortUrls: *urlShortener.NewFindAllShortUrls(urlRepository),
		SaveUrlShort: *urlShortener.NewSaveUrlShort(urlRepository),
	}
}

func (s *SurlController) FindAll(w http.ResponseWriter, r *http.Request) {
	dtos, err := s.FindAllShortUrls.FindAll()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(dtos)
}

func (s *SurlController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
    var data map[string]interface{}
    err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
    
    expiration, err := convertStringToTime(data["expiration"].(string))
    if err != nil {
        http.Error(w, err.Error(), 400)
    }

	s.SaveUrlShort.Save(data["url"].(string), data["ip"].(string), expiration)
}

func convertStringToTime(str string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
    
    expiration, err := time.Parse(layout, str,)
    if err != nil {
        return time.Now(), errors.New("Expiration in invalid format")
    }

	return expiration, err
}