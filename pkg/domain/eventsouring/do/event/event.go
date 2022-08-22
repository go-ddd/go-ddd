// Code generated by ent, DO NOT EDIT.

package event

const (
	// Label holds the string label denoting the event type in the database.
	Label = "event"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "nickname"
	// Table holds the table name of the event in the database.
	Table = "events"
)

// Columns holds all SQL columns for event fields.
var Columns = []string{
	FieldID,
	FieldAge,
	FieldName,
	FieldNickname,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
