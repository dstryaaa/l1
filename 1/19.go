package main

import "fmt"

// reverseWord принимает строку и возвращает ее обратную версию,
// формируя новую строку и добавляя символы по одному в обратном порядке.
func reverseWord(word string) string {
	reversed := ""
	for i := len(word) - 1; i >= 0; i-- {
		reversed += string(word[i])
	}
	return reversed
}

// reverseWordRunes принимает строку и возвращает ее обратную версию,
// преобразуя строку в слайс рун и меняя местами пары рун.
func reverseWordRunes(word string) string {
	runes := []rune(word)
	length := len(runes)

	for i := 0; i < length/2; i++ {
		runes[i], runes[length-i-1] = runes[length-i-1], runes[i]
	}

	return string(runes)
}

func main() {
	word := "killimangaro"
	reversed1 := reverseWord(word)
	reversed2 := reverseWordRunes(word)
	fmt.Println("Оригинальное слово:", word)
	fmt.Println("Переворот слова с помощью создания новой строки:", reversed1)
	fmt.Println("Переворот слова с помощью массива рун:", reversed2)
}
