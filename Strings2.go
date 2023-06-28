package main

import (
	"fmt"
	"strings"
)

func removeEspacos(str string) string {
	return strings.ReplaceAll(str, " ", "")
}

func main() {
	fmt.Println("Digite uma frase:")
	var frase string
	fmt.Scanln(&frase)

	fraseSemEspacos := removeEspacos(frase)

	fmt.Println("Frase sem espaços:", fraseSemEspacos)
}
