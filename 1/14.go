package main

import (
	"fmt"
	"reflect"
)

// checkType принимает reflect.Type и возвращает название базового типа данных.
func checkType(t reflect.Type) string {
	typename := ""
	switch t {
	case reflect.TypeOf(int(0)):
		typename = "int"
	case reflect.TypeOf(string("")):
		typename = "string"
	case reflect.TypeOf(bool(false)):
		typename = "bool"
	}
	return typename
}

// getType определяет тип значения и его направление, если это канал,
// и выводит информацию о типе в стандартный вывод.
func getType(v interface{}) {
	switch val := reflect.TypeOf(v); val.Kind() {
	case reflect.Int:
		fmt.Printf("v is of type %T\n", v)
	case reflect.String:
		fmt.Printf("v is of type %T\n", v)
	case reflect.Bool:
		fmt.Printf("v is of type %T\n", v)
	// Обрабатываем каналы
	case reflect.Chan:
		// Получаем тип элемента канала и его направление
		elemType := checkType(val.Elem())
		dir := ""
		if val.ChanDir() == reflect.BothDir {
			dir = "bidirectional"
		} else if val.ChanDir() == reflect.RecvDir {
			dir = "receive-only"
		} else if val.ChanDir() == reflect.SendDir {
			dir = "send-only"
		}
		if elemType != "" && dir != "" {
			fmt.Printf("v is of type %s %s channel of %s values\n", dir, val.Kind(), elemType)
		} else {
			fmt.Println("v is of unknown type")
		}
	default:
		fmt.Println("v is of unknown type")
	}
}

func main() {
	var v1 interface{} = make(chan bool)
	var v2 interface{} = make(chan int, 1)
	var v3 interface{} = make(chan<- string)
	v4 := 10
	v5 := "asd"
	v6 := true

	getType(v1) // v is of type bidirectional chan of bool values
	getType(v2) // v is of type bidirectional chan of int values
	getType(v3) // v is of type send-only chan of string values
	getType(v4) // v is of type int
	getType(v5) // v is of type string
	getType(v6) // v is of type bool
}
