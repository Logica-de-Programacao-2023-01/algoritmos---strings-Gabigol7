package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isCamelCase(str string) bool {
	if str == "" {
		return false
	}

	firstChar := rune(str[0])
	if !unicode.IsUpper(firstChar) {
		return false
	}

	hasLowerCase := false
	hasUpperCase := false

	for _, char := range str {
		if unicode.IsSpace(char) || char == '_' {
			return false
		}

		if unicode.IsUpper(char) {
			if hasLowerCase {
				return false
			}
			hasUpperCase = true
		}

		if unicode.IsLower(char) {
			hasLowerCase = true
		}
	}

	return hasUpperCase && hasLowerCase
}

func countWords(str string) int {
	if str == "" {
		return 0
	}

	words := strings.FieldsFunc(str, func(r rune) bool {
		return unicode.IsUpper(r) || unicode.IsSpace(r) || r == '_'
	})

	return len(words)
}

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	if isCamelCase(input) {
		wordCount := countWords(input)
		fmt.Printf("A string está em camelCase e possui %d palavra(s).\n", wordCount)
	} else {
		fmt.Println("A string não está em camelCase.")
	}
}
