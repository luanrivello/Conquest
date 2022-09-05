package dice

import (
	"math/rand"
	"time"
)

func newSeed() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func Roll(sides int) int {
	newSeed()
	return rand.Intn(sides)
}

func RollPlusOne(sides int) int {
	newSeed()
	return rand.Intn(sides) + 1
}