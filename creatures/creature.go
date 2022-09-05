package creatures

import (
	"math/rand"

	"github.com/luanrivello/conquest/dice"
	st "github.com/luanrivello/conquest/spacetime"
)

/*
* CREATURE STRUCTURE
 */
type Creature struct {
	name       string
	age        int
	sex        byte
	isAlive    bool
	relashions []relashionship

	st.Placeble
}

func NewCreature(givenName string, birthSex byte) Creature {
	return Creature{
		name:       givenName,
		age:        0,
		sex:        birthSex,
		isAlive:    true,
		relashions: []relashionship{},
	}
}

func Birth(name string) Creature {

	var birthSex byte
	if dice.Roll(2) == 0 {
		birthSex = 'M'
	} else {
		birthSex = 'F'
	}

	return NewCreature(name, birthSex)
}

// add new relashionship with another sentient
func (creature *Creature) AddRelashionship(other *Creature) {
	creature.relashions = append(creature.relashions, newRelationship(other))
}

func (creature Creature) String() string {
	aux := string(creature.Rune()) + " " + creature.name + " " + string(creature.sex)

	return aux
}

/*
* PLACEBLE INTERFACE
 */
func (creature *Creature) IsAlive() bool {
	return creature.isAlive
}

func (creature *Creature) IsPlaced() bool {
	return creature.Placeble.IsPlaced()
}

func (creature *Creature) Move() {
	x := dice.Roll(3) - 1
	y := dice.Roll(3) - 1

	if creature.IsPlaced() && creature.IsAlive() {
		creature.GetPlace().Move(creature, x, y)
	}
}

func (p *Creature) Exploded() {

	if rand.Intn(2) == 0 {
		p.isAlive = false
	}

}

func (creature *Creature) Rune() rune {
	if !creature.isAlive {
		return '💀'
	}

	return '🐍'
}

/*
* RELASHIONSHIP STRUCTURE
 */
type relashionship struct {
	creature *Creature
	value    int
}

func newRelationship(creature *Creature) relashionship {
	return relashionship{creature, 0}
}
