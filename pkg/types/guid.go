package types

type GUIDKind int

const (
	StringGUIDKind GUIDKind = iota + 1
	BytesGUIDKind
	Int64GUIDKind
)

type GUID interface {
	Kind() GUIDKind
	Value() any
	IsNull() bool
	Equaled(GUID) bool
}
