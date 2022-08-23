package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"

	"github.com/galaxyobe/go-ddd/pkg/types"
)

var _ ent.Mixin = (*UUIDMixin)(nil)

type UUIDMixin struct {
	mixin.Schema
	Field
}

func UUID(name string, opts ...FieldOption) UUIDMixin {
	obj := UUIDMixin{
		Field: Field{
			Name: name,
		},
	}
	obj.Field.Apply(opts...)
	return obj
}

func (o UUIDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.
			String(o.Name).
			StorageKey(o.StorageKey).
			GoType(types.UUID("")).
			NotEmpty().
			Annotations(o.Field.Annotations...).
			Comment(o.Comment),
	}
}

func (o UUIDMixin) Annotations() []schema.Annotation {
	return o.Field.Annotations
}

var _ ent.Mixin = (*IntIDMixin)(nil)

type IntIDMixin struct {
	mixin.Schema
	Field
}

func IntID(name string, opts ...FieldOption) IntIDMixin {
	obj := IntIDMixin{
		Field: Field{
			Name: name,
		},
	}
	obj.Field.Apply(opts...)
	return obj
}

func (o IntIDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.
			Int64(o.Name).
			StorageKey(o.StorageKey).
			GoType(types.IntID(0)).
			Positive().
			Annotations(o.Field.Annotations...).
			Comment(o.Comment),
	}
}

func (o IntIDMixin) Annotations() []schema.Annotation {
	return o.Field.Annotations
}

var _ ent.Mixin = (*Int32IDMixin)(nil)

type Int32IDMixin struct {
	mixin.Schema
	Field
}

func Int32ID(name string, opts ...FieldOption) Int32IDMixin {
	obj := Int32IDMixin{
		Field: Field{
			Name: name,
		},
	}
	obj.Field.Apply(opts...)
	return obj
}

func (o Int32IDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.
			Int32(o.Name).
			StorageKey(o.StorageKey).
			GoType(types.Int32ID(0)).
			Positive().
			Annotations(o.Field.Annotations...).
			Comment(o.Comment),
	}
}

func (o Int32IDMixin) Annotations() []schema.Annotation {
	return o.Field.Annotations
}
