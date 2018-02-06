package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const vowels string = "aeiouy"
const consonants string = "bcdfghijklmpqrstvwxz"

func simpleRandomName(length int) string {
	rand.Seed(int64(time.Now().Nanosecond()))
	name := ""
	letters := vowels + consonants
	first := true
	for i := 0; i < length; i++ {
		name = name + string(letters[rand.Intn(len(letters))])
		if first {
			name = strings.ToUpper(name)
			first = false
		}
	}
	return name
}

func main() {
	fmt.Println(simpleRandomName(5))
}
