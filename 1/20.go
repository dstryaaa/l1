package main

import (
	"fmt"
	"strings"
)

// Функция reverseStringUsingPackageStrings переворачивает последовательность слов в строке,
// используя функции из пакета "strings". Получает на вход строку и возвращает модифицированную строку.
func reverseStringUsingPackageStrings(str string) string {
	words := strings.Split(str, " ") // Разделение строки на слова по пробелу. Создается массив "слов".
	reversed := ""                   // Итоговая строка с перевернутыми словами
	// Проход по массиву слов и последовательное их соединение в новой строке,
	// выполняется перестановка слов в обратном порядке.
	for _, word := range words {
		if reversed == "" {
			reversed = word
		} else {
			reversed = word + " " + reversed
		}
	}
	return reversed
}

// Функция reverseString переворачивает последовательность слов в строке,
// без использования пакета "strings". Получает на вход строку и возвращает модифицированную строку.
func reverseString(str string) string {
	wordStart := -1 // Индекс начала текущего слова.
	wordEnd := -1   // Индекс конца текущего слова.
	reversed := ""  // Итоговая строка с перевернутыми словами

	// Проход по всем символам исходной строки.
	for i := 0; i < len(str); i++ {
		// Если находим первый символ слова и начало слова еще не было определено:
		if str[i] != ' ' && wordStart == -1 {
			wordStart = i
		}

		// Если находим пробел или последний символ в строке и начало слова определено:
		if (str[i] == ' ' || i == len(str)-1) && wordStart != -1 {
			wordEnd = i // Задаем конец слова

			// Если это последний символ (и не пробел), увеличиваем wordEnd на 1
			if i == len(str)-1 && str[i] != ' ' {
				wordEnd++
			}

			// Если строка с перевернутыми словами еще пуста, копируем туда первое найденное слово
			if reversed == "" {
				reversed = str[wordStart:wordEnd]
			} else {
				// Иначе склеиваем текущее найденное слово и уже существующую строку с разделителем - пробел
				reversed = str[wordStart:wordEnd] + " " + reversed
			}

			// Сбрасываем значение начальной и конечной позиций слова для поиска следующего слова
			wordStart = -1
			wordEnd = -1
		}
	}

	return reversed // Возвращаем строку с переставленными словами
}

func main() {
	words := "dog cat monkey"
	fmt.Println("Начальная строка", words)
	fmt.Println("С использованием пакета strings", reverseStringUsingPackageStrings(words))
	fmt.Println("Без использования пакета strings", reverseString(words))
}
