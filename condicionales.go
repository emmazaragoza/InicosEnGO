//! CONDICIONALES
package main

func main() {
	var numero int = 45

	if numero == 40 {
		println(">> ejecuto el codigo que esta entre las llaves del IF <<")
	} else if numero == 41 {
		//* Esta condicion puede ir cuantas veces sea necesarias.
		println(">> ejecuto el codigo que esta entre las llaves del ELSE IF <<")
	} else {
		//* Esta condicion se la como clausula residual, en el caso de que no entre ni en IF o ELSE IF se ejecutara este ELSE, con la diferencia que se va una sola ves en el codigo y no lleva condicion, siempre se ejecutara.
		println(">> Ejecuto el condigo que esta entre las llaves del ELSE <<")

	}
}
