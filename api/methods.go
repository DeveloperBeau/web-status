package api

import "net/url"

func IsValidUrl(u string) (string, error) {
	_, err := url.ParseRequestURI(u)
	if err != nil {
		u = "http://" + u
		_, err := url.ParseRequestURI(u)
		if err != nil {
			return "", ErrCorruptData
		}
	}
	return u, nil
}