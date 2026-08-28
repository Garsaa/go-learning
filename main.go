package main

import (
	"fmt"
	calculadora "go-learning/calculadora"
)

func main() {
	var a = calculadora.Somar(2, 3)
	fmt.Println("Hello, World!")
	fmt.Println("The sum is:", a)
}
