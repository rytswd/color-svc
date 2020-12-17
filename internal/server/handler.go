package server

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/rytswd/color-svc/internal/color"
)

var now = time.Now

func (s *Server) handleRandom() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(s.allColors) == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("ERROR: no color available\n"))
			return
		}

		rand.Seed(now().UTC().UnixNano())
		c := s.allColors[rand.Intn(len(s.allColors))]
		result := fmt.Sprintf("  Generated Color\n    \"%s\" - with HEX \"%s\"\n", c.Name, c.HexCode)

		log.Printf("About to respond with:\n%s", result)

		w.Write([]byte(result))
	}
}

func (s *Server) handleSpecific() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(s.allColors) == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("ERROR: no color available\n"))
			return
		}

		requestedColor := strings.ReplaceAll(r.RequestURI, "/get/", "") // Assumes /get/ as request uri
		c := color.Color{Name: "unknown", HexCode: "#888888"}
		for _, color := range s.allColors {
			if compareColorInput(requestedColor, color.Name) {
				c = color
				break
			}
		}
		result := fmt.Sprintf("  Generated Color\n    \"%s\" - with HEX \"%s\"\n", c.Name, c.HexCode)

		log.Printf("About to respond with:\n%s", result)

		w.Write([]byte(result))
	}
}

func (s *Server) handleRed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rand.Seed(now().UTC().UnixNano())
		c := color.RedishColors[rand.Intn(len(color.RedishColors))]
		result := fmt.Sprintf("  Generated Color\n    \"%s\" - with HEX \"%s\"\n", c.Name, c.HexCode)

		log.Printf("About to respond with:\n%s", result)

		w.Write([]byte(result))
	}
}

func (s *Server) handleGreen() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rand.Seed(now().UTC().UnixNano())
		c := color.GreenishColors[rand.Intn(len(color.GreenishColors))]
		result := fmt.Sprintf("  Generated Color\n    \"%s\" - with HEX \"%s\"\n", c.Name, c.HexCode)

		log.Printf("About to respond with:\n%s", result)

		w.Write([]byte(result))
	}
}
func (s *Server) handleBlue() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rand.Seed(now().UTC().UnixNano())
		c := color.BluishColors[rand.Intn(len(color.BluishColors))]
		result := fmt.Sprintf("  Generated Color\n    \"%s\" - with HEX \"%s\"\n", c.Name, c.HexCode)

		log.Printf("About to respond with:\n%s", result)

		w.Write([]byte(result))
	}
}
func (s *Server) handleYellow() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rand.Seed(now().UTC().UnixNano())
		c := color.YellowishColors[rand.Intn(len(color.YellowishColors))]
		result := fmt.Sprintf("  Generated Color\n    \"%s\" - with HEX \"%s\"\n", c.Name, c.HexCode)

		log.Printf("About to respond with:\n%s", result)

		w.Write([]byte(result))
	}
}

func compareColorInput(uri, color string) bool {
	c := strings.ToLower(color)
	u := strings.ToLower(strings.ReplaceAll(uri, "/", ""))
	return c == u
}
