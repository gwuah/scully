package matrix

import (
	"github.com/gwuah/scully/api"
)

type Matrix struct {
	api *api.Api
}

func NewMatrix(apiInstance *api.Api) *Matrix {
	return &Matrix{api: apiInstance}
}
