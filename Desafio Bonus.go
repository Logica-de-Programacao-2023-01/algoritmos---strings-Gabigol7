package main

import (
	"fmt"
	"strings"
)

func findPatternIndices(str, pattern string) []int {
	var indices []int
	start := 0

	for {
		index := strings.Index(str[start:], pattern)
		if index == -1 {
			break
		}
		indices = append(indices, start+index)
		start = start + index + len(pattern)
	}

	return indices
}

func main() {
	str := "A string de exemplo para teste de padrão"
	pattern := "exemplo"

	indices := findPatternIndices(str, pattern)
	if len(indices) == 0 {
		fmt.Println("O padrão não foi encontrado na string.")
	} else {
		fmt.Println("O padrão foi encontrado nos seguintes índices:")
		for _, index := range indices {
			fmt.Println(index)
		}
	}
}
