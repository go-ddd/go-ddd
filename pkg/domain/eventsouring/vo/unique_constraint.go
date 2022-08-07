package vo

// UniqueConstraint UniqueCheck represents all information about a unique attribute
type UniqueConstraint struct {
	// UniqueField is the field which should be unique
	UniqueField string
	// UniqueType is the type of the unique field
	UniqueType string
	// InstanceID represents the instance
	InstanceID string
	// Action defines if unique constraint should be added or removed
	Action UniqueConstraintAction
	// ErrorMessage is the message key which should be returned if constraint is violated
	ErrorMessage string
}

type UniqueConstraintAction int32

const (
	UniqueConstraintAdd UniqueConstraintAction = iota
	UniqueConstraintRemove

	uniqueConstraintActionCount
)

func (a UniqueConstraintAction) Valid() bool {
	return a >= 0 && a < uniqueConstraintActionCount
}

func NewAddUniqueConstraint(
	uniqueType,
	uniqueField,
	errMessage string) *UniqueConstraint {
	return &UniqueConstraint{
		UniqueType:   uniqueType,
		UniqueField:  uniqueField,
		ErrorMessage: errMessage,
		Action:       UniqueConstraintAdd,
	}
}

func NewRemoveUniqueConstraint(
	uniqueType,
	uniqueField string) *UniqueConstraint {
	return &UniqueConstraint{
		UniqueType:  uniqueType,
		UniqueField: uniqueField,
		Action:      UniqueConstraintRemove,
	}
}
