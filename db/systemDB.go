package db

import (
	"github.com/anchamber/genetics-system/db/model"
)

type SystemDB interface {
	Select(Options) ([]*model.System, error)
	SelectByName(name string) (*model.System, error)
	Insert(system *model.System) error
	Update(system *model.System) error
	Delete(name string) error
}
