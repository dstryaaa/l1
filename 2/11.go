package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}

// В результате выполнения этой программы мы получим дедлок, так как передаем WaitGroup в нашу функцию
// не по указателю. Из-за этого каждая наша горутина будет иметь свою копию WaitGroup объекта
// и между ними не будет синхронизации
