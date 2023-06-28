package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	if containsOnlyNumbers(input) {
		fmt.Println("A string contém apenas números.")
	} else {
		fmt.Println("A string não contém apenas números.")
	}
}

func containsOnlyNumbers(input string) bool {
	_, err := strconv.Atoi(input)
	return err == nil
}
