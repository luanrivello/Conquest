package eventlog

import (
	"conquest/creatures"
	"conquest/spacetime"
)

type AkashicRecords struct {
	actions map[string]actionLog
	events  map[string]eventLog
}

type actionLog struct {
	actor    creatures.Creature
	action   string
	target   creatures.Creature
	location spacetime.PlaceInterface
}

type eventLog struct {
	event    string
	location spacetime.PlaceInterface
}
