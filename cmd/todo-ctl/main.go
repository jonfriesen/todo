package main

import (
	"fmt"

	"github.com/jonfriesen/todo"
)

func main() {
	a := todo.List(true)
	fmt.Printf("%+v\n", a)
	a, err := todo.Set("rawr")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", a)
	a, err = todo.Set("asd")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", a)
	a, err = todo.Set("jgh")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", a)
	a = todo.List(true)
	fmt.Printf("%+v\n", a)

	a, err = todo.Complete("11")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", a)
}
