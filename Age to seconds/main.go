package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// 21/01/94
	//dd mm yy
	dob := os.Args[1]
	nums := strings.Split(dob, "/")
	var vals [3]int
	for i, v := range nums {
		x, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			fmt.Println(err)
		}
		vals[i] = int(x)
	}

	loc, err := time.LoadLocation("Local")
	if err != nil {
		fmt.Println(err)
	}
	date := time.Date(vals[2], time.Month(vals[1]), vals[0], 0, 0, 0, 0, loc)
	fmt.Println(time.Since(date).Seconds())
}
