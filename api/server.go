package api

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	errShutdown := make(chan error, 1)
	go api.shutdown(srv, errShutdown)

	api.logger.Info("starting server", "addr", srv.Addr)
	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-errShutdown
	if err != nil {
		return err
	}

	api.logger.Info("stopped server", "addr", srv.Addr)
	return nil
}

func (api *API) shutdown(server *http.Server, errShutdown chan error) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	s := <-quit

	api.logger.Info("caught signal", "signal", s.String())

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		errShutdown <- err
	}
	errShutdown <- nil
}
