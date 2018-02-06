package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func guessGame() bool {
	rand.Seed(int64(time.Now().Nanosecond()))
	num := rand.Intn(100)
	var guess string
	fmt.Println("Guess the number between 0 and 100. [ctrl+c to exit]")
	ans := "n"
	for {
		fmt.Scan(&guess)
		g, err := strconv.Atoi(guess)
		if err != nil {
			fmt.Println(err)
		}
		switch {
		case g == num:
			fmt.Println("You have guessed! Want to play again? [y/N]")
			fmt.Scanln(&ans)
			switch ans {
			case "yes", "y", "Yes":
				return true
			default:
				return false
			}
		case g > num:
			fmt.Println("Too high!")
		case g < num:
			fmt.Println("Too low!")
		default:
			fmt.Println("Try again!")
		}
	}
}

func main() {
	again := true
	for again {
		again = guessGame()
	}
}
