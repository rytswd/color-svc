package server

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
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

		writeResponse(w, r.URL.Query(), c)
	}
}

func (s *Server) handleSpecific() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(s.allColors) == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("ERROR: no color available\n"))
			return
		}

		requestedColor := strings.ReplaceAll(r.URL.Path, "/get/", "") // Assumes /get/ as request uri
		c := color.Color{Name: "unknown", HexCode: "#888888"}
		for _, color := range s.allColors {
			if compareColorInput(requestedColor, color.Name) {
				c = color
				break
			}
		}

		writeResponse(w, r.URL.Query(), c)
	}
}

func (s *Server) handleRed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if s.disableRed {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("ERROR: red is disabled\n"))
			return
		}

		rand.Seed(now().UTC().UnixNano())
		c := color.RedishColors[rand.Intn(len(color.RedishColors))]

		writeResponse(w, r.URL.Query(), color.Color(c))
	}
}

func (s *Server) handleGreen() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if s.disableGreen {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("ERROR: green is disabled\n"))
			return
		}

		rand.Seed(now().UTC().UnixNano())
		c := color.GreenishColors[rand.Intn(len(color.GreenishColors))]

		writeResponse(w, r.URL.Query(), color.Color(c))
	}
}

func (s *Server) handleBlue() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if s.disableBlue {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("ERROR: blue is disabled\n"))
			return
		}

		rand.Seed(now().UTC().UnixNano())
		c := color.BluishColors[rand.Intn(len(color.BluishColors))]

		writeResponse(w, r.URL.Query(), color.Color(c))
	}
}

func (s *Server) handleYellow() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if s.disableYellow {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("ERROR: yellow is disabled\n"))
			return
		}

		rand.Seed(now().UTC().UnixNano())
		c := color.YellowishColors[rand.Intn(len(color.YellowishColors))]

		writeResponse(w, r.URL.Query(), color.Color(c))
	}
}

func compareColorInput(uri, color string) bool {
	c := strings.ToLower(color)
	u := strings.ToLower(strings.ReplaceAll(uri, "/", ""))
	return c == u
}

func writeResponse(w http.ResponseWriter, query url.Values, c color.Color) {
	var format string
	if f, ok := query["fmt"]; ok {
		format = f[0]
	}

	switch format {
	case "json":
		w.Header().Set("Content-Type", "application/json")
		j, err := json.Marshal(c)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(j)
	default:
		result := fmt.Sprintf("  Generated Color\n    \"%s\" - with HEX \"%s\"\n", c.Name, c.HexCode)
		log.Printf("About to respond with:\n%s", result)
		w.Write([]byte(result))
	}
}
