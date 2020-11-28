package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"temp-converter-cli/helpers"
)

var (
	err                 error
	originValue         float64
	originUnit          string
	shouldConvertAgain  string
	errInvalidArguments = errors.New("Invalid Arguments")
	errReadingInput     = errors.New("Error reading input")
)

func main() {
	if len(os.Args) != 2 {
		helpers.PrintError(errInvalidArguments)
	}

	originUnit = strings.ToUpper(os.Args[1])

	for {
		fmt.Println("What is the current temperature in " + originUnit + " ? ")
		_, err = fmt.Scanln(&originValue)
		if err != nil {
			helpers.PrintError(errReadingInput)
		}

		if originUnit == "C" {
			helpers.ConvertToFahrenheit(originValue)
		} else {
			helpers.ConvertToCelsius(originValue)
		}

		fmt.Println("Would you like to convert another temperature ? (y/n)")
		_, err = fmt.Scanln(&shouldConvertAgain)
		if err != nil {
			helpers.PrintError(errReadingInput)
		}

		shouldConvertAgain = strings.ToUpper(strings.TrimSpace(shouldConvertAgain))

		if shouldConvertAgain != "Y" {
			fmt.Println("Okay. Bye there!")
			break
		}
	}
}
