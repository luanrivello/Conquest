package spacetime

import (
	"fmt"

	"github.com/luanrivello/conquest/dice"
)

/*
* PLANET STRUCTURE
 */
type Planet struct {
	name      string
	size      int
	longitude int
	tiles     [][]tile
	//?mayor &Mayor
}

// CONSTRUCTOR
func NewPlanet(radious int) *Planet {

	// totalLongitude = radious/4 central radious
	totalLongitude := radious / 4
	if totalLongitude%2 == 0 {
		totalLongitude++
	}

	center := totalLongitude / 2
	// Main Colum
	tileColum := make([][]tile, totalLongitude)
	// Central Radious
	tileColum[center] = make([]tile, radious)

	// Create Tiles
	fillDirections(tileColum, center, 0, tileColum[center], true)
	for currentLatitude := 0; currentLatitude < totalLongitude/2; currentLatitude++ {
		topRow := make([]tile, radious-currentLatitude*2)
		botRow := make([]tile, radious-currentLatitude*2)
		tileColum[center-1-currentLatitude] = topRow
		tileColum[center+currentLatitude+totalLongitude%2] = botRow

		fillDirections(tileColum, center, currentLatitude, topRow, true)
		fillDirections(tileColum, center, currentLatitude, botRow, false)

		/*
			for index, tile := range botColum {
				tile.right = &botColum[index+1]
			}
		*/
	}

	/*
		if radious%2 == 0 {
			tileColum[i] = make([]tile, radious)
			i++
		} else {
			i++
		}

		for j := 0; i+j < radious-1; j++ {
			tileColum[i+j] = make([]tile, radious-(radious/4+j*2)+1)
		}
	*/

	return &Planet{name: "Garden Of Eden", size: radious, longitude: totalLongitude, tiles: tileColum}
}

func fillDirections(tileColum [][]tile, center int, currentLatitude int, colum []tile, top bool) {
	for currentLongitude := range colum {
		tile := &colum[currentLongitude]

		//⬄
		fillLeftRight(currentLongitude, tile, colum)

		//⇳
		fillUpDown(currentLongitude, currentLatitude, tile, colum, tileColum, center, top)

	}
}

func fillLeftRight(currentLongitude int, tile *tile, colum []tile) {
	if currentLongitude == 0 {
		// Fist Tile
		tile.left = &colum[len(colum)-1]
		tile.right = &colum[currentLongitude+1]
	} else if currentLongitude == len(colum)-1 {
		// Last Tile
		tile.left = &colum[currentLongitude-1]
		tile.right = &colum[0]
	} else {
		// Between
		tile.left = &colum[currentLongitude-1]
		tile.right = &colum[currentLongitude+1]
	}
}

func fillUpDown(currentLongitude int, currentLatitude int, tile *tile, colum []tile, tileColum [][]tile, center int, top bool) {
	// TODO last latitudes
	// TODORE edge cases literaly
	if currentLatitude == 0 {
		// Parallel Lines
		if currentLongitude == 0 {
			if top {
				tile.up = &colum[len(colum)-1]
				tile.up.up = &colum[0]
				tile.down = &tileColum[center-currentLatitude][currentLongitude+1]
				tile.down.up = tile
			} else {
				tile.down = &colum[len(colum)-1]
				tile.down.down = &colum[0]
				tile.up = &tileColum[center+currentLatitude][currentLongitude+1]
				tile.up.down = tile
			}
		}

		if top {
			tile.down = &tileColum[center-currentLatitude][currentLongitude]
			tile.down.up = tile
		} else {
			tile.up = &tileColum[center+currentLatitude][currentLongitude]
			tile.up.down = tile
		}

	} else if currentLongitude == 0 {
		// Edge Tiles
		if top {
			tile.up = &colum[len(colum)-1]
			tile.up.up = &colum[0]
			tile.down = &tileColum[center-currentLatitude][currentLongitude+1]
			tile.down.up = tile
		} else {
			tile.down = &colum[len(colum)-1]
			tile.down.down = &colum[0]
			tile.up = &tileColum[center+currentLatitude][currentLongitude+1]
			tile.up.down = tile
		}
	} else {
		// Decresing Lines
		if top {
			tile.down = &tileColum[center-currentLatitude][currentLongitude+1]
			tile.down.up = tile
		} else {
			tile.up = &tileColum[center+currentLatitude][currentLongitude+1]
			tile.up.down = tile
		}
	}

}

// GET
func (planet *Planet) GetName() string {
	return planet.name
}

func (planet *Planet) GetSize() int {
	return planet.size
}

func (planet *Planet) GetLongitude() int {
	return planet.longitude
}

// place a sentient in a random position of the array
func (planet *Planet) Place(thing PlaceInterface) {
	x := dice.Roll(3)
	y := dice.Roll(3)

	thing.Placed(planet, &planet.tiles[x][y])

	planet.tiles[x][y].add(thing)
}

func (planet *Planet) Move(thing PlaceInterface, relativeX, relativeY int) {

	thing.GetTile().remove(thing)

	// Direction
	if relativeX == 1 && relativeY == 1 {
		// UP
		fmt.Print("UP")
		thing.GetTile().up.add(thing)
		thing.Placed(planet, thing.GetTile().up)

	} else if relativeX == 1 && relativeY == -1 {
		// RIGHT
		fmt.Print("RIGHT")
		thing.GetTile().right.add(thing)
		thing.Placed(planet, thing.GetTile().right)

	} else if relativeX == -1 && relativeY == 1 {
		// LEFT
		fmt.Print("LEFT")
		thing.GetTile().left.add(thing)
		thing.Placed(planet, thing.GetTile().left)

	} else {
		// DOWN
		fmt.Print("DOWN")
		thing.GetTile().down.add(thing)
		thing.Placed(planet, thing.GetTile().down)

	}

}

/*
func (planet *Planet) Move(thing PlaceInterface, x, y int) {
	actualX := thing.GetX()
	actualY := thing.GetY()
	planet.tiles[actualX][actualY].remove(thing)

	actualX, actualY = normalize(actualX, actualY, planet.GetSize())
	targetX := actualX + x
	targetY := actualY + y

	if targetX < -(planet.GetSize()/2 + 1) {
		targetX = actualX - x - 1
		targetY = -actualY
	} else if targetX > (planet.GetSize()/2 + 1) {
		targetX = actualX - x + 1
		targetY = -actualY
	}

	if targetY < -(planet.GetSize() / 2) {
		targetY = (planet.GetSize() / 2)
	} else if targetY > (planet.GetSize() / 2) {
		targetY = -(planet.GetSize() / 2)
	}

	fmt.Println("Moving from ", actualX, actualY,
		"\n             ", x, y,
		"\n            ", targetX, targetY)

	targetX, targetY = denormalize(targetX, targetY, planet.size)
	planet.tiles[targetY][targetX].add(thing)
	thing.Placed(planet, targetY, targetX)
}
*/

func (planet *Planet) GetTiles() [][]tile {
	return planet.tiles
}

func (planet *Planet) getTile(x, y int) *tile {
	return &planet.tiles[x][y]
}

func (planet *Planet) Bomb(x, y int) {
	planet.tiles[x][y].exploded()
	planet.tiles[x][y].up.exploded()
	planet.tiles[x][y].down.exploded()
	planet.tiles[x][y].left.exploded()
	planet.tiles[x][y].right.exploded()
}

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
	Move()
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

func (placeble *Placeble) Move() {
	x := dice.Roll(3) - 1
	y := dice.Roll(3) - 1

	if placeble.IsPlaced() && placeble.IsAlive() {
		placeble.GetPlace().Move(placeble, x, y)
	}
}

func (placeble *Placeble) IsAlive() bool     { return true }
func (placeble *Placeble) IsPlaced() bool    { return placeble.isPlaced }
func (placeble *Placeble) GetPlace() *Planet { return placeble.place }
func (placeble *Placeble) Exploded()         {}
func (placeble *Placeble) Rune() rune {
	return '❓'
}
