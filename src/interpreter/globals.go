  
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
}

func ResetGlobalEnv() {
	GlobalEnv = globals
}