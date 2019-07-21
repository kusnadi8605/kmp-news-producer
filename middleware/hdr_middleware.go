package handler

import (
	"encoding/json"
	dts "kmp-news-producer/datastruct"
	"net/http"
)

//Middleware ..
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

// Method ensures that url can only be requested with a specific method, else returns a 400 Bad Request
func Method(m string) Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			var MiddlewareResponse dts.MiddlewareResponse
			if r.Method != m {
				MiddlewareResponse.ResponseCode = http.StatusBadRequest
				MiddlewareResponse.ResponseDesc = http.StatusText(http.StatusBadRequest) + ": Invalid Request Methods"
				json.NewEncoder(w).Encode(MiddlewareResponse)
				return
			}
			f(w, r)
		}
	}
}

//ContentType ..
func ContentType(m string) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			var MiddlewareResponse dts.MiddlewareResponse
			if r.Header.Get("Content-Type") != m {
				MiddlewareResponse.ResponseCode = http.StatusBadRequest
				MiddlewareResponse.ResponseDesc = http.StatusText(http.StatusBadRequest) + ": Invalid Http Header"
				json.NewEncoder(w).Encode(MiddlewareResponse)
				return
			}
			f(w, r)
		}
	}
}
