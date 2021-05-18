package db

import (
	"github.com/anchamber/genetics-system/db/model"
)

type TankDB interface {
	Select(Options) ([]*model.Tank, error)
	SelectByNumber(number uint32) (*model.Tank, error)
	Insert(tank *model.Tank) error
	Update(tank *model.Tank) error
	Delete(number uint32) error
}
