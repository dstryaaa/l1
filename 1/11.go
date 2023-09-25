package main

import "fmt"

func main() {
	// Создание двух множеств
	set1 := map[int]bool{1: true, 2: true, 3: true}
	set2 := map[int]bool{2: true, 3: true, 4: true}

	// Создание множества для хранения пересечения
	intersection := make(map[int]bool)

	// Перебор элементов первого множества
	for k := range set1 {
		// Если элемент есть во втором множестве, добавить его в пересечение
		if set2[k] {
			intersection[k] = true
		}
	}

	fmt.Println(intersection)
}
