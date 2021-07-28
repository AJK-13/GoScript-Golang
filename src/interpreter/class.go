package interpreter

import (
	"GoScript/src/runtimeerror"
	"GoScript/src/token"
	"fmt"
)

type MetaClass struct {
	Methods map[string]*UserFunction
}

type PropertyAccessor interface {
	Get(name token.Token) (interface{}, error)
	Set(name token.Token, value interface{}) (interface{}, error)
}

type Class struct {
	Callable
	PropertyAccessor
	MetaClass  *MetaClass
	SuperClass *Class
	Name       string
	Methods    map[string]*UserFunction
	Fields     map[string]interface{}
}

func (c *Class) FindMethod(name token.Token) (*UserFunction, error) {
	if m, prs := c.Methods[name.Lexeme]; prs {
		return m, nil
	} else if c.SuperClass != nil {
		return c.SuperClass.FindMethod(name)
	}
	return nil, runtimeerror.Make(name, fmt.Sprintf("Undefined property '%s'", name.Lexeme))
}

func (c *Class) String() string {
	return fmt.Sprintf("<class %s>", c.Name)
}

func (c *Class) Get(name token.Token) (interface{}, error) {
	if v, prs := c.Fields[name.Lexeme]; prs {
		return v, nil
	}
	if m, prs := c.MetaClass.Methods[name.Lexeme]; prs {
		return m, nil
	}
	return nil, runtimeerror.Make(name, fmt.Sprintf("Undefined property '%s'", name.Lexeme))
}

func (c *Class) Set(name token.Token, value interface{}) (interface{}, error) {
	c.Fields[name.Lexeme] = value
	return nil, nil
}

func (c *Class) Call(arguments []interface{}) (interface{}, error) {
	instance := &ClassInstance{Class: c, fields: make(map[string]interface{})}
	if initializer, prs := c.Methods["init"]; prs {
		_, err := initializer.Bind(instance).Call(arguments)
		if err != nil {
			return nil, err
		}
	}

	return instance, nil
}

func (c *Class) Arity() int {
	if initializer, prs := c.Methods["init"]; prs {
		return initializer.Arity()
	}
	return 0
}

type ClassInstance struct {
	PropertyAccessor
	Class  *Class
	fields map[string]interface{}
}

func (c *ClassInstance) String() string {
	return fmt.Sprintf("<class-instance %s>", c.Class.Name)
}

func (c *ClassInstance) Get(name token.Token) (interface{}, error) {
	if v, prs := c.fields[name.Lexeme]; prs {
		return v, nil
	}

	m, err := c.Class.FindMethod(name)
	if err != nil {
		return nil, err
	}

	newMethod := m.Bind(c)
	if newMethod.Definition.IsProperty() {
		return newMethod.Call(nil)
	}
	return m.Bind(c), nil
}

func (c *ClassInstance) Set(name token.Token, value interface{}) (interface{}, error) {
	c.fields[name.Lexeme] = value
	return nil, nil
}
