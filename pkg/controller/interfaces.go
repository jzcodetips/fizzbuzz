package controller

import "context"

// API interface.
type API interface {
	Start(ctx context.Context) error
	Shutdown(ctx context.Context) error
}
