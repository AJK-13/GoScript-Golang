package interpreter

import (
	"fmt"
	"GoScript/src/ast"
	"GoScript/src/env"
	"GoScript/src/semantic"
	"GoScript/src/token"
)

type loxCallable func([]interface{}) (interface{}, error)


type Callable interface {
	Arity() int
	Call([]interface{}) (interface{}, error)
}


type NativeFunction struct {
	Callable
	nativeCall loxCallable
	arity      int
}


func (n *NativeFunction) Call(arguments []interface{}) (interface{}, error) {
	return n.nativeCall(arguments)
}


func (n *NativeFunction) Arity() int {
	return n.arity
}


func (n *NativeFunction) String() string {
	return fmt.Sprintf("<native/%p>", n.nativeCall)
}


type UserFunction struct {
	Callable
	Definition    *ast.Function
	Closure       *env.Environment
	Resolution    semantic.Resolution
	IsInitializer bool
	envSize       int
}


func NewUserFunction(def *ast.Function, closure *env.Environment, res semantic.Resolution, envSize int) *UserFunction {
	return &UserFunction{Definition: def, Closure: closure, Resolution: res, envSize: envSize, IsInitializer: false}
}


func (u *UserFunction) Call(arguments []interface{}) (interface{}, error) {
	env := env.NewSized(u.Closure, u.envSize)

	if !u.Definition.IsProperty() {
		for i, param := range u.Definition.Params {
			env.Define(param.Lexeme, arguments[i], i, "void")
		}
	}

	for _, stmt := range u.Definition.Body {
		_, err := Eval(stmt, env, u.Resolution)

		if err != nil {
			if r, ok := err.(returnError); ok {
				if u.IsInitializer {
					return u.Closure.GetAt(0, token.Token{Lexeme: "this"}, 0)
				}
				return r.value, nil
			}
			return nil, err
		}
	}

	if u.IsInitializer {
		return u.Closure.GetAt(0, token.Token{Lexeme: "this"}, 0)
	}
	return nil, nil
}


func (u *UserFunction) Arity() int {
	return len(u.Definition.Params)
}


func (u *UserFunction) String() string {
	return u.Definition.Name.Lexeme
}


func (u *UserFunction) Bind(instance *ClassInstance) *UserFunction {
	thisEnv := env.NewSized(u.Closure, 1)
	thisEnv.Define("this", instance, 0, "void")
	return &UserFunction{Definition: u.Definition, Closure: thisEnv, Resolution: u.Resolution, envSize: u.envSize, IsInitializer: u.IsInitializer}
}
