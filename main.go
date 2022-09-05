package main

import (
	"fmt"
	"strings"

	"github.com/luanrivello/conquest/display"
)

func main() {
	bars := strings.Repeat("✦", 20)
	fmt.Println(bars + " START " + bars)

	display.Menu()

	fmt.Println(bars + "  END  " + bars)
}
