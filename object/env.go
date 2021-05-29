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
