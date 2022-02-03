package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
	edad     int
}

func main() {
	persona1 := persona{"Emmanuel", "Zaragoza", 25}

	persona2 := persona{"Bill", "Gates", 65}

	fmt.Println("Persona 1: ", persona1.nombre, persona1.apellido, persona1.edad)
	fmt.Println("Persona 2: ", persona2)

	//! cambiar los valores dentro el obj.
	persona1.edad = 20
	fmt.Println("Persona 1: ", persona1.nombre, persona1.apellido, persona1.edad)
}
