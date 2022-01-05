package main

import (
	"io"
	"os"
)

func main() {
	myString := ""
	arguments := os.Args

	if len(arguments) == 1 {
		myString = "Please give me one argument!"
	} else {
		myString = arguments[2]
	}

	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")

}
