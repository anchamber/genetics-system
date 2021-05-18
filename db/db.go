package db

import (
	"fmt"
	apiModel "github.com/anchamber/genetics-api/model"
)

type ErrorCode string

type Options struct {
	Pageination *apiModel.Pageination
	Filters     []*apiModel.Filter
}

type EntityAlreadyExists struct {
	entity string
}

func (e *EntityAlreadyExists) Error() string {
	return fmt.Sprintf("%s already exists", e.entity)
}

type UnknownDBError struct {
	message string
}

func (e *UnknownDBError) Error() string {
	return e.message
}
