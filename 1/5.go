package main

import (
	"fmt"
	"time"
)

// функция sender отправляет целочисленные значения по каналу c в бесконечном цикле 2 раза в секунду
func sender(c chan int) {
	i := 0
	for {
		c <- i
		fmt.Printf("sent %d\n", i)
		i++
		time.Sleep(500 * time.Millisecond)
	}
}

// функция receiver принимает целочисленные значения из канала c в бесконечном цикле
func receiver(c chan int) {
	for {
		msg := <-c
		fmt.Printf("received %d\n", msg)
	}
}

func main() {
	var endTime int
	fmt.Scan(&endTime)
	c := make(chan int)
	go sender(c)
	go receiver(c)
	time.Sleep(time.Duration(endTime) * time.Second)
}
