package server

import (
	"context"
	"fizzbuzz/pkg/controller/rest/handler"
	"fizzbuzz/pkg/core/service"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// errors.
var (
	ErrConfigRequired = errors.New("config is required")
	ErrPortRequired   = errors.New("port is required")
)

// Config schema.
type Config struct {
	Port string
}

// Validate AppConfig.
func (s *Config) Validate() error {
	if s.Port == "" {
		return ErrPortRequired
	}

	return nil
}

// Server struct.
type Server struct {
	config *Config
	e      *echo.Echo
}

// Option type.
type Option func(s *Server)

// WithMiddlewares option.
func WithMiddlewares(middlewares ...echo.MiddlewareFunc) Option {
	return func(s *Server) {
		s.e.Use(middlewares...)
	}
}

func NewServer(config *Config, opts ...Option) (*Server, error) {
	if config == nil {
		return nil, ErrConfigRequired
	}

	if err := config.Validate(); err != nil {
		return nil, errors.Wrap(err, "config.Validate")
	}

	s := &Server{
		config: config,
		e:      echo.New(),
	}

	for _, opt := range opts {
		if opt != nil {
			opt(s)
		}
	}

	return s, nil
}

// InitRoutes.
func (s *Server) InitRoutes() error {
	//  endpoint.
	serviceFizzBuzz := service.NewFizzBuzz()
	fizzBuzz, err := handler.NewFizzBuzz(serviceFizzBuzz)
	if err != nil {
		return errors.Wrap(err, "handler.NewFizzBuzz")
	}

	// FizzBuzz
	s.e.GET("/fizzbuzz", fizzBuzz.Handle)

	s.e.GET("/metrics", handler.Metrics)

	return nil
}

// Start start the server and listen from the given port.
func (s *Server) Start(ctx context.Context) error {
	return s.e.Start(":" + s.config.Port)
}

// Shutdown the server.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.e.Shutdown(ctx)
}
