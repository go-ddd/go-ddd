package do

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type StringArray []string

func (s *StringArray) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *StringArray) Scan(value any) error {
	data, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("expect []byte")
	}
	return json.Unmarshal(data, s)
}
