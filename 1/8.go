package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	// Получение аргументов командной строки
	args := os.Args
	// Проверка на правильное количество аргументов
	if len(args) != 4 {
		fmt.Println("Usage: go run main.go <значение 1> <значение 2> <значение 3>") // число бит значение
		return
	}

	// Конвертирование первого аргумента в целое число (int64)
	tempnum, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Use the integer as the first argument")
		return
	}
	num := int64(tempnum)

	// Конвертирование второго аргумента в целое число (bit)
	bit, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Первый аргумент должен быть целочисленным")
		return
	}

	// Конвертирование третьего аргумента в целое число (val)
	val, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Println("Второй аргумент должен быть целочисленным")
		return
	}

	// Проверка, что бит относится к числу
	if bit <= int(math.Log2(float64(num))) {
		// Установка бита в 1
		if val == 1 {
			fmt.Printf("num в десятичной системе счисления: %d\n", num)
			num |= (1 << bit)
			fmt.Printf("num после установки %d-го бита в 1: %d\n", bit, num)
			// Установка бита в 0
		} else if val == 0 {
			fmt.Printf("num в десятичной системе счисления: %d\n", num)
			num &= ^(1 << bit)
			fmt.Printf("num после установки %d-го бита в 0: %d\n", bit, num)
		} else {
			// Если третий аргумент не равен 1 или 0
			fmt.Println("Третий аргумент должен быть 1 или 0")
		}
	} else {
		// Недопустимый бит
		fmt.Println("У числа нет такого бита")
	}

}
