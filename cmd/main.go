package main

import (
	"e-commerce/config"
	"e-commerce/internal/app/handler"
	"e-commerce/internal/app/service"
	"e-commerce/internal/app/storage"
	"e-commerce/pkg/gorm"

	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal("main: ", err)
	}
}

func run() error {
	cfg, err := config.NewConfig(config.DefaultPath)
	if err != nil {
		return err
	}
	orm, err := gorm.New(cfg.PostgresDSN)
	if err != nil {
		return err
	}
	app := handler.Handler{
		Service: service.Service{
			Storage: storage.Storage{
				ORM: orm,
			},
			Config: *cfg,
		},
	}

	router := handler.NewRouter(app, cfg)

	addr := net.JoinHostPort(cfg.AppHost, cfg.AppPort)
	server := http.Server{
		Addr:              addr,
		Handler:           router,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
	}

	errCh := make(chan error)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- fmt.Errorf("listen and serve: %w", err)
		}
	}()
	wg.Wait()
	return err
}
