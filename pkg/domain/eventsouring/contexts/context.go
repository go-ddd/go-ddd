package contexts

import (
	"context"

	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/vo"
)

type key int

const (
	orgIDKey key = iota + 1
	instanceIDKey
	serviceKey
	creatorKey
)

func WithOrgID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, orgIDKey, id)
}

func GetOrgID(ctx context.Context) string {
	id, ok := ctx.Value(orgIDKey).(string)
	if !ok {
		return ""
	}
	return id
}

func WithInstanceID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, instanceIDKey, id)
}

func GetInstanceID(ctx context.Context) string {
	id, ok := ctx.Value(instanceIDKey).(string)
	if !ok {
		return ""
	}
	return id
}

func WithService(ctx context.Context, service string) context.Context {
	return context.WithValue(ctx, serviceKey, service)
}

func GetService(ctx context.Context) string {
	service, ok := ctx.Value(serviceKey).(string)
	if !ok {
		return ""
	}
	return service
}

func WithCreator(ctx context.Context, creator vo.GUID) context.Context {
	return context.WithValue(ctx, creatorKey, creator)
}

func GetCreator(ctx context.Context) vo.GUID {
	creator, ok := ctx.Value(creatorKey).(vo.GUID)
	if !ok {
		return nil
	}
	return creator
}
