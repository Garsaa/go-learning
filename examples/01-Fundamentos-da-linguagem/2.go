// escopo de pacotes, funções e blocos
package main

// packages não compartilham imports
import "fmt"

//declarações de variáveis e constantes podem ser acessadas em blocos inferiores/internos a blocos mas não exteriores

func SecondStep() {
	// Da pra declarar variável com declaração curta no if
	if idade2, idade := firstStep(), 2; idade > 11 {
		// Não tem valor truthy ou falsy em Go, tem que ser booleano MESMO pra ser usado de condition
	} else if idade2 == 2 {

	}

	// println(idade) aq nao existe
	// versao eh uma variável no escopo do package
	fmt.Println(versao)
}
