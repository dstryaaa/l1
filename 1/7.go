package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	var mu sync.Mutex // Создаем мьютекс
	var wg sync.WaitGroup

	m := make(map[int]string)

	// Конкурентный доступ к map
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			mu.Lock() // Блокируем доступ к map для других горутин
			m[i] = fmt.Sprint(rand.Intn(100))
			mu.Unlock() // Разблокируем доступ к map
			wg.Done()
		}(i)
	}

	// Ждем, пока все горутины не завершат свое выполнение
	wg.Wait()
	fmt.Println(m)
}
