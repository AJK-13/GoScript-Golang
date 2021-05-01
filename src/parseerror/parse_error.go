package parseerror

import (
	"fmt"
	"GoScript/src/token"
	"os"
)


var HadError = false


func LogMessage(line int, message string) {
	report(line, "", message)
	HadError = true
}


func LogError(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", err.Error())
	HadError = true
}


func MakeError(tok token.Token, message string) error {
	if tok.Type == token.EOF {
		return fmt.Errorf("\033[31m[Line \033[97m%v\033[31m]\033[97m Error at end: " + message, tok.Line)
	}
	return fmt.Errorf("\033[31m[Line \033[97m%v\033[31m]\033[97m Error at '" + tok.Lexeme + "': " + message, tok.Line)
}

func report(line int, where string, message string) {
	fmt.Printf("\033[31m[Line \033[97m%v\033[31m]\033[97m Error at '" + where + "': " + message + "\n", line)
}
