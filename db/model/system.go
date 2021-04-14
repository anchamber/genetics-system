package model

import "time"

type System struct {
	Name             string
	Location         string
	Type             SystemType
	CleaningInterval int32
	LastCleaned      time.Time
}

type SystemType int32

const (
	Unkown      SystemType = 0
	Glass       SystemType = 1
	Techniplast SystemType = 2
)
