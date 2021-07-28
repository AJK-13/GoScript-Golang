package semanticerror

import (
	"fmt"
	"os"
)

func Print(message string) {
	fmt.Fprintf(os.Stderr, "%v\n", message)
	HadError = true
}

func Make(message string) error {
	return fmt.Errorf("%s", message)
}

var HadError = false
