package gifs

import (
	"encoding/json"
	"errors"
	"fmt"
	"gus/static"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

var (
	rateLimited     bool
	rateLimitedLock sync.Mutex
)

type tenorResponse struct {
	Results []struct {
		URL string `json:"url"`
	} `json:"Results"`
}

func SearchForGif(searchTerm string) (results []string, err error) {

	var sanitizedSearchTerm string

	for _, char := range searchTerm {

		if string(char) == " " {
			sanitizedSearchTerm += "%20"
		}

		if strings.Contains(static.ALPHABET, string(char)) {
			sanitizedSearchTerm += string(char)
		}
	}

	url := fmt.Sprintf(static.TENOR_URL, sanitizedSearchTerm, static.TENOR_API_KEY)
	resp, err := http.Get(url)

	if err != nil {
		return results, err
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return results, err
	}

	var gifs tenorResponse
	err = json.Unmarshal(b, &gifs)
	if err != nil {
		return results, err
	}

	if len(gifs.Results) > 0 {
		for _, gif := range gifs.Results {
			results = append(results, gif.URL)
		}
	} else {
		return results, errors.New(static.ERROR_TENOR_NO_GIFS)
	}

	switch resp.StatusCode {
	case 200:
		return results, nil
	case 429:
		return results, errors.New(static.ERROR_TENOR_RATE_LIMIT)
	default:
		return results, err
	}
}

func IsRateLimited() bool {
	defer rateLimitedLock.Unlock()
	rateLimitedLock.Lock()

	return rateLimited
}

func RateLimit() {
	rateLimitedLock.Lock()
	if rateLimited {
		return
	}
	rateLimitedLock.Unlock()

	go func() {
		rateLimitedLock.Lock()
		rateLimited = true
		rateLimitedLock.Unlock()

		time.Sleep(time.Second * 30)

		rateLimitedLock.Lock()
		rateLimited = false
		rateLimitedLock.Unlock()
	}()
}
