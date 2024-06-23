package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")

	if val == "" {
		return "", errors.New("you are not authorized to perform this action")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("invalid authorization code")

	}
	if vals[0] != "APIKey" {
		return "", errors.New("invalid authorization code")

	}
	return vals[1],nil
}
