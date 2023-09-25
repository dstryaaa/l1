package main

import (
	"context"
	"fmt"
	"time"
)

var flag bool

// doWork использует канал для отправки сигнала о завершении работы.
func doWork(done chan bool) {
	i := 0
	for {
		fmt.Println(i)
		if i == 9 {
			done <- true
		}
		i++
		time.Sleep(500 * time.Millisecond)
	}

	// отправка значения на канал, чтобы сигнализировать о завершении работы

}

// doWork2 использует глобальный флаг для остановки выполнения цикла
func doWork2() {
	i := 0
	for {
		fmt.Println(i)
		i++
		if flag {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
}

// doWork3 использует контекст для получения сигнала об отмене выполнения цикла.
func doWork3(ctx context.Context) {
	i := 0
	for {
		fmt.Println(i)
		i++
		time.Sleep(500 * time.Millisecond)

		// Проверяем, был ли контекст отменен. Если да, прерываем выполнение.
		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}

func main() {
	// // Первый пример с использованием канала
	done := make(chan bool)
	go doWork(done)
	// ожидание получения значения на канале
	<-done

	// // Второй пример с использованием флага
	// go doWork2()
	// time.Sleep(5 * time.Second)
	// flag = true

	// Третий пример с использованием контекста
	// ctx, cancel := context.WithCancel(context.Background())
	// go doWork3(ctx)
	// time.Sleep(5 * time.Second)
	// cancel()
}
