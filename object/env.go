package object

import (
	_ "fmt"
)

func NewEnv() *Env {
	s := make(map[string]Object)
	return &Env{
		store: s,
	}
}

type Env struct {
	store map[string]Object
}

func (e *Env) SetEnv(name string, obj Object) {
	e.store[name] = obj
}

func (e *Env) GetEnv(name string) (Object, bool) {

	val, ok := e.store[name]

	return val, ok
}
