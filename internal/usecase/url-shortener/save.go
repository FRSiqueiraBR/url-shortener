package urlShortener

import (
	"time"

	"github.com/FRSiqueiraBR/url-shortener/internal/entity"
	"github.com/FRSiqueiraBR/url-shortener/internal/usecase/hash"
	"github.com/FRSiqueiraBR/url-shortener/internal/infra/database"
)

type SaveUrlShortInterface interface {
	Save(url string, ip string) (*entity.ShortUrl, error)
}

type SaveUrlShort struct {
	UrlRepository database.UrlRepositoryInterface
}

func NewSaveUrlShort(urlRepository database.UrlRepositoryInterface) *SaveUrlShort {
	return &SaveUrlShort{
		UrlRepository: urlRepository,
	}
}

func (s *SaveUrlShort) Save(url string, ip string, expiration time.Time) (*entity.ShortUrl, error) {
	hash, err := hash.CreateHash(url, ip)
	if err != nil {
		return &entity.ShortUrl{}, err
	}

	urlEntity, err := entity.NewShortUrl(url, hash, expiration, time.Now())
	if err != nil {
		return &entity.ShortUrl{}, err
	}

	s.UrlRepository.Save(urlEntity)

	return urlEntity, nil
}
