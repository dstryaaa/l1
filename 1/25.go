package main

import (
	"fmt"
	"time"
)

func sleep(seconds int) {
	// получение текущего времени
	currentTime := time.Now()
	for time.Since(currentTime).Seconds() < float64(seconds) {
		// в цикле for проверяем прошедшее время с момента вызова
		// функции и ждем пока не прошло заданное количество секунд
	}
}

func main() {
	fmt.Println("lets")
	sleep(3)
	fmt.Println("gooo")
}
