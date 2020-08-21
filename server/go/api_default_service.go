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
	_ "errors"
)

// DefaultApiService is a service that implents the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type DefaultApiService struct {
}

// NewDefaultApiService creates a default api service
func NewDefaultApiService() DefaultApiServicer {
	return &DefaultApiService{}
}

// ListUsers -
func (s *DefaultApiService) ListUsers(active string) (interface{}, error) {
	// TODO - update ListUsers with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return []*User{{Name: "tom"}}, nil

	// return nil, errors.New("service method 'ListUsers' not implemented")
}
