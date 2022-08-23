package schema

import (
	"time"

	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/vo"
	reposchema "github.com/galaxyobe/go-ddd/pkg/repository/schema"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").GoType(vo.EventType("")).Immutable().NotEmpty().Annotations(entproto.Field(7)).Comment("event type"),
		field.String("aggregate_type").GoType(vo.AggregateType("")).Immutable().NotEmpty().Annotations(entproto.Field(8)).Comment("event aggregate type"),
		field.Bytes("metadata").GoType(vo.Metadata{}).Optional().Immutable().Annotations(entproto.Field(9)).Comment("metadata JSON"),
		field.Bytes("data").Optional().Immutable().Annotations(entproto.Field(10)).Comment("event data JSON"),
		field.Uint64("sequence").Immutable().Annotations(entproto.Field(11)).Comment("event sequence"),
		field.Uint64("previous_aggregate_sequence").Immutable().Annotations(entproto.Field(12)).Comment("previous aggregate sequence"),
		field.Uint64("previous_aggregate_type_sequence").Immutable().Annotations(entproto.Field(13)).Comment("previous aggregate type sequence"),
		field.String("service").Immutable().NotEmpty().Annotations(entproto.Field(14)).Comment("event create service"),
		field.Time("create_time").Immutable().Default(time.Now()).Annotations(entproto.Field(15)).Comment("create event time"),
	}
}

// Mixin of the reposchema.
func (Event) Mixin() []ent.Mixin {
	return []ent.Mixin{
		reposchema.UUID("id", reposchema.WithAnnotations(entproto.Field(1)), reposchema.WithComment("event id")),
		reposchema.UUID("aggregate_id", reposchema.WithAnnotations(entproto.Field(2)), reposchema.WithComment("aggregate id")),
		reposchema.UUID("org_id", reposchema.WithAnnotations(entproto.Field(3)), reposchema.WithComment("organisation id")),
		reposchema.UUID("instance_id", reposchema.WithAnnotations(entproto.Field(4)), reposchema.WithComment("instance id")),
		reposchema.Version(reposchema.WithAnnotations(entproto.Field(5)), reposchema.WithComment("aggregate semver version")),
		reposchema.UUID("creator", reposchema.WithAnnotations(entproto.Field(6)), reposchema.WithComment("event creator")),
	}
}

func (Event) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(
			entproto.PackageName("io.entgo.apps.todo"),
		),
		entproto.Service(),
	}
}
