package api

import "net/url"

func IsValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	} else {
		return true
	}
}