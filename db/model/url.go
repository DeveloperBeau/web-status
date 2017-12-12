package model

// URL Struct
type Url struct {
	Id  int    `json:"id"`
	Url string `json:"url"`
	IsEnabled bool `json:"enabled"`
}
