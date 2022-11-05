//go:generate mockgen -source=cache.go -destination=cache_mocks.go -package=cache

package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

var (
	ErrNotFound = errors.New("not found")
)

type (
	Key interface {
		fmt.Stringer
		Int64() int64
	}

	Cachable interface {
		json.Marshaler
		json.Unmarshaler
		Key() Key
	}

	Cache interface {
		Set(context.Context, Cachable) error
		Get(context.Context, Cachable) error
		Drop(context.Context, Key) error
		Invalidate(context.Context) error
	}
)
