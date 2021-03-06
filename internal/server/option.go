package server

import "time"

// WithDelay adds provided duration as server response delay. If no duration provided,
func WithDelay(d time.Duration) Option {
	return func(s *Server) {
		s.isDelayEnabled = true
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

// WithCORSEnabled allows CORS. This is NOT recommended for production use
// cases, and should be used only for testing.
func WithCORSEnabled() Option {
	return func(s *Server) {
		s.isCORSEnabled = true
	}
}

// WithLoggingDisabled disables logging for each incoming request. Even with
// this option, you will see startup and error logs.
func WithLoggingDisabled() Option {
	return func(s *Server) {
		s.isLoggingDisabled = true
	}
}
