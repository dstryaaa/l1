package main

import "fmt"

func main() {
	nums := []int{2, 4, 8, 10}
	ch1 := make(chan int)
	ch2 := make(chan int)

	// запускаем первую анонимную горутину для записи чисел в первый канал
	go func() {
		defer close(ch1)
		for _, num := range nums {
			ch1 <- num // отправляем значения из списка nums в канал ch1
		}
	}()

	// запускаем вторую анонимную горутину для получения данных из первого канала,
	// умножения их на 2 и запись этого во второй канал
	go func() {
		defer close(ch2)
		for num := range ch1 {
			ch2 <- num * 2 // умножаем значения полученые из ch1 на 2, и отправляем их в ch2
		}
	}()

	// цикл для чтения значений из канала ch2
	for result := range ch2 {
		fmt.Println(result)
	}
}
