package board

import (
	"time"

	"conquest/creatures"
	"conquest/eventlog"
	"conquest/spacetime"
)

type Board struct {
	//calendar
	AkashicRecord eventlog.AkashicRecords
	galaxy   spacetime.Galaxy
	calendar int
}

// * Constructor * //
func NewBoard() *Board {
	return &Board{
		galaxy: *spacetime.NewGalaxy(),
	}
}

// * Get * //
func (b *Board) GetGalaxy() spacetime.Galaxy {
	return b.galaxy
}

// * Gameloop * //
func (b *Board) Run() {
	planet := b.galaxy.GetSystem().GetPlanet()
	adam := creatures.NewSentient("Adam", 'M')
	eve := creatures.NewSentient("Eve", 'F')
	snake := creatures.NewCreature("Snake", 'X')

	//planet.Place(&adam)
	//planet.Place(&eve)
	planet.Place(&snake)

	b.loop(adam.Creature, eve.Creature, snake)
}

func (b *Board) loop(adam, eve, snake creatures.Creature) {
	for {
		//adam.Move()
		//eve.Move()
		snake.Move()

		b.calendar += 1
		time.Sleep(2 * time.Second)

		if b.calendar >= 1000 {
			break
		}
	}
}
