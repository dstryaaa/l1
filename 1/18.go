package main

import (
	"fmt"
	"sync"
)

// Count - структура, содержащая счетчик и мьютекс для безопасного доступа к счетчику.
type Count struct {
	val   int
	mutex sync.Mutex
}

// Increment - Метод увеличивает счетчик на единицу, безопасно обрабатывая конкурентный доступ.
func (c *Count) Increment() {
	c.mutex.Lock()         // Захватываем мьютекс для обеспечения безопасного доступа к счетчику.
	defer c.mutex.Unlock() // Отпускаем мьютекс после завершения данной функции.
	c.val++                // Увеличиваем счетчик на 1.
}

func main() {
	c := &Count{}
	wg := &sync.WaitGroup{} // Создаем экземпляр WaitGroup для ожидания всех горутин.
	for i := 0; i < 1000; i++ {
		wg.Add(1) // Добавляем одну горутину в WaitGroup перед запуском.
		go func() {
			defer wg.Done() // Указываем, что горутина завершена после выполнения функции Increment.
			c.Increment()   // Вызываем метод Increment для увеличения счетчика.
		}()
	}
	// Ожидаем завершения всех горутин.
	wg.Wait()
	fmt.Println(c.val)
}
