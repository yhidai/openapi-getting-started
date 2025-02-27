/*
 * Swagger Getting Started
 *
 * getting started
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A DefaultApiController binds http requests to an api service and writes the service results to the http response
type DefaultApiController struct {
	service DefaultApiServicer
}

// NewDefaultApiController creates a default api controller
func NewDefaultApiController(s DefaultApiServicer) Router {
	return &DefaultApiController{ service: s }
}

// Routes returns all of the api route for the DefaultApiController
func (c *DefaultApiController) Routes() Routes {
	return Routes{ 
		{
			"AddUser",
			strings.ToUpper("Post"),
			"/users",
			c.AddUser,
		},
		{
			"DeleteUser",
			strings.ToUpper("Delete"),
			"/users/${id}",
			c.DeleteUser,
		},
		{
			"GetUser",
			strings.ToUpper("Get"),
			"/users/${id}",
			c.GetUser,
		},
		{
			"ListUsers",
			strings.ToUpper("Get"),
			"/users",
			c.ListUsers,
		},
		{
			"UpdateUser",
			strings.ToUpper("Put"),
			"/users/${id}",
			c.UpdateUser,
		},
	}
}

// AddUser - 
func (c *DefaultApiController) AddUser(w http.ResponseWriter, r *http.Request) { 
	newUser := &NewUser{}
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.AddUser(*newUser)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}

// DeleteUser - 
func (c *DefaultApiController) DeleteUser(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	id, err := parseIntParameter(params["id"])
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.DeleteUser(id)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}

// GetUser - 
func (c *DefaultApiController) GetUser(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	id, err := parseIntParameter(params["id"])
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.GetUser(id)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
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

// UpdateUser - 
func (c *DefaultApiController) UpdateUser(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	id, err := parseIntParameter(params["id"])
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	user := &User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.UpdateUser(id, *user)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}
