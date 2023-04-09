package eventlog

import (
	"conquest/creatures"
	"conquest/spacetime"
)

type AkashicRecords struct {
	actions map[string]ActionLog
	events  map[string]EventLog
}

type ActionLog struct {
	actor    creatures.Creature
	action   string
	target   creatures.Creature
	location spacetime.PlaceInterface
}

type EventLog struct {
	event    string
	location spacetime.PlaceInterface
}
