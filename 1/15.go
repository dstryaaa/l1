package main

import "strings"

// // Этот код может привести к утечке памяти
// var justString string
// // В этой функции создается большая строка и из нее берется только 100 символов.
// // А оставшийся кусок строки не будет удален из памяти, что может привести к ее нехватке.
// func someFunc() {
// 	v := createHugeString(1 << 10)
// 	justString = v[:100]
// }

// func main() {
// 	someFunc()
// }

var justString string

func someFunc() {
	var builder strings.Builder
	builder.Grow(1 << 10)
	builder.WriteString(createHugeString())

	justString = builder.String()[:100]
}

func main() {
	someFunc()
}
