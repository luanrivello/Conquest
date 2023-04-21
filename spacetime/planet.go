package spacetime

import (
	"strconv"

	"conquest/dice"
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

	planet := &Planet{name: "Garden Of Eden", size: radious}

	// totalLongitude = radious/4 central radious
	totalLongitude := radious / 4
	if totalLongitude%2 == 0 {
		totalLongitude++
	}
	planet.longitude = totalLongitude

	center := totalLongitude / 2
	// Main Colum
	tileColum := make([][]tile, totalLongitude)
	planet.tiles = tileColum
	// Central Radious
	tileColum[center] = make([]tile, radious)

	// Create Tiles
	//Middle allways 0
	planet.fillDirections(tileColum, center, 0, tileColum[center], true)

	//Loop afte the middle was filled
	for currentLatitude := 0; currentLatitude < totalLongitude/2; currentLatitude++ {
		//CreatLines
		topRow := make([]tile, radious) //-currentLatitude*2)
		botRow := make([]tile, radious) //-currentLatitude*2)
		tileColum[center-1-currentLatitude] = topRow
		tileColum[center+currentLatitude+totalLongitude%2] = botRow

		planet.fillDirections(tileColum, center, currentLatitude, topRow, true)
		planet.fillDirections(tileColum, center, currentLatitude, botRow, false)

	}

	return planet
}

func (planet *Planet) fillDirections(tileColum [][]tile, center int, currentLatitude int, row []tile, top bool) {
	for currentLongitude := range row {
		tile := &row[currentLongitude]
		tile.SetBuilding("<" + strconv.Itoa(currentLatitude) + " " + strconv.Itoa(currentLongitude) + ">")

		//⬄
		planet.fillLeftRight(currentLongitude, tile, row)

		//⇳
		planet.fillUpDown(currentLongitude, currentLatitude, tile, row, tileColum, center, top)

	}
}

func (planet *Planet) fillLeftRight(currentLongitude int, tile *tile, row []tile) {
	if currentLongitude == 0 {
		// Fist Tile
		tile.left = &row[len(row)-1]
		tile.right = &row[currentLongitude+1]
	} else if currentLongitude == len(row)-1 {
		// Last Tile
		tile.left = &row[currentLongitude-1]
		tile.right = &row[0]
	} else {
		// Between
		tile.left = &row[currentLongitude-1]
		tile.right = &row[currentLongitude+1]
	}
}

func (planet *Planet) fillUpDown(currentLongitude int, currentLatitude int, tile *tile, row []tile, colum [][]tile, center int, top bool) {
	if top {
		tile.down = &colum[center-currentLatitude][currentLongitude]
		tile.down.up = tile
	} else {
		tile.up = &colum[center+currentLatitude][currentLongitude]
		tile.up.down = tile
	}

	if currentLatitude == planet.longitude/2-1 {
		// Last Rows
		// [&][][][*][][][]
		var middle int = len(row) / 2
		if middle+currentLongitude < len(row) {
			if top {
				tile.up = &row[middle+currentLongitude]
				tile.up.up = tile
			} else {
				tile.down = &row[middle+currentLongitude]
				tile.down.down = tile
			}
		}
	}
}

/*
CurrentLongitude = Y
CurrentLatitude = X
Colum = |
*/
/*
//func (planet *Planet) fillUpDown(currentLongitude int, currentLatitude int, tile *tile, row []tile, colum [][]tile, center int, top bool) {
//	// TODORE edge cases literaly
//
//	if currentLatitude == 0 {
//		// Parallel Rows
//		if top {
//			tile.down = &colum[center-currentLatitude][currentLongitude]
//			tile.down.up = tile
//		} else {
//			tile.up = &colum[center+currentLatitude][currentLongitude]
//			tile.up.down = tile
//		}
//	} else if currentLongitude == 0 {
//		//*Fix this shit
//		// Left Edge Tiles of Rest
//		if top {
//			//   [*]
//			//[&][ ]
//			colum[center-currentLatitude][0].up = tile
//
//			//   [*]
//			//[ ][&]
//			tile.down = &colum[center-currentLatitude][1]
//			tile.down.up = tile
//		} else {
//			//[&][ ]
//			//   [*]
//			colum[center+currentLatitude][0].down = tile
//
//			//[ ][&]
//			//   [*]
//			tile.up = &colum[center+currentLatitude][1]
//			tile.up.down = tile
//		}
//	} else if currentLongitude == len(row)-1 {
//		// Right Edge Tiles of Rest
//		if top {
//			//[*]
//			//[ ][&]
//			colum[center-currentLatitude+1][len(row)].up = tile
//
//			//[*]
//			//[&][ ]
//			tile.down = &colum[center-currentLatitude+1][len(row)-1]
//			tile.down.up = tile
//		} else {
//			//[ ][&]
//			//[*]
//			colum[center+currentLatitude][len(row)].down = tile
//
//			//[&][ ]
//			//[*]
//			tile.up = &colum[center+currentLatitude][len(row)-1]
//			tile.up.down = tile
//		}
//	} else {
//		// The Rest
//		if top {
//			tile.down = &colum[center-currentLatitude][currentLongitude+1]
//			tile.down.up = tile
//		} else {
//			tile.up = &colum[center+currentLatitude][currentLongitude+1]
//			tile.up.down = tile
//		}
//	}
//
//	if currentLatitude == planet.longitude/2-1 {
//		// Last Rows
//		// [&][][][*][][][]
//		var middle int = len(row) / 2
//		if middle+currentLongitude < len(row) {
//			if top {
//				tile.up = &row[middle+currentLongitude]
//				tile.up.up = tile
//			} else {
//				tile.down = &row[middle+currentLongitude]
//				tile.down.down = tile
//			}
//		}
//	}
//
//}
*/

// * GET * //
func (planet *Planet) GetName() string {
	return planet.name
}

func (planet *Planet) GetSize() int {
	return planet.size
}

func (planet *Planet) GetLongitude() int {
	return planet.longitude
}

// * Place a sentient in a random position of the array * //
func (planet *Planet) Place(thing PlaceInterface) {
	x := dice.Roll(5)
	y := dice.Roll(5)

	tile := planet.tiles[x][y]
	tile.add(thing)
	thing.Placed(planet, &tile)
}

func (planet *Planet) Move(thing PlaceInterface, relativeX, relativeY int) string {
	var direction string

	if relativeX == 0 && relativeY == 0 {
		return "nowhere"
	}

	// Current Tile
	var targetTile *tile = thing.GetTile()

	//⇳
	if relativeY == 1 {
		targetTile = targetTile.up
		direction = "up"

	} else if relativeY == -1 {
		targetTile = targetTile.down
		direction = "down"
	}

	//⬄
	if relativeX == 1 {
		targetTile = targetTile.right
		direction += "right"

	} else if relativeX == -1 {
		targetTile = targetTile.left
		direction += "left"
	}

	thing.GetTile().remove(thing)
	targetTile.add(thing)
	thing.Placed(planet, targetTile)

	return direction
}

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
