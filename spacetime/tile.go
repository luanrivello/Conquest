package spacetime

import "fmt"

/*
* TILE STRUCTURE
 */
type tile struct {
	up       *tile
	down     *tile
	right    *tile
	left     *tile
	terrain  string
	building string
	stack    []PlaceInterface
}

func (t *tile) newTile() {
	t.terrain = "grass"
	t.building = "none"
	t.stack = make([]PlaceInterface, 0)
}

func (t *tile) add(p PlaceInterface) {
	t.stack = append(t.stack, p)
}

func (t *tile) remove(p PlaceInterface) {
	for i, v := range t.stack {
		if v == p {
			t.stack[i] = t.stack[len(t.stack)-1]
			t.stack = t.stack[:len(t.stack)-1]
			return
		}
	}
}

func (t *tile) String() string {
	var aux string

	if len(t.stack) > 0 {
		for _, placeble := range t.stack {
			aux += fmt.Sprintf("%q", placeble.Rune())
		}
	}
	return aux
}

func (t *tile) exploded() {
	for i := range t.stack {
		t.stack[i].Exploded()
	}
}

func (t *tile) SetBuilding(building string) {
	t.building = building
}
