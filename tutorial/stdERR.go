package main

import (
	"io"
	"os"
)

func main() {
	myString := ""

	if len(os.Args) == 1 {
		myString = "Please give me one argument!"
	} else {
		myString = os.Args[1]
	}

	io.WriteString(os.Stdout, "This is standard Output\n")
	io.WriteString(os.Stderr, myString)
	io.WriteString(os.Stderr, "\n")

}
