package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	// INTEIROS
	// signed integers...
	var numero1 int8 = 100
	fmt.Println(numero1)

	var numero2 int16 = 9999
	fmt.Println(numero2)

	var numero3 int32 = 999999999
	fmt.Println(numero3)

	var numero4 int64 = 999999999999999999
	fmt.Println(numero4)

	// Cria um int baseado na arquitetura do processador
	var numero5 int = 999999999
	fmt.Println(numero5)

	// Inferencia de tipo
	numero6 := 125444
	fmt.Println("A variavel 6 é do tipo: ", reflect.TypeOf(numero6))

	// byte é equivalente à uint8
	var numero7 byte = 123
	fmt.Println(numero7)
	fmt.Println("A variavel 7 é do tipo: ", reflect.TypeOf(numero7))

	// rune é equivalente a int32
	var numero8 rune = 1000000000
	fmt.Println(numero8)
	fmt.Println("A variavel 8 é do tipo: ", reflect.TypeOf(numero8))

	// unsigned integers...
	var numero10 uint = 10
	fmt.Println(numero10)
	fmt.Println("O numero 10 é do tipo:", reflect.TypeOf(numero10))

	var numero11 uint8 = 11
	fmt.Println(numero11)
	fmt.Println("O numero 11 é do tipo:", reflect.TypeOf(numero11))

	var numero12 uint16 = 12
	fmt.Println(numero12)
	fmt.Println("O numero 12 é do tipo:", reflect.TypeOf(numero12))

	var numero13 uint32 = 13
	fmt.Println(numero13)
	fmt.Println("O numero 13 é do tipo:", reflect.TypeOf(numero13))

	var numero14 uint64 = 143
	fmt.Println(numero14)
	fmt.Println("O numero 14 é do tipo:", reflect.TypeOf(numero14))
	// FIM INTEIROS

	// BOOLEANS
	var ok bool = false
	fmt.Println(ok)
	fmt.Println("Ok é do tipo:", reflect.TypeOf(ok))

	var nok bool = false
	fmt.Println(nok)
	fmt.Println("NOK é do tipo:", reflect.TypeOf(nok))
	// FIM BOOLEAN

	// ERROR
	var erro error = errors.New("Algo errado não está certo")
	fmt.Println(erro)
	fmt.Println("Erro é do tipo:", reflect.TypeOf(erro))
	// FIM ERROR
}
