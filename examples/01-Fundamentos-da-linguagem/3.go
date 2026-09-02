package main

import "fmt"

// thirdStep demonstra os zero values das principais categorias de tipos em Go.
//
// Zero value é o valor que uma VARIÁVEL recebe automaticamente quando é
// declarada sem uma inicialização explícita. Constantes sempre precisam ter
// um valor definido.
func thirdStep() {
	// Booleanos recebem false.
	var booleano bool

	// Todos os tipos numéricos recebem zero.
	var inteiro int
	var inteiroComSinal int64
	var inteiroSemSinal uint64
	var byteZero byte // byte é um alias de uint8
	var runeZero rune // rune é um alias de int32
	var endereco uintptr
	var decimal float64
	var complexo complex128

	// Strings recebem a string vazia: "".
	var texto string

	fmt.Println("\n--- Zero values dos tipos básicos ---")
	fmt.Printf("bool:       %v\n", booleano)
	fmt.Printf("int:        %d\n", inteiro)
	fmt.Printf("int64:      %d\n", inteiroComSinal)
	fmt.Printf("uint64:     %d\n", inteiroSemSinal)
	fmt.Printf("byte:       %d\n", byteZero)
	fmt.Printf("rune:       %d\n", runeZero)
	fmt.Printf("uintptr:    %d\n", endereco)
	fmt.Printf("float64:    %g\n", decimal)
	fmt.Printf("complex128: %v\n", complexo)
	fmt.Printf("string:     %q (tamanho %d)\n", texto, len(texto))

	// O zero value de ponteiros, slices, maps, channels, funções e interfaces é
	// nil. nil significa que a variável ainda não referencia um valor concreto.
	var ponteiro *int
	var numeros []int
	var notas map[string]float64
	var mensagens chan string
	var operacao func(int, int) int
	var qualquerCoisa any

	fmt.Println("\n--- Tipos cujo zero value é nil ---")
	fmt.Printf("ponteiro:  %v; é nil? %t\n", ponteiro, ponteiro == nil)
	fmt.Printf("slice:     %v; é nil? %t\n", numeros, numeros == nil)
	fmt.Printf("map:       %v; é nil? %t\n", notas, notas == nil)
	fmt.Printf("channel:   %v; é nil? %t\n", mensagens, mensagens == nil)
	fmt.Printf("função:    %v; é nil? %t\n", operacao, operacao == nil)
	fmt.Printf("interface: %v; é nil? %t\n", qualquerCoisa, qualquerCoisa == nil)

	// Arrays e structs não recebem nil. Cada elemento ou campo recebe o zero
	// value correspondente ao próprio tipo.
	var array [3]int
	var pessoa struct {
		nome  string
		idade int
		ativo bool
	}

	fmt.Println("\n--- Tipos compostos preenchidos recursivamente ---")
	fmt.Printf("array:  %#v\n", array)
	fmt.Printf("struct: %#v\n", pessoa)
}
