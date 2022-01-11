package goshort

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/url"
)

type URL struct {
	UrlCode  string `json:"code"`
	ShortUrl string `json:"short_url"`
	LongUrl  string `json:"long_url"`
}

const BASE_URL string = "https://goshort.io"

var urls = []URL{}

func uuid(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	ret := make([]byte, n)

	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

// Looks up a URL by the provided code and returns it's long url, or errors if the URL is not found
func GetUrl(c string) (URL, error) {
	var url URL
	var err error

	for i := range urls {
		if urls[i].UrlCode == c {
			url = urls[i]
			break
		} else {
			err = fmt.Errorf("no url found")
			break
		}
	}

	return url, err
}

// creates a short URL based on the provided long URL
func ShortenUrl(longUrl string) (URL, error) {
	// validate url
	_, err := url.ParseRequestURI(longUrl)
	if err != nil {
		return URL{}, err
	}

	// make sure the long url does not exist in DB, if so, just return its short one
	for _, u := range urls {
		if u.LongUrl == longUrl {
			return u, fmt.Errorf("existing record found")
		}
	}

	// generate unique url code(short id)
	code, _ := uuid(8)

	// create a new url and return the short one: which is: BASE_URL/url code
	u := URL{
		UrlCode:  code,
		ShortUrl: fmt.Sprintf("%s/%s", BASE_URL, code),
		LongUrl:  longUrl,
	}

	urls = append(urls, u)

	return u, nil
}
