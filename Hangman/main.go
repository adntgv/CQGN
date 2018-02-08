package main

import (
	"fmt"
)

var word = []rune{'s', 'u', 'c', 'c', 'e', 's'}
var figure = []string{
	"/==v==\\ \n",
	"   |\n",
	"   O\n",
	"  /", "|", "\\ \n",
	"  /", " \\ \n"}

func check(s string, arr *[]rune, ch int) (int, bool) {
	r := rune(s[0])
	if contains(word, r) {
		*arr = append(*arr, r)
	} else {
		ch--
	}
	for _, w := range word {
		if !contains(*arr, w) {
			return ch, false
		}
	}
	return ch, true
}

func contains(arr []rune, x rune) bool {
	l := len(arr)
	for i := 0; i < l; i++ {
		if arr[i] == x {
			return true
		}
	}
	return false
}

func printWord(known *[]rune) {
	for _, v := range word {
		if contains(*known, v) {
			fmt.Printf("[%c]", v)
		} else {
			fmt.Print("[_]")
		}
	}
	fmt.Println()
}

func drawPicture(count int) {
	for i := 0; i < count; i++ {
		fmt.Print(figure[i])
	}
	fmt.Println()
}

func main() {
	var known = new([]rune)
	chances := len(figure)
	win := false
	var guess string
	fmt.Println("Guess the word!")
	for chances > 0 {
		// Draw picture orz
		drawPicture(len(figure) - chances)
		// Draw letters line [_] [_] [_] [_] [_]..
		printWord(known)
		// Ask letter
		fmt.Scanln(&guess)
		// test existance
		chances, win = check(guess, known, chances)
		if win {
			fmt.Println("You won!")
			printWord(known)
			return
		}
	}
	fmt.Println("You have lost!")
	drawPicture(0)
	return
}
