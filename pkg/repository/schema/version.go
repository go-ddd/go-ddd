package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"

	"github.com/galaxyobe/go-ddd/pkg/types"
)

var _ ent.Mixin = (*VersionMixin)(nil)

type VersionMixin struct {
	mixin.Schema
	Field
}

func Version(opts ...FieldOption) VersionMixin {
	obj := VersionMixin{
		Field: Field{
			Name: "version",
		},
	}
	obj.Field.Apply(opts...)
	return obj
}

func (o VersionMixin) Fields() []ent.Field {
	return []ent.Field{
		field.
			String(o.Name).
			StorageKey(o.StorageKey).
			GoType(types.Version("")).
			NotEmpty().
			Match(types.VersionRegexp).
			Annotations(o.Field.Annotations...).
			Comment(o.Comment),
	}
}

func (o VersionMixin) Annotations() []schema.Annotation {
	return o.Field.Annotations
}

var _ ent.Mixin = (*IntVersionMixin)(nil)

type IntVersionMixin struct {
	mixin.Schema
	Field
}

func IntVersion(opts ...FieldOption) IntVersionMixin {
	obj := IntVersionMixin{
		Field: Field{
			Name: "version",
		},
	}
	obj.Field.Apply(opts...)
	return obj
}

func (o IntVersionMixin) Fields() []ent.Field {
	return []ent.Field{
		field.
			Int64(o.Name).
			StorageKey(o.StorageKey).
			GoType(types.IntVersion(0)).
			Positive().
			Annotations(o.Field.Annotations...).
			Comment(o.Comment),
	}
}

func (o IntVersionMixin) Annotations() []schema.Annotation {
	return o.Field.Annotations
}
