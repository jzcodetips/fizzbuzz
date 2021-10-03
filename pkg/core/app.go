package core

import (
	"context"
	"fizzbuzz/pkg/controller"

	"github.com/pkg/errors"
)

// errors.
var (
	ErrDependenciesRequired = errors.New("Dependencies is required")
	ErrAPIRequired          = errors.New("API is required")
)

// Dependencies holds app deps.
type Dependencies struct {
	API controller.API
}

// Validate Dependencies.
func (d *Dependencies) Validate() error {
	if d.API == nil {
		return ErrAPIRequired
	}

	return nil
}

// App holds the server and his configuration.
type App struct {
	deps *Dependencies
}

// NewApp returns a new instance of API.
func NewApp(deps *Dependencies) (*App, error) {
	if deps == nil {
		return nil, ErrDependenciesRequired
	}

	if err := deps.Validate(); err != nil {
		return nil, err
	}

	return &App{deps: deps}, nil
}

// Run the App.
func (a *App) Run(ctx context.Context) error {
	return a.deps.API.Start(ctx)
}
