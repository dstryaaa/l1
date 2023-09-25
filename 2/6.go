// В Go нет поддержки перегрузки методов или операторов.

// Каждый метод или функция, определенная в Go, должна иметь уникальное имя в пределах пакета.
// Это означает, что невозможно определить две функции, имеющие одинаковое имя в рамках одного пакета,
// независимо от типов параметров или их количества.

// Это связано с философией языка Go, которая нацелена на простоту и удобство чтения кода.
// Вместо перегрузки методов и операторов Go использует интерфейсы для достижения полиморфизма и расширяемости
package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func GetArea(s Shape) float64 {
	return s.Area()
}

func main() {
	rectangle := Rectangle{Width: 5, Height: 2}
	circle := Circle{Radius: 4}
	shapes := []Shape{rectangle, circle}

	for _, shape := range shapes {
		fmt.Println("площадь фигуры", GetArea(shape))
	}
}
