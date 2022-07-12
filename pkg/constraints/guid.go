package constraints

type GUID interface {
	int64 | string | []byte
}
