package database

import (
	"database/sql"

	"github.com/FRSiqueiraBR/url-shortener/internal/entity"
)

type UrlRepositoryInterface interface {
	Save(url *entity.ShortUrl) error
	FindAll() (*[]entity.ShortUrl, error)
}

type UrlRepository struct {
	Db *sql.DB
}

func NewUrlRepository(db *sql.DB) *UrlRepository {
	return &UrlRepository{
		Db: db,
	}
}

func (r *UrlRepository) Save(url *entity.ShortUrl) error {
	_, err := r.Db.Exec("INSERT INTO short_url(long, hash, expiration, timestamp) values (?,?,?,?)", url.Long, url.Hash, url.Expiration, url.Timestamp)
	if err != nil {
		return err
	}
	return nil
}

func (r *UrlRepository) FindAll() (*[]entity.ShortUrl, error) {
	rows, err := r.Db.Query("SELECT long, hash, expiration, timestamp FROM short_url")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    // An album slice to hold data from returned rows.
    var shortUrls []entity.ShortUrl

    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var shortUrl entity.ShortUrl
        if err := rows.Scan(&shortUrl.Long, &shortUrl.Hash, &shortUrl.Expiration, &shortUrl.Timestamp); err != nil {
            return &shortUrls, err
        }
        shortUrls = append(shortUrls, shortUrl)
    }
    if err = rows.Err(); err != nil {
        return &shortUrls, err
    }
    return &shortUrls, nil
}
