package main

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/jfilipedias/tidy-url/api"
)

func getEnvWithDefault(key, defaultValue string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}
	return defaultValue
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	port, err := strconv.Atoi(getEnvWithDefault("API_PORT", "8080"))
	if err != nil {
		logger.Error("invalid port value", "err", err)
		os.Exit(1)
	}

	mongoURI := getEnvWithDefault("MONGODB_URI", "")
	if mongoURI == "" {
		logger.Error("mongodb uri is required")
		os.Exit(1)
	}

	cfg := api.Config{
		Port:     port,
		MongoURI: mongoURI,
	}

	app := api.New(cfg)
	app.Serve()
}
