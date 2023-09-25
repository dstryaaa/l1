package main

import "fmt"

func main() {
	n := 0
	if true {
		n := 1
		n++
	}
	fmt.Println(n)
}

// Мы создаем переменную n снаружи блока if, и потом заново создаем ее в блоке локально. В результате чего
// создавая переменную локально в блоке if мы затенняем переменную созданную в мейне и выходя из блока if она
// выходит из области видимости и уничтожается, а старая переменная n снова станет видимой.
// Соответственно выведется значение 0
