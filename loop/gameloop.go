package loop

import (
	"fmt"
	"time"

	"conquest/creatures"
	"conquest/dice"
	"conquest/spacetime"
	"conquest/tui/colors"
)

func Gameloop() {
	var spaceCalendar uint64 = 0
	galaxy := spacetime.NewGalaxy()
	system := galaxy.GetSystem()
	sun := system.GetSun()
	planet := system.GetPlanet()
	adam := creatures.NewSentient("Adam", 'M')
	eve := creatures.NewSentient("Eve", 'F')
	snake := creatures.NewCreature("Snake", 'X')

	planet.Place(&adam)
	planet.Place(&eve)
	planet.Place(&snake)

	for {

		fmt.Println("\033[H\033[2J")
		fmt.Printf("%s✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦\n", colors.BLUE)
		fmt.Printf("%sDate:   %d\n", colors.YELLOW, spaceCalendar)
		fmt.Printf("%sSystem: %s\n", colors.PURPLE, system.GetName())
		fmt.Printf("%sSun:    %s\n", colors.RED, sun.GetName())
		fmt.Printf("%sPlanet: %s\n", colors.GREEN, planet.GetName())
		fmt.Printf("%s✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦\n", colors.BLUE)
		fmt.Println(colors.Reset)

		x := -1
		y := -1
		fmt.Printf("%c\n", '>')
		if dice.Roll(2) == 0 {
			x = dice.Roll(planet.GetLongitude())
			y = dice.Roll(planet.GetSize())
			planet.Bomb(x, y)
			fmt.Printf("%c\n", '💣')
		} else {
			fmt.Println()
		}
		fmt.Printf("%s%c %s\n", colors.RED, '🌎', colors.Reset)
		fmt.Println()

		fmt.Println()

		tiles := planet.GetTiles()

		for i, tileLine := range tiles {

			//fmt.Printf("%d", i+1)
			if i < planet.GetLongitude()/2 {
				for space := 0; space < planet.GetLongitude()/2-i; space++ {
					//fmt.Print("  ")
				}
			} else if i == planet.GetLongitude()/2 {
				//fmt.Print("  ")
			} else {
				for space := 0; space < -planet.GetLongitude()/2+i; space++ {
					//fmt.Print("  ")
				}
			}
			//for space := 0; space < -planet.GetSize()+i*2+8; space++ {
			//	fmt.Print("   ")
			//}
			//make a loop that

			for j, tile := range tileLine {

				var aux = ""
				if i == x && j == y {
					fmt.Print(colors.RED)
				} else if len(tile.String()) != 0 {
					fmt.Print(colors.GREEN)

				}

				aux += "["

				//stack := tile.String()
				//aux += stack
				stack := ""

				// Fill aux string untill it has lenght 9
				for len(aux) < 1+len(stack)/3 {
					aux += " "
				}

				aux += "]"
				fmt.Print(aux + colors.Reset)
			}
			fmt.Println()
		}

		fmt.Println()

		adam.Move()
		eve.Move()
		snake.Move()

		// wait 1 second before next iteration of the loop
		spaceCalendar += 1
		time.Sleep(2 * time.Second)

		if spaceCalendar >= 1000 {
			break
		}
	}

}
