package core_http_server

import (
	"fmt"
	"net/http"
)

type ApiVersion string

var (
	ApiVersion1 = ApiVersion("v1")
	ApiVersion2 = ApiVersion("v2")
	ApiVersion3 = ApiVersion("v3")
)

type APIVersionRoter struct {
	*http.ServeMux
	apiVersion ApiVersion
}

func NewAPIVersionRouter(apiVersion ApiVersion) *APIVersionRoter {
	return &APIVersionRoter{
		ServeMux:   http.NewServeMux(),
		apiVersion: apiVersion,
	}
}

func (r *APIVersionRoter) RegisterRoutes(routes ...Route) {
	for _, router := range routes {
		pattern := fmt.Sprintf("%s %s", router.Method, router.Path)

		r.Handle(pattern, router.Handler)
	}
}
