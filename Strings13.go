package main

import (
	"fmt"
	"strconv"
)

func isNumericSequence(str string) bool {
	length := len(str)
	if length < 2 {
		return false
	}

	for i := 0; i < length-1; i++ {
		current, _ := strconv.Atoi(string(str[i]))
		next, _ := strconv.Atoi(string(str[i+1]))

		if current+1 != next {
			return false
		}
	}

	return true
}

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	if isNumericSequence(input) {
		fmt.Println("A string é uma sequência numérica crescente.")
	} else {
		fmt.Println("A string não é uma sequência numérica crescente.")
	}
}
