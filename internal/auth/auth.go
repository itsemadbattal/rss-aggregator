package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey, it extracts an api key from the headers of an http request
// Example, Authorization: APIKey {insert APIKey here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no Authentication info found!")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("maloformed auth header!")
	}

	if vals[0] != "APIKey" {
		return "", errors.New("maloformed first part of auth header!")
	}

	return vals[1], nil
}
