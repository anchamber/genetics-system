package db

import (
	"github.com/anchamber/genetics-system/db/model"
)

type SystemDB interface {
	SelctAll() []*model.System
}
