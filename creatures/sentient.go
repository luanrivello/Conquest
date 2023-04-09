package creatures

import (
	"conquest/creatures/mind/personality"
	"conquest/dice"
)

/*
* SENTIENT CREATURE STRUCTURE
 */
type Sentient struct {
	Creature

	personality *personality.Personality
}

func NewSentient(givenName string, birthSex byte) Sentient {
	isfp := &personality.ISFP

	return Sentient{
		Creature:    NewCreature(givenName, birthSex),
		personality: isfp,
	}
}

/*
* PLACEBLE INTERFACE
 */
func (sentient *Sentient) Move() {
	x := dice.Roll(3) - 1
	y := dice.Roll(3) - 1

	if sentient.IsPlaced() && sentient.IsAlive() {
		sentient.GetPlace().Move(sentient, x, y)
	}
}

func (sentient Sentient) Rune() rune {
	if !sentient.isAlive {
		return 'X'
	} else if sentient.sex == 'M' {
		return 'M'
	} else if sentient.sex == 'F' {
		return 'F'
	} else {
		return 'S'
	}
}

// ToString
func (s Sentient) String() string {
	aux := string(s.Rune()) + " " + s.name + " " + string(s.sex)
	return aux
}

// GET
func (s *Sentient) Age() int {
	return s.age
}

func (s *Sentient) GetName() string {
	return s.name
}

func (s *Sentient) GetAge() int {
	return s.age
}

// SET
func (s *Sentient) SetName(name string) {
	s.name = name
}

func (s *Sentient) SetAge(age int) {
	s.age = age
}
