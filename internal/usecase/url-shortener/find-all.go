package urlShortener

import (
	"github.com/FRSiqueiraBR/url-shortener/internal/dto"
	"github.com/FRSiqueiraBR/url-shortener/internal/entity"
	"github.com/FRSiqueiraBR/url-shortener/internal/infra/database"
)


type FindAllShortUrlsInterface interface {
	FindAll() (*[]entity.ShortUrl, error)
}

type FindAllShortUrls struct {
	UrlRepository database.UrlRepositoryInterface
}

func NewFindAllShortUrls(urlRepository database.UrlRepositoryInterface) *FindAllShortUrls {
	return &FindAllShortUrls{
		UrlRepository: urlRepository,
	}
}

func (f *FindAllShortUrls) FindAll() (*[]dto.Surl, error) {
	entities, err := f.UrlRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var dtos []dto.Surl

	for _, ent := range *entities {
		surl, err := ent.ToDTO()
		if err != nil {
			return nil, err
		}
		
		dtos = append(dtos, *surl)
    }

	return &dtos, err
}