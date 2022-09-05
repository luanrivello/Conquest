package display

import (
	////"bufio"
	"fmt"
)

//Cores do Print
var colorReset = "\033[0m"
var colorRed = "\033[31m"
var colorGreen = "\033[32m"
var colorYellow = "\033[33m"
var colorBlue = "\033[34m"
var colorPurple = "\033[35m"
var colorCyan = "\033[36m"

func Menu() {
	for {
		fmt.Println("\033[H\033[2J")

		fmt.Println(colorYellow, "1- Explore")
		fmt.Println(colorGreen, "2- Conquest")
		fmt.Println(colorPurple, "3- God")
		fmt.Println(colorCyan, "4- Dev")
		fmt.Println(colorRed, "0- Exit")
		fmt.Println(colorReset)

		fmt.Printf(" => ")
		/*
		//scanner := bufio.NewScanner(os.Stdin)
		//scanner.Scan()
		//resp, err := strconv.ParseInt(scanner.Text(), 10, 64)
		//if err != nil {
		//	resp = -1
		//}
		*/
		resp := 4

		fmt.Println("")

		switch resp {
		case 0:
			fmt.Println("Exiting...")
			return

		case 1:
			fmt.Println("Explore Choosen")

		case 2:
			fmt.Println("Conquest Choosen")

		case 3:
			fmt.Println("God Choosen")

		case 4:
			fmt.Println("Dev Choosen")
			devmode()

		default:
			fmt.Println("Invalid")
			continue
		}

	}
}

func devmode() {

	Gameloop()

}

/*
//func explore() {
//
//}
//
//func conquest() {
//
//}
//
//func godmode() {
//
//}
*/