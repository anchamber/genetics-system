package db

import (
	"time"

	"github.com/anchamber/genetics-system/db/model"
)

type SytemDBMock struct{}

var systems []*model.System = []*model.System{
	{Name: "doctor", Location: "tardis", Type: model.Techniplast, CleaningInterval: 90, LastCleaned: time.Now()},
	{Name: "rick", Location: "c-137", Type: model.Techniplast, CleaningInterval: 90, LastCleaned: time.Now()},
	{Name: "morty", Location: "herry-herpson", Type: model.Techniplast, CleaningInterval: 90, LastCleaned: time.Now()},
	{Name: "obi", Location: "high_ground", Type: model.Techniplast, CleaningInterval: 90, LastCleaned: time.Now()},
}

func (db *SytemDBMock) SelctAll() []*model.System {
	return systems
}
