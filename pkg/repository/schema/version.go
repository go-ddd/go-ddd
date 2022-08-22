package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"

	"github.com/galaxyobe/go-ddd/pkg/types"
)

var _ ent.Mixin = (*Version)(nil)

type Version struct {
	mixin.Schema
	Field
}

func NewVersion(opts ...FieldOption) Version {
	obj := Version{
		Field: Field{
			Name: "version",
		},
	}
	obj.Field.Apply(opts...)
	return obj
}

func (o Version) Fields() []ent.Field {
	return []ent.Field{
		field.
			String(o.Name).
			StorageKey(o.StorageKey).
			NotEmpty().
			GoType(types.Version("")).
			Match(types.VersionRegexp),
	}
}
