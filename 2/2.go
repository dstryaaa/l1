// Интерфейсы в Go являются абстракциями, которые определяют набор методов, которые тип должен реализовать,
// чтобы удовлетворять интерфейсу.
// Интерфейсы в Go используются для создания гибких и расширяемых систем. Они позволяют работать с различными
// типами данных, которые реализуют один и тот же интерфейс, без необходимости знать конкретный тип этого данных.
// Кроме того, интерфейсы в Go часто используются как контракты между пакетами, что упрощает разработку
// сложных приложений
package main

import (
	"fmt"
	"strconv"
)

// Formatter - это интерфейс, который требует реализации метода Format() string.
type Formatter interface {
	Format() string
}

type Person struct {
	Name string
}

// Format - метод, который реализует интерфейс Formatter для структуры Person.
func (p *Person) Format() string {
	return "Name: " + p.Name
}

type Number struct {
	Value int
}

// Format - метод, который реализует интерфейс Formatter для структуры Number.
func (n *Number) Format() string {
	return "Value: " + strconv.Itoa(n.Value)
}

// PrintFormated - функция, которая принимает форматтер и печатает отформатированную строку.
func PrintFormated(f Formatter) {
	fmt.Println(f.Format())
}

func main() {
	p := Person{Name: "John"}
	PrintFormated(&p)
	n := Number{Value: 26}
	PrintFormated(&n)
}
