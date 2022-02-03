package main

func main() {
	miFuncion()
	miFuncion()

	saludar("Emma")

	println(suma(100, 200))
}

func miFuncion() {
	println("Esto es una lina de codigo")
	println("Esto es una nueva linea de codigo")
}

func saludar(nombre string) {
	println("Hola " + nombre + " como estas?")
}

func suma(num1 int, num2 int) int {
	suma := num1 + num2
	return suma
}
