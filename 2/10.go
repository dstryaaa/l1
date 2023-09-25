package main

import "fmt"

func update(p *int) {
	b := 2
	p = &b
}

func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p)
	update(p)
	fmt.Println(*p)
}

// вывод будет 1 1, так как в функции update() изменение указателя p локальное. Потому что мы передаем в нее
// указатель по значению, а не по ссылке
