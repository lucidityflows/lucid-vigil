package actions

import (
	"context"
)

// Action defines the interface for any defensive action Sentinel can take.
type Action interface {
	Name() string
	Execute(ctx context.Context, data map[string]interface{}) error
}
