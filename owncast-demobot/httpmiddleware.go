package main

import (
	"log"
	"net/http"
)

func RequireHttpPost(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Unsupported Method", http.StatusMethodNotAllowed)
			return;
		}

		next.ServeHTTP(w, r)
	   }
}

func RequireJsonContentType(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			http.Error(w, "The specified ContentType is not supported", http.StatusUnsupportedMediaType)
		}
		next.ServeHTTP(w, r)
	}
}

func LogRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Tracing request for '%s' from '%s'", r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
	   }
}