package main

import (
	"fmt"
	"math"
)

// Point - структура содержит координаты x и y точки
type Point struct {
	x float64
	y float64
}

// PointerInterface - интерфейс, который определяет два метода - первый для получения координаты x точки,
// второй - для координаты y
type PointerInterface interface {
	getX() float64
	getY() float64
}

// Pointers - структура содержит две точки типа PointerInterface для вычисления расстояния между ними
type Pointers struct {
	a PointerInterface
	b PointerInterface
}

// NewPoint - функция создает и возвращает новый объект типа Point с заданными координатами x и y
func NewPoint(x, y float64) Point {
	a := Point{x: x, y: y}
	return a
}

// getX - метод для получения координаты x точки
func (p *Point) getX() float64 {
	return p.x
}

// getY - метод для получения координаты y точки
func (p *Point) getY() float64 {
	return p.y
}

// distance - метод для структуры Pointers, который вычисляет расстояние между двумя точками
// с использованием теоремы Пифагора
func (p *Pointers) distance() float64 {
	x := p.a.getX() - p.b.getX()
	y := p.a.getY() - p.b.getY()
	if x < 0 {
		x = x * (-1)
	}
	if y < 0 {
		y = y * (-1)
	}
	dstnc := math.Sqrt(x*x + y*y)
	return dstnc
}

func main() {
	aP := NewPoint(2, 2)
	bP := NewPoint(5, 5)
	var ds Pointers
	ds.a = &aP
	ds.b = &bP
	fmt.Printf("%.2f", ds.distance())
}
