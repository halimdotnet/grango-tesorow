package hxxp

import (
	"context"
)

type Server interface {
	RunServer() error
	Shutdown(ctx context.Context) error
	BuildRouter() *Router
}
