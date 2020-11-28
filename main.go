package main

import (
	"fmt"
)

var (
	originUnit         string
	shouldConvertAgain string
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
