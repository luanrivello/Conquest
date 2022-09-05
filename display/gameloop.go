package display

import (
	"fmt"
	"time"

	"github.com/luanrivello/conquest/creatures"
	"github.com/luanrivello/conquest/dice"
	"github.com/luanrivello/conquest/spacetime"
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
	//planet.Place(&eve)
	//planet.Place(&snake)

	for {

		fmt.Println("\033[H\033[2J")
		fmt.Printf("%s✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦\n", colorBlue)
		fmt.Printf("%sDate:   %d\n", colorYellow, spaceCalendar)
		fmt.Printf("%sSystem: %s\n", colorPurple, system.GetName())
		fmt.Printf("%sSun:    %s\n", colorRed, sun.GetName())
		fmt.Printf("%sPlanet: %s\n", colorGreen, planet.GetName())
		fmt.Printf("%s✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦\n", colorBlue)
		fmt.Println(colorReset)

		x := -1
		y := -1
		fmt.Printf("%c\n", '🛩')
		if dice.Roll(2) == 0 {
			x = dice.Roll(3)
			y = dice.Roll(3)
			planet.Bomb(x, y)
			fmt.Printf("%c\n", '💣')
		} else {
			fmt.Println()
		}
		fmt.Printf("%s%c %s\n", colorRed, '🌎', colorReset)
		fmt.Println()

		adam.Move()
		eve.Move()
		snake.Move()

		fmt.Println()

		tiles := planet.GetTiles()

		for i, tileLine := range tiles {

			//fmt.Printf("%d", i+1)
			if i < planet.GetLongitude()/2 {
				for space := 0; space < planet.GetLongitude()/2-i; space++ {
					fmt.Print("  ")
				}
			} else if i == planet.GetLongitude()/2 {
				fmt.Print("  ")
			} else {
				for space := 0; space < -planet.GetLongitude()/2+i; space++ {
					fmt.Print("  ")
				}
			}
			//for space := 0; space < -planet.GetSize()+i*2+8; space++ {
			//	fmt.Print("   ")
			//}
			//make a loop that

			for j, tile := range tileLine {

				var aux = ""
				if i == x && j == y {
					fmt.Print(colorRed)
				} else if len(tile.String()) != 0 {
					fmt.Print(colorGreen)

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
				fmt.Print(aux + colorReset)
			}
			fmt.Println()
		}

		fmt.Println()

		// wait 1 second before next iteration of the loop
		spaceCalendar += 1
		time.Sleep(2 * time.Second)

		if spaceCalendar >= 1000 {
			break
		}
	}

}
