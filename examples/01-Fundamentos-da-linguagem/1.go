// Estrutura, Variáveis, constantes e zero values
package main

import (
	"fmt"
)

// em Go a declaração de uma variável eh:
// constante ou variável identificador tipo = valorDeInicialização
const versao string = "212"

// O escopo das variáveis em go é definido por bloco
// seja um pacote, uma função, um if ou enfim.

//go também pode inferir o tipo, não eh necessário explícitar]
//uma vez declarado, o tipo não muda, Go é estáticamente tipado

// zero value eh o valor q 1 variável recebe quando não é inicializada.
func firstStep() int {
	// dentro de funções da pra fazer declaração curta, essas são sempre vars
	nome, idade := "gabriel", 23
	// da pra deixar o go inferir o tipo
	const something = 2.20
	const linguagem string = "Go"
	// Da pra fazer declaração de varias variáveis em 1 linha
	// de forma posicional
	var cor, comprimento, largura = "Azul", 2, 3
	// também é possível declarar variáveis em bloco
	var (
		tipo              = "carro"
		quantidadeDeRodas = 4
	)
	// Existe esse iota ai, ele inicia 1 contagem em 0 e atribui esse valor pra iniciar as const
	// Em um bloco de declarações se nenhum valor for atribuído às consts seguintes, ele reutiliza
	// q o valor das anteriores (nesse caso, iota)
	const (
		numero0 = iota
		numero1
		numero2
		numero3
	)
	nome = "something3"
	fmt.Println(
		linguagem, versao, nome, idade,
		cor, comprimento, largura,
		tipo, quantidadeDeRodas, numero0, numero1, numero2, numero3,
	)
	return 3
}
