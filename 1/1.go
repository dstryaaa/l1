// Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в
// структуре Action от родительской структуры Human (аналог наследования).

package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h *Human) SayHi() {
	fmt.Printf("Hello, my name is %s and I am %d years old\n", h.name, h.age)
}

type Action struct {
	// встраивание структуры Human
	// поля и методы Human будут доступны для использования в структуре Action
	*Human
	action string
}

func (a *Action) DoAction() {
	fmt.Printf("%s did an action: %s\n", a.name, a.action)
}

func main() {
	human := &Human{name: "Kolya", age: 26}
	action := &Action{Human: human, action: "jumped"}

	human.SayHi()
	action.SayHi()

	action.DoAction()
}
