package tools

import (
	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/vo"
)

func IsGUIDNull(guid vo.GUID) bool {
	if guid == nil {
		return true
	}
	return guid.IsNull()
}
