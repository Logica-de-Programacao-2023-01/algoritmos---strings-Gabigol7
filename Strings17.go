package main

import "fmt"

func main() {
	var input string

	fmt.Println("Digite uma string:")
	fmt.Scanln(&input)

	uniqueLetters := findUniqueLetters(input)

	fmt.Println("Letras únicas na string:", uniqueLetters)
}

func findUniqueLetters(input string) string {
	letterCount := make(map[rune]int)

	// Conta a ocorrência de cada letra na string
	for _, char := range input {
		letterCount[char]++
	}

	uniqueLetters := ""

	// Verifica quais letras aparecem apenas uma vez e as adiciona à string de letras únicas
	for _, char := range input {
		if letterCount[char] == 1 {
			uniqueLetters += string(char)
		}
	}

	return uniqueLetters
}
