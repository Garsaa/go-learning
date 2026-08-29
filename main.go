package main

import (
	"fmt"
	"go-learning/calculadora"
	// "github.com/google/uuid"
)

func main() {
	var a = calculadora.Somar(2, 3)
	// var b = uuid.NewString()
	fmt.Println("Hello, World!")
	fmt.Println("The sum is:", a)
	fmt.Println("2x da soma", calculadora.Faz2x(a))
}
