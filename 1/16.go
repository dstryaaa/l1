package main

import "fmt"

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	// Определение опорного элемента в массиве
	pivot := arr[(left+right)/2]

	i, j := left, right

	// Сортировка элементов в массиве по отношению к опорному элементу
	for i <= j {
		// Ищем элемент слева, который больше или равен опорному
		for arr[i] < pivot {
			i++
		}
		// Ищем элемент справа, который меньше или равен опорному
		for arr[j] > pivot {
			j--
		}
		// Если индексы не "пересеклись", меняем местами элементы и продолжаем искат
		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}

	// Рекурсивный вызов функции сортировки для левой части массива
	if left < j {
		quickSort(arr, left, j)
	}
	// Рекурсивный вызов функции сортировки для правой части массива
	if right > i {
		quickSort(arr, i, right)
	}
}

func main() {
	// Входной массив
	arr := []int{6, 4, 9, 3, 1, 7, 2, 8, 5}

	// Вывод неотсортированного массива
	fmt.Println("Unsorted array:", arr)

	// Сортировка и вывод отсортированного массива
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("Sorted array:  ", arr)
}
