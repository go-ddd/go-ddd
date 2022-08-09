package types

type GUID interface {
	Value() any
	IsNull() bool
	Equaled(GUID) bool
}
