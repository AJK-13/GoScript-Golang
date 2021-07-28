package main

import (
	"GoScript/src/env"
	"GoScript/src/interpreter"
	"GoScript/src/parseerror"
	"GoScript/src/parser"
	"GoScript/src/runtimeerror"
	"GoScript/src/scanner"
	"GoScript/src/semantic"
	"GoScript/src/semanticerror"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

var VERSION = "4.0.0"

func check(err error) {
	if err != nil {
		fmt.Println("\033[31m[Line \033[97m0\033[31m]\033[97m No Such File or Directory")
		os.Exit(0)
	}
}

func runFile(file string) {
	dat, err := ioutil.ReadFile(file)
	check(err)
	fmt.Print("\033[34m\nRunning GoScript, \033[97mVersion: " + VERSION + "\n________________________________\n\n")
	run(string(dat), interpreter.GlobalEnv)
	if parseerror.HadError {
		os.Exit(0)
	} else if runtimeerror.HadError || semanticerror.HadError {
		os.Exit(0)
	}
}

func run(src string, env *env.Environment) {
	scanner := scanner.New(src)
	tokens := scanner.ScanTokens()
	parser := parser.New(tokens)
	statements := parser.Parse()
	if parseerror.HadError {
		return
	}
	resolution, err := semantic.Resolve(statements)
	if err != nil || semanticerror.HadError {
		semanticerror.Print(err.Error())
		return
	}
	interpreter.Interpret(statements, env, resolution)
}
func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	env := interpreter.GlobalEnv
	for {
		fmt.Print("> ")
		dat, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Errorf("\033[31m[Line \033[97m%v\033[31m]\033[97m File Not Found", 0)
		}
		run(string(dat), env)
		parseerror.HadError = false
		runtimeerror.HadError = false
		semanticerror.HadError = false
	}
}

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "-t" {
			runFile("src/Test/Test.gs")
		} else if os.Args[1] == "-v" {
			fmt.Println("\033[34m\nGoScript: \033[97mv" + VERSION)
			os.Exit(0)
		} else {
			runFile(os.Args[1])
		}
	} else {
		runPrompt()
	}
}
