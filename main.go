package main

import (
	"errors"
	"fmt"
	"os"
)

var (
	originUnit          string
	shouldConvertAgain  string
	errInvalidArguments = errors.New("Invalid Arguments")
	errReadingInput     = errors.New("Error reading input")
)

func main() {
	for {
		fmt.Println("What is the current temperature in" + originUnit + " ? ")

		fmt.Println("Would you like to convert another temperature ? (y/n)")

		if shouldConvertAgain != "Y" {
			fmt.Println("Okay. Bye there!")
			break
		}
	}
}

func printError(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}
