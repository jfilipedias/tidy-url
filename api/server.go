package api

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/jfilipedias/tidy-url/url"
	"github.com/jfilipedias/tidy-url/url/mongodb"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Config struct {
	Port    int
	Env     string
	MongoDB struct {
		URI string
	}
}

type API struct {
	config     Config
	logger     *slog.Logger
	urlUseCase url.UseCase
}

func New(config Config) *API {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	client, err := mongo.Connect(options.Client().ApplyURI(config.MongoDB.URI))
	if err != nil {
		panic(err)
	}

	repo := mongodb.NewMongoDB(client)
	urlUseCase := url.NewService(repo)

	return &API{config, logger, urlUseCase}
}

func (app *API) Serve() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.Port),
		Handler:      app.newRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(app.logger.Handler(), slog.LevelError),
	}

	return srv.ListenAndServe()
}
