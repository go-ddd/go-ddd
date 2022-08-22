// Code generated by ent, DO NOT EDIT.

package do

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/galaxyobe/go-ddd/pkg/domain/database/do"
	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/repository/do/event"
	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/vo"
	"github.com/galaxyobe/go-ddd/pkg/types"
)

// Event is the model entity for the Event schema.
type Event struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// AggregateID holds the value of the "aggregate_id" field.
	AggregateID string `json:"aggregate_id,omitempty"`
	// Version holds the value of the "version" field.
	Version types.Version `json:"version,omitempty"`
	// Creator holds the value of the "creator" field.
	Creator string `json:"creator,omitempty"`
	// event type
	Type vo.EventType `json:"type,omitempty"`
	// event aggregate type
	AggregateType vo.AggregateType `json:"aggregate_type,omitempty"`
	// organisation id
	OrgID string `json:"org_id,omitempty"`
	// instance id
	InstanceID string `json:"instance_id,omitempty"`
	// metadata JSON
	Metadata do.StringMap `json:"metadata,omitempty"`
	// event data JSON
	Data []byte `json:"data,omitempty"`
	// event sequence
	Sequence uint64 `json:"sequence,omitempty"`
	// previous aggregate sequence
	PreviousAggregateSequence uint64 `json:"previous_aggregate_sequence,omitempty"`
	// previous aggregate type sequence
	PreviousAggregateTypeSequence uint64 `json:"previous_aggregate_type_sequence,omitempty"`
	// event create service
	Service string `json:"service,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Event) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case event.FieldData:
			values[i] = new([]byte)
		case event.FieldMetadata:
			values[i] = new(do.StringMap)
		case event.FieldSequence, event.FieldPreviousAggregateSequence, event.FieldPreviousAggregateTypeSequence:
			values[i] = new(sql.NullInt64)
		case event.FieldID, event.FieldAggregateID, event.FieldCreator, event.FieldType, event.FieldAggregateType, event.FieldOrgID, event.FieldInstanceID, event.FieldService:
			values[i] = new(sql.NullString)
		case event.FieldCreateTime:
			values[i] = new(sql.NullTime)
		case event.FieldVersion:
			values[i] = new(types.Version)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Event", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Event fields.
func (e *Event) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case event.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				e.ID = value.String
			}
		case event.FieldAggregateID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field aggregate_id", values[i])
			} else if value.Valid {
				e.AggregateID = value.String
			}
		case event.FieldVersion:
			if value, ok := values[i].(*types.Version); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value != nil {
				e.Version = *value
			}
		case event.FieldCreator:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field creator", values[i])
			} else if value.Valid {
				e.Creator = value.String
			}
		case event.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				e.Type = vo.EventType(value.String)
			}
		case event.FieldAggregateType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field aggregate_type", values[i])
			} else if value.Valid {
				e.AggregateType = vo.AggregateType(value.String)
			}
		case event.FieldOrgID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field org_id", values[i])
			} else if value.Valid {
				e.OrgID = value.String
			}
		case event.FieldInstanceID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field instance_id", values[i])
			} else if value.Valid {
				e.InstanceID = value.String
			}
		case event.FieldMetadata:
			if value, ok := values[i].(*do.StringMap); !ok {
				return fmt.Errorf("unexpected type %T for field metadata", values[i])
			} else if value != nil {
				e.Metadata = *value
			}
		case event.FieldData:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field data", values[i])
			} else if value != nil {
				e.Data = *value
			}
		case event.FieldSequence:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sequence", values[i])
			} else if value.Valid {
				e.Sequence = uint64(value.Int64)
			}
		case event.FieldPreviousAggregateSequence:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field previous_aggregate_sequence", values[i])
			} else if value.Valid {
				e.PreviousAggregateSequence = uint64(value.Int64)
			}
		case event.FieldPreviousAggregateTypeSequence:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field previous_aggregate_type_sequence", values[i])
			} else if value.Valid {
				e.PreviousAggregateTypeSequence = uint64(value.Int64)
			}
		case event.FieldService:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field service", values[i])
			} else if value.Valid {
				e.Service = value.String
			}
		case event.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				e.CreateTime = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Event.
// Note that you need to call Event.Unwrap() before calling this method if this Event
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Event) Update() *EventUpdateOne {
	return (&EventClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the Event entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Event) Unwrap() *Event {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("do: Event is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Event) String() string {
	var builder strings.Builder
	builder.WriteString("Event(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("aggregate_id=")
	builder.WriteString(e.AggregateID)
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", e.Version))
	builder.WriteString(", ")
	builder.WriteString("creator=")
	builder.WriteString(e.Creator)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", e.Type))
	builder.WriteString(", ")
	builder.WriteString("aggregate_type=")
	builder.WriteString(fmt.Sprintf("%v", e.AggregateType))
	builder.WriteString(", ")
	builder.WriteString("org_id=")
	builder.WriteString(e.OrgID)
	builder.WriteString(", ")
	builder.WriteString("instance_id=")
	builder.WriteString(e.InstanceID)
	builder.WriteString(", ")
	builder.WriteString("metadata=")
	builder.WriteString(fmt.Sprintf("%v", e.Metadata))
	builder.WriteString(", ")
	builder.WriteString("data=")
	builder.WriteString(fmt.Sprintf("%v", e.Data))
	builder.WriteString(", ")
	builder.WriteString("sequence=")
	builder.WriteString(fmt.Sprintf("%v", e.Sequence))
	builder.WriteString(", ")
	builder.WriteString("previous_aggregate_sequence=")
	builder.WriteString(fmt.Sprintf("%v", e.PreviousAggregateSequence))
	builder.WriteString(", ")
	builder.WriteString("previous_aggregate_type_sequence=")
	builder.WriteString(fmt.Sprintf("%v", e.PreviousAggregateTypeSequence))
	builder.WriteString(", ")
	builder.WriteString("service=")
	builder.WriteString(e.Service)
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(e.CreateTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Events is a parsable slice of Event.
type Events []*Event

func (e Events) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}
