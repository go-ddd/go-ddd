package schema

type Field struct {
	Name       string
	StorageKey string
}

type FieldOption func(field *Field)

func (f *Field) Apply(opts ...FieldOption) {
	for _, opt := range opts {
		opt(f)
	}
}

func NewField(opts ...FieldOption) Field {
	field := Field{}
	field.Apply(opts...)
	return field
}
