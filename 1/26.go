package main

import "fmt"

// uniqueString проверяет, состоит ли строка только из уникальных символов.
// Возвращает true, если все символы уникальны, и false в противном случае.
func uniqueString(str string) bool {
	// Создается отображение для хранения обнаруженных символов.
	chars := make(map[rune]bool)
	for _, char := range str {
		// Если символ уже был обнаружен, возвращается false.
		if chars[char] {
			return false
			// Если символ встречается впервые, добавляется в отображение.
		} else {
			chars[char] = true
		}
	}
	return true
}

func main() {
	s1 := "abcd"
	s2 := "abCdefAaf"
	s3 := "aabcd"
	fmt.Println(uniqueString(s1), uniqueString(s2), uniqueString(s3))
}
