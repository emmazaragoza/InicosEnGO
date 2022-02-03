package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
	edad     int
}

//* Metodo
func (p persona) saludar(saludo string) {
	fmt.Println(saludo + " " + p.nombre + " " + p.apellido)
}

func main() {
	persona1 := persona{"Emmanuel", "Zaragoza", 25}

	persona2 := persona{"Bill", "Gates", 65}

	fmt.Println("Persona 1: ", persona1.nombre, persona1.apellido, persona1.edad)
	fmt.Println("Persona 2: ", persona2)

	//! cambiar los valores dentro el obj.
	persona1.edad = 20
	fmt.Println("Persona 1: ", persona1.nombre, persona1.apellido, persona1.edad)

	//! Metodos en GO.
	persona1.saludar("Hola")
	persona2.saludar("Hello")
}
