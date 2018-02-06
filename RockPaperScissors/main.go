package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	fmt.Println("Welcome to Rock Paper Scissors game! Press [Enter] to finish")

	var userChoice string
	for {
		pi := rand.Int63n(2)
		fmt.Println("Please make your choice:\n1) Rock\n2) Paper\n3) Scissors")
		fmt.Scanf("%s", &userChoice)
		if len(userChoice) == 0 {
			return
		}
		ui, err := strconv.ParseInt(userChoice, 10, 32)
		ui--
		if err != nil {
			fmt.Println(err)
		}
		sum := pi - ui
		result := ""
		switch sum {
		case 0:
			result = "It's a Draw!"
		case 1, -2:
			result = "You lose!"
		case -1, 2:
			result = "You win!"
		}
		fmt.Println(result, "You:", ui+1, "PC:", pi+1)
		userChoice = ""
	}

}

//   0  1  2
//0  0  1  2
//1 -1  0  1
//2 -2 -1  0

//  p r  p  s
//u
//r   d  l  w
//p   w  d  l
//s   l  w  d
