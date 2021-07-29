package object

import (
	_ "fmt"
)

func NewEnv() *Env {
	s := make(map[string]Object)
	return &Env{
		Store: s,
	}
}

type Env struct {
	Store map[string]Object
}

func (e *Env) SetEnv(name string, obj Object) {

	e.Store[name] = obj
}

func (e *Env) GetEnv(name string) (Object, bool) {

	val, ok := e.Store[name]
	// fmt.Println("sampling +> ", val)
	return val, ok
}

func (e *Env) ChangeVal(name string, elm Object) bool {

	_, ok := e.Store[name]

	if ok {
		e.Store[name] = elm

		return ok
	}

	return ok
}

func (e *Env) IndexValChange(name string, index int, value Object) bool {

	val, ok := e.Store[name]

	if ok {
		val.(*ArrayContents).Body[index] = value
		e.ChangeVal(name, val)

		return ok
	}

	return ok
}
