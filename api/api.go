package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

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
	return &Api{token: token}, nil
}

func (a *Api) BuildRequest(params []string, queryParams map[string]string) (string, error) {
	urlParams := strings.Join(params, "/")
	req, err := http.NewRequest("GET", "https://api.mapbox.com/"+urlParams, nil)

	if err != nil {
		return "", err
	}

	q := req.URL.Query()

	for key, value := range queryParams {
		q.Add(key, value)
	}
	q.Add("access_token", a.token)
	req.URL.RawQuery = q.Encode()

	return req.URL.String(), nil
}

func (a *Api) MakeRequest(url string) (map[string]interface{}, error) {
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == 422 {
		return nil, errors.New("Invalid Input")
	} else if res.StatusCode == 404 {
		return nil, errors.New("Not Found")
	}

	data, _ := ioutil.ReadAll(res.Body)

	res.Body.Close()

	var jsonResponse map[string]interface{}

	err = json.Unmarshal([]byte(data), &jsonResponse)

	if err != nil {
		return nil, err
	}

	fmt.Println(jsonResponse)

	return jsonResponse, nil
}
