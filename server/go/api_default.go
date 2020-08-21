/*
 * Swagger Getting Started
 *
 * getting started
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	_ "encoding/json"
	"net/http"
	"strings"

	_ "github.com/gorilla/mux"
)

// A DefaultApiController binds http requests to an api service and writes the service results to the http response
type DefaultApiController struct {
	service DefaultApiServicer
}

// NewDefaultApiController creates a default api controller
func NewDefaultApiController(s DefaultApiServicer) Router {
	return &DefaultApiController{service: s}
}

// Routes returns all of the api route for the DefaultApiController
func (c *DefaultApiController) Routes() Routes {
	return Routes{
		{
			"ListUsers",
			strings.ToUpper("Get"),
			"/users",
			c.ListUsers,
		},
	}
}

// ListUsers -
func (c *DefaultApiController) ListUsers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	active := query.Get("active")
	result, err := c.service.ListUsers(active)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	EncodeJSONResponse(result, nil, w)
}
