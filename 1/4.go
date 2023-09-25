package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var numWorkers int // количество воркеров
	fmt.Scan(&numWorkers)

	c := make(chan int) // создание канала

	// Запуск задачи, записывающей произвольные числа в канал
	go func() {
		for {
			c <- rand.Intn(100)     // генерация произвольного числа и запись в канал
			time.Sleep(time.Second) // задержка 1 секунду
		}
	}()

	// Запуск воркеров
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				data := <-c                             // чтение данных из канала
				fmt.Printf("Worker %d: %d\n", id, data) // вывод данных в стандартный вывод
			}
		}(i)
	}

	wg.Wait() // ожидание завершения всех воркеров
}
