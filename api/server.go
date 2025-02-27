package api

import (
	"log/slog"
	"net/http"
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

type App struct {
	config Config
	logger *slog.Logger
}

func NewApp(config Config, logger *slog.Logger) *App {
	return &App{config, logger}
}

func (app *App) Serve() error {
	client, err := mongo.Connect(options.Client().ApplyURI(app.config.MongoDB.URI))
	if err != nil {
		panic(err)
	}

	repo := mongodb.NewMongoDB(client)
	urlUseCase := url.NewService(repo)

	handler := NewRoutes(urlUseCase)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(app.logger.Handler(), slog.LevelError),
	}

	return srv.ListenAndServe()
}
