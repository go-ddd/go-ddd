package schema

type Field struct {
	Name       string
	StorageKey string
	Comment    string
}

type FieldOption func(field *Field)

func (f *Field) Apply(opts ...FieldOption) *Field {
	for _, opt := range opts {
		opt(f)
	}
	return f
}

func (f *Field) SetComment(comment string) *Field {
	f.Comment = comment
	return f
}

func NewField(opts ...FieldOption) Field {
	field := Field{}
	field.Apply(opts...)
	return field
}

func WithComment(comment string) FieldOption {
	return func(field *Field) {
		field.Comment = comment
	}
}
