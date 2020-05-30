package scully

import (
	"github.com/gwuah/scully/api"
	"github.com/gwuah/scully/matrix"
)

type Scully struct {
	api    *api.Api
	Matrix *matrix.Matrix
}

func NewScully(token string) (*Scully, error) {
	s := &Scully{}

	apiInstance, err := api.NewApi(token)

	if err != nil {
		return nil, err
	}

	matrixInstance := matrix.NewMatrix(apiInstance)

	s.api = apiInstance
	s.Matrix = matrixInstance

	return s, nil
}

func (s *Scully) Greet() string {
	return "Hello World"
}
