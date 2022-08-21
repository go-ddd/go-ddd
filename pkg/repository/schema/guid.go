package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

var _ ent.Mixin = (*GUIDMixin)(nil)

type GUIDMixin struct {
	mixin.Schema
	Field
}

func NewGUIDMixin(name string, opts ...FieldOption) GUIDMixin {
	obj := GUIDMixin{
		Field: Field{
			Name: name,
		},
	}
	obj.Field.Apply(opts...)
	return obj
}

func (m GUIDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").Positive(),
		field.String("name").NotEmpty(),
	}
}
