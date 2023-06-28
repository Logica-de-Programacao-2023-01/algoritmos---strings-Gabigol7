package main

import (
	"fmt"
	"strings"
)

func replaceVowels(str string) string {
	vowels := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}
	replacer := strings.NewReplacer(vowels...)

	return replacer.Replace(str)
}

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	result := replaceVowels(input)
	fmt.Println("String com vogais substituídas:", result)
}
