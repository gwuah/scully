package api

import "errors"

const (
	BaseURL = "https://api.mapbox.com"
)

type Api struct {
	token string
}

func NewApi(token string) (*Api, error) {
	if token == "" {
		return nil, errors.New("API token is required")
	}

	return &Api{token}, nil
}
