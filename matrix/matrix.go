package matrix

import (
	"fmt"

	"github.com/gwuah/scully/api"
)

type Matrix struct {
	api              *api.Api
	destinationIndex string
}

func NewMatrix(apiInstance *api.Api) *Matrix {
	return &Matrix{api: apiInstance}
}

func (m *Matrix) SetDestinationIndex(index string) {
	m.destinationIndex = index
}

func (m *Matrix) GetMatrix(points string) (map[string]interface{}, error) {
	urlParts := []string{"directions-matrix", "v1", "mapbox/driving", points}

	requestUrl, err := m.api.BuildRequest(urlParts, map[string]string{
		"annotations":  "distance,duration",
		"sources":      "all",
		"destinations": m.destinationIndex,
	})

	if err != nil {
		return nil, err
	}

	fmt.Println(requestUrl)
	response, err := m.api.MakeRequest(requestUrl)

	if err != nil {
		return nil, err
	}

	return response, nil

}
