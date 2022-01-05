package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("give me one or more floats.")
		os.Exit(1)
	}

	argument := os.Args
	min, _ := strconv.ParseFloat(argument[1], 64)
	max, _ := strconv.ParseFloat(argument[1], 64)

	for i := 2; i < len(argument); i++ {
		temp, _ := strconv.ParseFloat(argument[i], 64)
		if temp > max {
			max = temp
		}
		if temp < min {
			min = temp
		}
	}

	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
}
