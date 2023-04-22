package eventlog

import (
	"conquest/creatures"
	st "conquest/spacetime"
	"fmt"
)

type AkashicRecords struct {
	LogLines        []logInterface
	Actions         []actionLog
	actionsToTarget []actionToTargetLog
	events          []eventLog
}

type logInterface interface {
	String() string
}

type actionLog struct {
	Actor    *creatures.Creature
	Action   string
	Location *st.Planet
	Time     int
}

func (log actionLog) String() string {
	return fmt.Sprintf("[time: %04d] %s %s at %s", log.Time, log.Actor.String(), log.Action, log.Location.GetName())
}

func (records *AkashicRecords) AddAction(actor *creatures.Creature, action string, location *st.Planet, time int) {
	entry := actionLog{
		Actor:    actor,
		Action:   action,
		Location: location,
		Time:     time,
	}
	records.Actions = append(records.Actions, entry)
}

type actionToTargetLog struct {
	actor    creatures.Creature
	action   string
	target   creatures.Creature
	location st.Planet
	time     int
}

type eventLog struct {
	event    string
	location st.Planet
	time     int
}
