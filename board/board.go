package board

import (
	"time"

	"conquest/creatures"
	"conquest/eventlog"
	"conquest/spacetime"
)

type Board struct {
	AkashicRecord eventlog.AkashicRecords
	galaxy        spacetime.Galaxy
	Calendar      int
}

// * Constructor * //
func NewBoard() *Board {
	return &Board{
		AkashicRecord: eventlog.AkashicRecords{},
		galaxy:        *spacetime.NewGalaxy(),
		Calendar:      0,
	}
}

// * Get * //
func (b *Board) GetGalaxy() spacetime.Galaxy {
	return b.galaxy
}

// * Gameloop * //
func (b *Board) Run() {
	planet := b.galaxy.GetSystem().GetPlanet()
	adam := creatures.NewCreature("Adam", 'M')
	eve := creatures.NewCreature("Eve", 'F')
	snake := creatures.NewCreature("Snake", 'X')

	planet.Place(&adam)
	planet.Place(&eve)
	planet.Place(&snake)

	b.Calendar += 1
	time.Sleep(2 * time.Second)

	b.loop(&adam, &eve, &snake)
}

func (b *Board) loop(creatures ...*creatures.Creature) {
	for {
		for i := range creatures {
			action := creatures[i].TakeAction()
			b.AkashicRecord.AddAction(creatures[i], action, creatures[i].GetPlace(), b.Calendar)
		}

		b.Calendar += 1
		time.Sleep(2 * time.Second)

		if b.Calendar >= 100 {
			break
		}
	}
}
