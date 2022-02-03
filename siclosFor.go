package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	//* Recorer una cadena de texto. str

	nombre := "Emmanuel"

	for i := 0; i < len(nombre); i++ {
		fmt.Println(i, string(nombre[i]))
	}

	//! IMPORTANTE: En Golang NO EXITE el while, se usa el siclo for.

	numero := 0

	for numero <= 10 {
		fmt.Println(numero, ": Ejecuto el codigo del siclo for reemplazando al while")
		numero++
	}
}
