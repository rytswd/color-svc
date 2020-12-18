package server

import (
	"log"
	"net/http"
	"time"

	"github.com/rytswd/color-svc/internal/color"
)

// Server holds onto routing setup and dependencies.
type Server struct {
	router *http.ServeMux

	middlewares []func(h http.HandlerFunc) http.HandlerFunc
	allColors   []color.Color

	// Option driven fields
	enableDelay   bool
	delay         time.Duration
	disableRed    bool
	disableGreen  bool
	disableBlue   bool
	disableYellow bool
}

// Option allows modifying the default server setup.
type Option func(*Server)

// NewServer creates fully functional server. You can provide Option to tweak
// its behaviour.
func NewServer(router *http.ServeMux, opts ...Option) {
	s := &Server{
		router: router,
	}

	for _, opt := range opts {
		opt(s)
	}

	// Set up the full list of colors
	s.allColors = color.EnabledColors(!s.disableRed, !s.disableGreen, !s.disableBlue, !s.disableYellow)

	s.compileMiddlewares()
	s.routes()

	s.logStartupSetup()
}

func (s *Server) compileMiddlewares() {
	middlewares := []func(h http.HandlerFunc) http.HandlerFunc{}

	// Make sure log middleware is the first one
	middlewares = append(middlewares, s.log)

	// Delay handling
	if s.enableDelay {
		middlewares = append(middlewares, s.addDelay)
	}

	// Reverse the order of middlewares, this is to ensure that middlewares
	// are applied in the definition order.
	//
	// Example:
	//   Middlewares {A, B, C}
	// 	 As is    -> C(B(A(someHandler()))) -> C gets called first
	//   Reversed -> A(B(C(someHandler()))) -> A gets called first
	for i, j := 0, len(middlewares)-1; i < j; i, j = i+1, j-1 {
		middlewares[i], middlewares[j] = middlewares[j], middlewares[i]
	}

	s.middlewares = middlewares
}

func (s *Server) routes() {
	s.registerHandleFunc("/random", s.handleRandom())

	s.registerHandleFunc("/random/red", s.handleRed())
	s.registerHandleFunc("/random/green", s.handleGreen())
	s.registerHandleFunc("/random/blue", s.handleBlue())
	s.registerHandleFunc("/random/yellow", s.handleYellow())

	s.registerHandleFunc("/", s.handleSpecific())
}

// registerHandleFunc is essentially (*http.ServeMux).HandleFunc, but it chains
// some middlewares.
func (s *Server) registerHandleFunc(path string, h http.HandlerFunc) {
	wrapped := h
	for _, m := range s.middlewares {
		wrapped = m(wrapped)
	}

	s.router.HandleFunc(path, wrapped)
}

func (s *Server) logStartupSetup() {
	log.Printf(`Server setup complete.

	Per Request Delay: %s

	Red    : %s
	Green  : %s
	Blue   : %s
	Yellow : %s

	Total Available Colors: %d

`,
		boolToEnabled(s.enableDelay),
		boolToEnabled(!s.disableRed),
		boolToEnabled(!s.disableGreen),
		boolToEnabled(!s.disableBlue),
		boolToEnabled(!s.disableYellow),
		len(s.allColors))
}

func boolToEnabled(b bool) string {
	if b {
		return "Enabled"
	}
	return "Disabled"
}
