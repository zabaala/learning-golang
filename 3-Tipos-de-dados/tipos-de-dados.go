package main

import (
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

	// FIM INTEIROS

}
