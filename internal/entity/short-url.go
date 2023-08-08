package entity

import (
	"errors"
	"time"
)

type shortUrlInteface interface {
	NewShortUrl(fullUrl string, hash string, expiration time.Time, timestamp time.Time)
}

type ShortUrl struct {
	Long    string
	Hash       string
	Expiration time.Time
	Timestamp  time.Time
}

func NewShortUrl(fullUrl string, hash string, expiration time.Time, timestamp time.Time) (*ShortUrl, error) {
	url := &ShortUrl{
		Long:    fullUrl,
		Hash:       hash,
		Expiration: expiration,
		Timestamp:  timestamp,
	}

	err := url.Validate()
	if err != nil {
		return nil, err
	}

	return url, nil
}

func (u *ShortUrl) Validate() error {
	if u.Long == "" {
		return errors.New("full url is required")
	} else if u.Hash == "" {
		return errors.New("Hash is required")
	}

	return nil
}
