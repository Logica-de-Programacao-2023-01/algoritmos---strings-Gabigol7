package main

import (
	"fmt"
	"strings"
)

func countWords(str string) int {
	// Remove os espaços em branco extras no início e no fim da string
	str = strings.TrimSpace(str)

	// Se a string for vazia, retorna 0
	if str == "" {
		return 0
	}

	// Divide a string em palavras usando espaços em branco como separador
	words := strings.Split(str, " ")

	// Retorna o número de palavras encontradas
	return len(words)
}

func main() {
	fmt.Println("Digite uma frase:")
	var input string
	fmt.Scanln(&input)

	numPalavras := countWords(input)
	fmt.Printf("A frase digitada contém %d palavras.\n", numPalavras)
}
