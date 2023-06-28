package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	result := strings.ToUpper(input)

	fmt.Println("Resultado:", result)
}
