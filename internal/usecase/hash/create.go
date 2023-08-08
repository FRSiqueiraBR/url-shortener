package hash

import (
	"crypto/md5"
	"encoding/base32"
	"strconv"
	"strings"
	"time"
)

type CreateHashInterface interface {
	CreateHash(url string, ip string) (string, error)
}

/*
Generating the hash with ip + timestamp in md5 and base32 and getting the first 7 characters.
*/
func CreateHash(url string, ip string) (string, error) {
	now := time.Now()
	sec := now.UnixNano()

	ipWithoutDots := strings.ReplaceAll(ip, ".", "")
	stringToHash := ipWithoutDots + strconv.FormatInt(sec, 10)

	data := []byte(stringToHash)
	md5Hash := md5.Sum(data)
	
	urlShortened := base32.HexEncoding.EncodeToString(md5Hash[:])[0:7]

	return urlShortened, nil
}