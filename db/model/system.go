package model

import "time"

type System struct {
	Name             string     `db:"name"`
	Location         string     `db:"location"`
	Type             SystemType `db:"type"`
	CleaningInterval int32      `db:"cleaning_interval"`
	LastCleaned      time.Time  `db:"last_cleaned"`
}

type SystemType int32

const (
	Unkown      SystemType = 0
	Glass       SystemType = 1
	Techniplast SystemType = 2
)
