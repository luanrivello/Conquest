package spacetime

import (
	"conquest/dice"
)

/*
* PLACEBLE THINGS
 */
type Placeble struct {
	isPlaced bool
	place    *Planet
	tile     *tile
	x        int
	y        int
}

type PlaceInterface interface {
	GetX() int
	GetY() int
	Placed(planet *Planet, tile *tile)
	IsAlive() bool
	IsPlaced() bool
	GetPlace() *Planet
	GetTile() *tile
	Move() (int, int)
	Rune() rune
	Exploded()
}

// GET
func (placeble *Placeble) GetTile() *tile {
	return placeble.tile
}

func (placeble *Placeble) GetX() int {
	return placeble.x
}

func (placeble *Placeble) GetY() int {
	return placeble.y
}

func (placeble *Placeble) Placed(planet *Planet, tile *tile) {
	placeble.isPlaced = true
	placeble.place = planet
	placeble.tile = tile
}

func (placeble *Placeble) Move() (int, int) {
	x := dice.Roll(3) - 1
	y := dice.Roll(3) - 1

	if placeble.IsPlaced() && placeble.IsAlive() {
		placeble.GetPlace().Move(placeble, x, y)
	}

	return x, y
}

func (placeble *Placeble) IsAlive() bool     { return true }
func (placeble *Placeble) IsPlaced() bool    { return placeble.isPlaced }
func (placeble *Placeble) GetPlace() *Planet { return placeble.place }
func (placeble *Placeble) Exploded()         {}
func (placeble *Placeble) Rune() rune {
	return '❓'
}
