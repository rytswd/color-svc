package server

import "time"

// WithDelay adds provided duration as server response delay. If no duration provided,
func WithDelay(d time.Duration) Option {
	return func(s *Server) {
		s.enableDelay = true
		if d != 0 {
			d = 1 * time.Second
		}
		s.delay = d
	}
}

// WithDisabledRed disables redish colors.
func WithDisabledRed() Option {
	return func(s *Server) {
		s.disableRed = true
	}
}

// WithDisabledGreen disables greenish colors.
func WithDisabledGreen() Option {
	return func(s *Server) {
		s.disableGreen = true
	}
}

// WithDisabledBlue disables bluish colors.
func WithDisabledBlue() Option {
	return func(s *Server) {
		s.disableBlue = true
	}
}

// WithDisabledYellow disables yellowish colors.
func WithDisabledYellow() Option {
	return func(s *Server) {
		s.disableYellow = true
	}
}
