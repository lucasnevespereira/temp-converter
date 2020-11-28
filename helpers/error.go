package helpers

import (
	"fmt"
	"os"
)

func PrintError(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}
