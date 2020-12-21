package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/rytswd/color-svc/internal/server"
)

func main() {
	options := []server.Option{}
	delay := envOrDefaultInt("DELAY_DURATION_MILLISECOND", 1000)
	if isEnvTrue("ENABLE_DELAY") {
		options = append(options, server.WithDelay(time.Duration(delay)*time.Millisecond))
	}
	if isEnvTrue("ENABLE_CORS") {
		options = append(options, server.WithCORSEnabled())
	}
	if isEnvTrue("DISABLE_RED") {
		options = append(options, server.WithDisabledRed())
	}
	if isEnvTrue("DISABLE_GREEN") {
		options = append(options, server.WithDisabledGreen())
	}
	if isEnvTrue("DISABLE_BLUE") {
		options = append(options, server.WithDisabledBlue())
	}
	if isEnvTrue("DISABLE_YELLOW") {
		options = append(options, server.WithDisabledYellow())
	}

	mux := http.NewServeMux()
	server.NewServer(mux, options...)

	log.Println("Server starting")

	if err := http.ListenAndServe(":8800", mux); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func isEnvTrue(envName string) bool {
	result := false
	if v, ok := os.LookupEnv(envName); ok {
		result, _ = strconv.ParseBool(strings.ToLower(v))
	}
	return result
}

func envOrDefault(envName, fallback string) string {
	result := fallback
	if v, ok := os.LookupEnv(envName); ok {
		result = v
	}
	return result
}

func envOrDefaultInt(envName string, fallback int) int {
	v, ok := os.LookupEnv(envName)
	if !ok {
		return fallback
	}
	result, err := strconv.Atoi(v)
	if err != nil {
		return fallback
	}
	return result
}
