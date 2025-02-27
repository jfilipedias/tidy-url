package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/jfilipedias/tidy-url/api"
)

func main() {
	var cfg api.Config

	flag.IntVar(&cfg.Port, "port", 4000, "API server port")
	flag.StringVar(&cfg.Env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.MongoDB.URI, "db-uri", "", "MongoDB URI")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := api.NewApp(cfg, logger)
	app.Serve()
}
