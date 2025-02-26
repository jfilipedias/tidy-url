package api

import (
	"net/http"
	"time"
)

func Serve() error {
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return srv.ListenAndServe()
}
