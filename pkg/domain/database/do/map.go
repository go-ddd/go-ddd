package do

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
)

// StringMap string map data
type StringMap map[string]any

func (m *StringMap) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("want to []byte type, but got type: %s", reflect.TypeOf(value))
	}
	return json.Unmarshal(bytes, m)
}

func (m *StringMap) Value() (driver.Value, error) {
	if m == nil {
		return nil, nil
	}
	return json.Marshal(m)
}
