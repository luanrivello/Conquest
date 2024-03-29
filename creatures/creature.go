package creatures

import (
	"fmt"
	"math/rand"
	"time"

	"conquest/dice"
	st "conquest/spacetime"
)

/*
* CREATURE STRUCTURE
 */
type Creature struct {
	//TODO shorterm memory
	//TODO longterm memory
	//TODO objectives memory
	//TODO maslow memory
	name       string
	age        int
	sex        byte
	isAlive    bool
	happines   int
	relashions []relashionship
	actions    []func(*Creature) string

	//* Stats
	Strength     int
	Endurance    int
	Dexterity    int
	Wisdom       int
	Intelligence int
	Charisma     int

	st.Placeble
}

func NewCreature(givenName string, birthSex byte) Creature {
	return Creature{
		name:       givenName,
		age:        0,
		sex:        birthSex,
		isAlive:    true,
		happines:   50,
		relashions: []relashionship{},
		actions: []func(*Creature) string{
			Move,
		},
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

// * Add new relashionship with another sentient
func (creature *Creature) AddRelashionship(other *Creature) {
	creature.relashions = append(creature.relashions, newRelationship(other))
}

func (creature Creature) String() string {
	aux := string(" " + creature.name + " " + string(creature.sex))

	return aux
}

/*
* Actions
 */
func (creature *Creature) TakeAction() string {
	var lenght int = len(creature.actions)
	if lenght == 0 {
		return "Cant do nothing"
	}

	//* Choose a random action from the list of possible actions.
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(lenght)
	action := creature.actions[index]

	//* Call the chosen action function with the current person as its argument.
	return action(creature)
}

func (creature *Creature) AddAction(action func(*Creature) string) {
	creature.actions = append(creature.actions, action)
}

// * Define action functions
func Move(creature *Creature) string {
	x := dice.Roll(3) - 1
	y := dice.Roll(3) - 1

	direction := creature.GetPlace().Move(creature, x, y)
	return fmt.Sprintf("moved %s", direction)
}

func Sleep(creature *Creature) string {
	return "fell asleep"
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

func (p *Creature) Exploded() {

	if rand.Intn(2) == 0 {
		p.isAlive = false
	}

}

func (creature Creature) Rune() rune {
	if !creature.isAlive {
		return 'X'
	} else if creature.sex == 'M' {
		return 'M'
	} else if creature.sex == 'F' {
		return 'F'
	} else {
		return 'S'
	}
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
