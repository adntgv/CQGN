package main

import (
	"fmt"
)

func fizzBuzz(fizz, buzz, finish int) error {
	if finish <= 1 {
		return fmt.Errorf("finish must be positive number higher than 1")
	}
	for i := 1; i < finish; i++ {
		switch {
		case i%fizz == 0 && i%buzz == 0:
			fmt.Println("FizzBuzz")
		case i%fizz == 0:
			fmt.Println("Fizz")
		case i%buzz == 0 && i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}

	}
	return nil
}

func main() {
	fizzBuzz(3, 5, 100)
}
