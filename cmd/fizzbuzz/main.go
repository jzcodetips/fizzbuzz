package main

import (
	"context"
	"fizzbuzz/pkg/controller/rest/server"
	"fizzbuzz/pkg/core"
	"flag"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
)

// flags.
var pFlag = flag.String("p", "", "port to expose for the API.")

func init() {
	flag.Parse()
}

func setupDependencies() (*core.Dependencies, error) {
	serverCfg := &server.Config{Port: *pFlag}

	s, err := server.NewServer(serverCfg, server.WithMiddlewares(middleware.Logger(), middleware.Recover()))
	if err != nil {
		return nil, errors.Wrap(err, "server.NewServer")
	}

	if err = s.InitRoutes(); err != nil {
		return nil, errors.Wrap(err, "s.InitRoutes")
	}

	deps := &core.Dependencies{API: s}

	return deps, nil
}

func startApp(deps *core.Dependencies) error {
	app, err := core.NewApp(deps)
	if err != nil {
		return errors.Wrap(err, "core.NewApp")
	}

	ctx := context.Background()

	// create a context with cancel.
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// set signal catching.
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {
		defer wg.Done()

		signal := <-c

		log.Printf("signal %v catched", signal)
		cancel()

		ctxWithTimeout, cancelTimeout := context.WithTimeout(context.Background(), time.Second*3)
		defer cancelTimeout()

		log.Printf("shutdown server ...")

		if err := deps.API.Shutdown(ctxWithTimeout); err != nil {
			log.Printf("server.Shutdown: err=%v", err)
		}

		log.Printf("server is shutdown!")
	}()

	if err = app.Run(ctx); err != nil {
		log.Printf("app.Run end with err=%v", err)
	}

	wg.Wait()

	return nil
}

func main() {
	deps, err := setupDependencies()
	if err != nil {
		log.Printf("app end with err=%v", err)
	}

	startApp(deps)
}
