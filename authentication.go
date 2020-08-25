package server

import (
	"encoding/json"
	"log"
	"net/http"
)

const ApiKey = "xxxxxxxxxxxx"

// Authenticate authenticates a request whether the request has a correct APIKey
func Authenticate(inner http.Handler, apiKey string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		key := r.Header.Get("X-Api-Key")
		successfullyAuthenticated := key == apiKey

		log.Printf(
			"%s %s auth result %v",
			r.Method,
			r.RequestURI,
			successfullyAuthenticated,
		)

		if !successfullyAuthenticated {
			w.Header().Set("Content-Type", "application/json")
			// w.Header().Set("WWW-Authentication", "API key is missing or invalid.")
			w.WriteHeader(401)

			response := Error{
				Code:    401,
				Message: "API key is missing or invalid.",
			}
			encoder := json.NewEncoder(w)
			err := encoder.Encode(response)
			if err != nil {
				log.Printf("write error response failed. %v\n", err)
			}
			return
		}

		// clall next operation
		inner.ServeHTTP(w, r)
	})
}
