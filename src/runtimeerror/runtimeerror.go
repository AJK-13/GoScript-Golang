package runtimeerror

import (
	"fmt"
	"GoScript/src/token"
	"os"
)


func Print(message string) {
	fmt.Fprintf(os.Stderr, "%v\n", message)
	HadError = true
}

func Make(token token.Token, message string) error {
	return fmt.Errorf("\033[31m[Line \033[97m%v\033[31m]\033[97m: " + message, token.Line)
}


var HadError = false
