package main

func main() {

	edad := 18

	if edad >= 18 {
		println("Eres major de edad")
	} else if edad < 18 && edad >= 0 {
		println("Eres menor de edad")
	} else {
		println("La edad no es valida")
	}
}
