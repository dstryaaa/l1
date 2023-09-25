package main

import "fmt"

// первый Способ
func deleteByIndex(slice []int, index int) []int {
	var newSlice []int
	// Цикл проходится по всем элементам среза.
	for i := 0; i <= len(slice)-1; i++ {
		// Если индекс текущего элемента не равен индексу элемента для удаления,
		// добавляем его в новый срез.
		if i != index {
			newSlice = append(newSlice, slice[i])
		}
	}
	return newSlice
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	index := 4
	fmt.Println("Начальный срез ", slice)
	fmt.Printf("Срез после удаление %d-го элемента %v", index, deleteByIndex(slice, index))
}

// второй способ
func main() {
	slice := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	fmt.Println("Начальный срез ", slice)
	index := 5
	// Удаляем элемент с помощью встроенной функции append, "склеивая" два среза,
	// исключая элемент с заданным индексом.
	slice = append(slice[:index], slice[index+1:]...)
	fmt.Printf("Срез после удаление %d-го элемента %v", index, slice)
}

// Третий способ
func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Начальный срез ", slice)
	index := 4

	// Используем функцию copy для копирования элементов,
	// начиная с (index+1)-го, на место элементов, начиная с индекса.
	copy(slice[index:], slice[index+1:])
	// После копирования удаляем последний элемент из среза, так как он дублируется.
	slice = slice[:len(slice)-1]
	fmt.Printf("Срез после удаление %d-го элемента %v", index, slice)
}
