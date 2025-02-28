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
	Port     int
	MongoURI string
}

type API struct {
	config     Config
	logger     *slog.Logger
	urlUseCase url.UseCase
}

func New(config Config) *API {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	client, err := mongo.Connect(options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		panic(err)
	}

	repo := mongodb.NewMongoDB(client)
	urlUseCase := url.NewService(repo)

	return &API{config, logger, urlUseCase}
}

func (api *API) Serve() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", api.config.Port),
		Handler:      api.newRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(api.logger.Handler(), slog.LevelError),
	}

	api.logger.Info("starting server", "addr", srv.Addr)
	return srv.ListenAndServe()
}
