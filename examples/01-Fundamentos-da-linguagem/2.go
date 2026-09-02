// escopo de pacotes, funções e blocos e shadowing de variáveis
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

	if fruta := "Melancia"; fruta == "Melancia" {
		// Apesar de ter 1 variável em nível de pacote com esse mesmo identificador
		// pode-se declarar com := uma outra de mesmo nome, quando se faz shadowing
		// a variável do menor nível verá a declaração de seu nível e assim sussetivamente
		// sempre se olha apra o nível atual e vai indo para cima!
		versao := "332"
		if fruta == "Melancia" {
			versao := "444"
			fmt.Println(versao)
		}
		fmt.Println(versao)
	}

	// versao eh uma variável no escopo do package
	fmt.Println(versao)
}
