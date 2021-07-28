package env

import (
	"GoScript/src/runtimeerror"
	"GoScript/src/token"
	"fmt"
	"os"
)

type uninitialized struct{}

var needsInitialization = &uninitialized{}

type Environment struct {
	values        map[string]interface{}
	enclosing     *Environment
	indexedValues []interface{}
}

func New(env *Environment) *Environment {
	return NewSized(env, 0)
}

func NewSized(env *Environment, size int) *Environment {
	return &Environment{values: make(map[string]interface{}), enclosing: env, indexedValues: make([]interface{}, size)}
}

func NewGlobal() *Environment {
	return New(nil)
}

func (e *Environment) Define(name string, value interface{}, index int, mut string) {
	if e.values[name] == nil {
		if index == -1 {
			e.values[name] = value
		} else {
			e.indexedValues[index] = value
		}
	} else {
		if mut == "final" {
			fmt.Printf("\033[31m[Error\033[31m]\033[97m: Final %v is already defined\n", name)
			os.Exit(0)
			return
		}
		if index == -1 {
			e.values[name] = value
		} else {
			e.indexedValues[index] = value
		}
	}
}

func (e *Environment) DefineUnitialized(name string, index int) {
	if index == -1 {
		e.values[name] = needsInitialization
	} else {
		e.indexedValues[index] = needsInitialization
	}
}

func (e *Environment) Get(name token.Token, index int) (interface{}, error) {
	if index == -1 {
		v, prs := e.values[name.Lexeme]
		if prs {
			if v == needsInitialization {
				return nil, runtimeerror.Make(name, fmt.Sprintf("Uninitialized variable access: '%s'", name.Lexeme))
			}
			return v, nil
		}
		if e.enclosing != nil {
			return e.enclosing.Get(name, index)
		}
		return nil, runtimeerror.Make(name, fmt.Sprintf("Undefined variable '%v'", name.Lexeme))
	}
	return e.indexedValues[index], nil
}

func (e *Environment) GetAt(distance int, name token.Token, index int) (interface{}, error) {
	return e.Ancestor(distance).Get(name, index)
}

func (e *Environment) Ancestor(distance int) *Environment {
	env := e
	for i := 0; i < distance; i++ {
		env = env.enclosing
	}
	return env
}

func (e *Environment) Assign(name token.Token, index int, value interface{}) error {
	if index == -1 {
		if _, prs := e.values[name.Lexeme]; prs {
			e.values[name.Lexeme] = value
			return nil
		}
		if e.enclosing != nil {
			return e.enclosing.Assign(name, index, value)
		}
		return runtimeerror.Make(name, fmt.Sprintf("Undefined variable '%s'.", name.Lexeme))
	}
	e.indexedValues[index] = value
	return nil
}

func (e *Environment) AssignAt(distance int, index int, name token.Token, value interface{}) error {
	return e.Ancestor(distance).Assign(name, index, value)
}
