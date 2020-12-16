package server

import (
	"log"
	"net/http"
	"time"
)

func (s *Server) addDelay(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(s.delay)
		h(w, r)
	}
}

func (s *Server) log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request received: %+v", r)
		h(w, r)
		log.Printf("Responding")
	}
}
