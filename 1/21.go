package main

import "fmt"

// Adapter - интерфейс, содержащий метод Request для адаптированных запросов
type Adapter interface {
	Request() string
}

// Adaptee - это адаптируемый класс, который содержит специфичный метод SpecificRequest
type Adaptee struct {
	Name string
}

// SpecificRequest - метод класса Adaptee, возвращающий специфичную строку

func (a *Adaptee) SpecificRequest() string {
	return "Adaptee - " + a.Name
}

// AdapterImpl - структура, реализующая подход Адаптер, внедряя Adaptee
type AdapterImpl struct {
	*Adaptee
}

// Request - метод, который реализут интерфейс Adapter и делегирует вызов метода SpecificRequest
func (a *AdapterImpl) Request() string {
	return a.SpecificRequest()
}

func main() {
	// создаем объект класса Adaptee
	adaptee := &Adaptee{"Tobi"}

	// адаптер AdapterImpl адаптирует вызовы к методу SpecificRequest класса Adaptee
	adapter := &AdapterImpl{adaptee}

	fmt.Println(adapter.Request())
}
