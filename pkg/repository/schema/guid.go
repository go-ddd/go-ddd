package schema

import (
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"

	"github.com/galaxyobe/go-ddd/pkg/types"
)

var _ ent.Mixin = (*GUID)(nil)

const (
	StringGUIDKind = types.StringGUIDKind
	BytesGUIDKind  = types.BytesGUIDKind
	Int64GUIDKind  = types.Int64GUIDKind
)

type GUID struct {
	mixin.Schema
	Field
	Kind types.GUIDKind
}

func NewGUID(name string, kind types.GUIDKind, opts ...FieldOption) GUID {
	obj := GUID{
		Field: Field{
			Name: name,
		},
	}
	obj.Field.Apply(opts...)
	return obj
}

func (o GUID) Fields() []ent.Field {
	switch o.Kind {
	case types.StringGUIDKind:
		return []ent.Field{
			field.String(o.Name).StorageKey(o.StorageKey).Immutable().NotEmpty(),
		}
	case types.BytesGUIDKind:
		return []ent.Field{
			field.Bytes(o.Name).StorageKey(o.StorageKey).Immutable().NotEmpty(),
		}
	case types.Int64GUIDKind:
		return []ent.Field{
			field.Int64(o.Name).StorageKey(o.StorageKey).Immutable().Positive(),
		}
	}
	panic(fmt.Errorf("unexpect guid kind: %s", o.Kind))
}
