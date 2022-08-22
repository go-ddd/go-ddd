package types

import (
	"fmt"
)

type GUIDKind int

const (
	StringGUIDKind GUIDKind = iota + 1
	BytesGUIDKind
	Int64GUIDKind
)

func (k GUIDKind) String() string {
	switch k {
	case StringGUIDKind:
		return "string"
	case BytesGUIDKind:
		return "bytes"
	case Int64GUIDKind:
		return "int64"
	default:
		return fmt.Sprintf("GUIDKind(%d)", k)
	}
}

type GUID interface {
	Kind() GUIDKind
	Value() any
	IsNull() bool
	Equaled(GUID) bool
}
