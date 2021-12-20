package infrastructure

import (
	"context"
)

// Redis client.
//
// A redis client will be called in many goroutines, so it must be implemented gocoutine-safe.
type Redis interface {
	// Do a redis command and return it's result.
	Do(ctx context.Context, cmd string, args ...interface{}) (interface{}, error)
}
