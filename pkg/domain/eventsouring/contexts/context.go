package contexts

import (
	"context"
)

type key int

const (
	orgIDKey key = iota + 1
	instanceIDKey
)

func WithOrgID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, orgIDKey, id)
}

func GetOrgID(ctx context.Context) string {
	instance, ok := ctx.Value(orgIDKey).(string)
	if !ok {
		return ""
	}
	return instance
}

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
