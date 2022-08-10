package interfaces

import (
	"github.com/galaxyobe/go-ddd/pkg/domain/database/vo"
)

type IStatement interface {
	Create(options vo.ExecOptions) string
	Upsert(options vo.ExecOptions) string
	Update(options vo.ExecOptions) string
	Delete(options vo.ExecOptions) string
}
