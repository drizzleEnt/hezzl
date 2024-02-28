package app

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/drizzleent/hezzl/internal/config"
)

type App struct {
	serviceProvicer *serviceProvicer
	httpServer      *http.Server
}

func New(ctx context.Context) (*App, error) {
	a := &App{}
	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	err := a.runHTTPServer()
	if err != nil {
		log.Fatalf("failed to run http server %s", err.Error())
	}

	return nil

}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initHTTPServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}
	return nil
}

func (a *App) initServiceProvider(ctx context.Context) error {
	a.serviceProvicer = newServiceProvider()
	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	srv := &http.Server{
		Addr:           a.serviceProvicer.HttpConfig().Address(),
		Handler:        a.serviceProvicer.Handler(ctx).InitRoutes(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	a.httpServer = srv
	return nil
}

func (a *App) runHTTPServer() error {
	log.Printf("Server running on: ")

	err := a.httpServer.ListenAndServe()

	if err != nil {
		return err
	}

	return nil
}
