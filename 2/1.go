package main

// Какой самый эффективный способ конкатенации строк?

// С помощью strings.Builder, в данном случае не происходит создания новой строки как например в
// конкатенации строк с помощью оператора "+". strings.Builder использует буфер для хранения
// строки и добавляет новые значения в этот буфер.

import (
	"fmt"
	"strings"
)

func main() {
	var builder strings.Builder
	builder.WriteString("Привет, ")
	builder.WriteString("Мир!")
	result := builder.String()
	fmt.Println(result)
}
