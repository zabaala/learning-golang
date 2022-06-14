package main

import "fmt"

func main() {
	// variavel tipada
	var variavel1 string = "conteúdo da PRIMEIRA variável"
	fmt.Println(variavel1)

	// tipagem implicita
	variavel2 := "conteúdo da SEGUNDA variavel"
	fmt.Println(variavel2)

	// Declaração conjunta de variáveis
	var (
		variavel3 string = "Conteúdo da variavel 3"
		variavel4 string = "Conteúdo da variavel 4"
	)
	fmt.Println(variavel3, "\n", variavel4)

	variavel5, variavel6 := "Conteúdo da QUINTA variavel", "Conteúdo da SEXTA variavel"

	fmt.Println(variavel5)
	fmt.Println(variavel6)

	// trocando valores das variaveis
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5)
	fmt.Println(variavel6)

	// Constante
	const constante1 string = "Conteúdo da constante"
	fmt.Println(constante1)
}
