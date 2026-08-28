package main

import (
	"fmt"
	calculadora "go-learning/calculadora"

	uuid "github.com/google/uuid"
)

func main() {
	var a = calculadora.Somar(2, 3)
	uuid.DisableRandPool()
	fmt.Println("Hello, World!")
	fmt.Println("The sum is:", a)
}
