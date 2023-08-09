package dto

import (
	"errors"
	"time"
)

type Surl struct {
	Long       string `json:"long"`
	Hash       string `json:"hash"`
	Expiration time.Time `json:"expiration"`
	Timestamp  time.Time `json:"timestamp"`
}

type SurlInteface interface {
	NewSurl(long string, hash string, expiration time.Time, timestamp time.Time)
}

func NewSurl(long string, hash string, expiration time.Time, timestamp time.Time) (*Surl, error) {
	url := &Surl{
		Long:       long,
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

func (u *Surl) Validate() error {
	if u.Long == "" {
		return errors.New("full url is required")
	} else if u.Hash == "" {
		return errors.New("Hash is required")
	}

	return nil
}
