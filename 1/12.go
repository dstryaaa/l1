package main

import "fmt"

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем пустое множество с помощью оператора map
	set := make(map[string]bool)

	// Проходим по всем элементам последовательности и добавляем их в множество
	for _, str := range sequence {
		set[str] = true
	}

	// Выводим все элементы множества на экран
	for key, _ := range set {
		fmt.Println(key)
	}

}
