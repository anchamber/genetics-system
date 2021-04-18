package db

import (
	"github.com/anchamber/genetics-system/db/model"
)

type SystemDB interface {
	SelectAll() ([]*model.System, error)
	SelectByName(name string) (*model.System, error)
	Insert(system *model.System) error
	Update(system *model.System) error
}

type ErrorCode string

const (
	SystemAlreadyExists ErrorCode = "system already exists"
	Unknown             ErrorCode = "unknown error with db occured"
)
