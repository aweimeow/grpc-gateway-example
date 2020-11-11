package main

import "fmt"

type Gender uint32
const (
	MALE Gender = iota
	FEMALE
	TRANSGENDER
	NOTDEFINED
)

type Employee struct {
	name string
	gender Gender
	age uint32
}

func (e *Employee) String() string {
	return fmt.Sprintf("Employee{name=%s, age=%d}", e.name, e.age)
}