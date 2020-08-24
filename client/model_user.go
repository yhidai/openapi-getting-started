/*
 * Swagger Getting Started
 *
 * getting started
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// User struct for User
type User struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Active bool `json:"active,omitempty"`
}
