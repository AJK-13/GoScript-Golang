package runtimeerror

import (
	"GoScript/src/token"
	"fmt"
	"os"
)

func Print(message string) {
	fmt.Fprintf(os.Stderr, "%v\n", message)
	HadError = true
}

func Make(token token.Token, message string) error {
	return fmt.Errorf("\033[31m[Line \033[97m%v\033[31m]: Runtime Error:\033[97m: "+message, token.Line)
}

var HadError = false
