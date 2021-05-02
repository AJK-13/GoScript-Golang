  
package interpreter

import (
	"GoScript/src/env"
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
	// GlobalEnv.Define("Number", &NativeFunction{
	// 	arity: 1,
	// 	nativeCall: func(args []interface{}) (interface{}, error) {
	// 		// fmt.Println(strconv.ParseInt("10", 10, 64))
	// 		return strconv.ParseInt("10", 10, 64)
	// 	},
	// }, -1, "void")
}

func ResetGlobalEnv() {
	GlobalEnv = globals
}