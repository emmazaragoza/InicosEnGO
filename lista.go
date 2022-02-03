package main

func main() {
	frutas := []string{"mazana", "pera", "naranja"}
	println(frutas[1])

	//* añadir mas elementos a la lista.

	frutas = append(frutas, "melocoton", "melon")

	for i := 0; i < len(frutas); i++ {
		println(frutas[i])
	}

	//* Para canbien algun valor dentro de la lista.

	frutas[0] = "sandia"
	for i := 0; i < len(frutas); i++ {
		println(frutas[i], "DE NUEVO")
	}

}
