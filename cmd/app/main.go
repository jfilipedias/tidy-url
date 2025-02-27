package main

import (
	"flag"

	"github.com/jfilipedias/tidy-url/api"
)

func main() {
	var c api.Config
	flag.IntVar(&c.Port, "port", 4000, "API server port")
	flag.StringVar(&c.Env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&c.MongoDB.URI, "db-uri", "", "MongoDB URI")
	flag.Parse()

	app := api.New(c)
	app.Serve()
}
