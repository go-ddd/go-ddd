package contexts

import (
	"context"
)

type key int

const (
	instanceIDKey key = iota + 1
)

func WithInstanceID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, instanceIDKey, id)
}

func GetInstanceID(ctx context.Context) string {
	instance, ok := ctx.Value(instanceIDKey).(string)
	if !ok {
		return ""
	}
	return instance
}
