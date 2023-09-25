package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 5)
	var slice3 []int

	map1 := map[int]int{1: 1, 2: 2}
	map2 := make(map[int]int)
	var map3 map[int]int
	fmt.Println(slice1, len(slice1), cap(slice1))
	fmt.Println(slice2, len(slice3), cap(slice3))
	fmt.Println(slice3, len(slice3), cap(slice3))
	fmt.Println(map1, len(map1))
	fmt.Println(map2, len(map2))
	fmt.Println(map3, len(map3))
}

// Есть по три способа задать переменные типа slice и map:
// 1) С помощью литерала :=
// 2) С помощью функции make. Создавая slice нужно так же обязательно указать длину и если хотите размерность.
// Если не указать размерность, то она будет равна длине. Для map создастся просто пустая мапа.
// 3) С помощью ключевого слова var создается пустой slice и map. После этого можно инициализировать их значениями
