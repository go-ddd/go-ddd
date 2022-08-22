package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/vo"
	"github.com/galaxyobe/go-ddd/pkg/repository/schema"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").GoType(vo.EventType("")).Immutable().NotEmpty().Comment("event type"),
		field.String("aggregate_type").GoType(vo.AggregateType("")).Immutable().NotEmpty().Comment("event aggregate type"),
		field.String("org_id").Immutable().NotEmpty().Comment("organisation id"),
		field.String("instance_id").Immutable().NotEmpty().Comment("instance id"),
		field.Bytes("metadata").GoType(vo.Metadata{}).Optional().Immutable().Comment("metadata JSON"),
		field.Bytes("data").Optional().Immutable().Comment("event data JSON"),
		field.Uint64("sequence").Immutable().Comment("event sequence"),
		field.Uint64("previous_aggregate_sequence").Immutable().Comment("previous aggregate sequence"),
		field.Uint64("previous_aggregate_type_sequence").Immutable().Comment("previous aggregate type sequence"),
		field.String("service").Immutable().NotEmpty().Comment("event create service"),
		field.Time("create_time").Immutable().Default(time.Now()),
	}
}

// Mixin of the schema.
func (Event) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schema.NewGUID("id", schema.StringGUIDKind, schema.WithComment("event id")),
		schema.NewGUID("aggregate_id", schema.StringGUIDKind, schema.WithComment("aggregate id")),
		schema.NewVersion(schema.WithComment("aggregate semver version")),
		schema.NewGUID("creator", schema.StringGUIDKind, schema.WithComment("event creator")),
	}
}
