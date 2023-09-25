package main

import "fmt"

// binarySearch осуществляет бинарный поиск элемента x в отсортированном массиве arr.
// Он возвращает индекс элемента в массиве или -1, если элемент не найден.
func binarySearch(arr []int, x int) int {
	// Объявляем левую и правую границы поиска
	left := 0
	right := len(arr) - 1
	// Повторяем поиск, пока левая граница не станет больше правой
	for left <= right {
		mid := (left + right) / 2 // определение середины массива
		if arr[mid] == x {        // если элемент найден, возвращаем индекс
			return mid
		} else if arr[mid] < x { // если элемент меньше искомого, смещаем левую границу вправо
			left++
		} else { // если элемент больше искомого, смещаем правую границу влево
			right--
		}
	}
	return -1
}

func main() {
	arr := []int{2, 4, 7, 10, 13}
	x := 10
	index := binarySearch(arr, x)
	if index == -1 {
		fmt.Printf("%d not found in array\n", x)
	} else {
		fmt.Printf("%d found at index %d\n", x, index)
	}
}
