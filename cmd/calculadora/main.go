// Command calculadora demonstra o uso do pacote de mesmo nome.
package main

import (
	"fmt"

	"github.com/Garsaa/go-learning/projects/calculadora"
)

func main() {
	soma := calculadora.Somar(2, 3)

	fmt.Println("Hello, World!")
	fmt.Println("A soma é:", soma)
	fmt.Println("O dobro da soma é:", calculadora.Dobrar(soma))
}
