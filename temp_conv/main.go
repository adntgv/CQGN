package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	modePtr := flag.String("mode", "CF", "conversion mode. Example: CF, Celsius to Fahrenheit")
	flag.Parse()
	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments. ")
		flag.Usage()
		return
	}
	val, err := strconv.ParseFloat(flag.Args()[0], 64)
	if err != nil {
		fmt.Println(err)
	}
	var res float64
	switch *modePtr {
	case "CF":
		res = (val * 9 / 5) + 32
	case "CK":
		res = val + 273.15
	case "FC":
		res = (val - 32) * 5 / 9
	case "FK":
		res = (val-32)*5/9 + 273.15
	case "KF":
		res = (val-273.15)*9/5 + 32
	case "KC":
		res = val - 273.15
	}
	fmt.Println(res)
}
