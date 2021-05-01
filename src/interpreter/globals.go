  
package interpreter

import (
	"GoScript/src/env"
	// "bufio"
	// "os"
	// "fmt"
)

var GlobalEnv = env.NewGlobal()
var globals = GlobalEnv
func check(err error) {
	if err != nil {
		panic(err)
	}
}
func init() {
	// reader := bufio.NewReader(os.Stdin)
	// GlobalEnv.Define("ask", &NativeFunction{
	// 	arity: 0,
	// 	nativeCall: func(args []interface{}) (interface{}, error) {
	// 		for {
	// 			fmt.Print("> ")
	// 			dat, err := reader.ReadBytes('\n')
	// 			check(err)
	// 			return dat, err
	// 		}
	// 	},
	// }, -1, "void")
}

func ResetGlobalEnv() {
	GlobalEnv = globals
}