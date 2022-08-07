package vo

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
)

// Metadata meta data
type Metadata map[string]any

func (m *Metadata) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("want to []byte type, but got type: %s", reflect.TypeOf(value))
	}
	return json.Unmarshal(bytes, m)
}

func (m *Metadata) Value() (driver.Value, error) {
	if m == nil {
		return nil, nil
	}
	return json.Marshal(m)
}
