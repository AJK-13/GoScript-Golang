package main

import (
    "fmt"
    "io/ioutil"
    "os"
	"GoScript/src/env"
	"GoScript/src/interpreter"
	"GoScript/src/parseerror"
	"GoScript/src/parser"
	"GoScript/src/runtimeerror"
	"GoScript/src/scanner"
	"GoScript/src/semantic"
	"GoScript/src/semanticerror"
)

var VERSION = "3.9.8";
// var Red    = "\033[31m"
// var Blue   = "\033[34m"
// var White  = "\033[97m"

func check(err error) {
	if err != nil {
		panic(err)
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
func main() {
	if os.Args[1] == "-v" {
		fmt.Println("\033[34m\nGoScript: \033[97mv" + VERSION)
		os.Exit(0)
	} else if os.Args[1] == "-t" {
		runFile("src/Test/Test.gs")
	} else if len(os.Args) == 1 {
		runFile(os.Args[1])
	}
}
