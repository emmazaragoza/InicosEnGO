package main

// Que es el package: Es una palabra reservada de GO que define el paquete de codigo que pertenece este archivo.
// 		solo puede haber un solo paquete por carpeta y cada archivo (.go) debe declarar el mismo nombre del paquete
// 		en la parte superior. En este caso el ccodigo pertenece al paquete "main"

// Que es import: Es otra palabra resevada que le indica al compilador de GO que requiere otros componentes para
// 		usar en este archivo. EJ: import "fmt"

// Variables en Golang
// en Golang se define la variable y se le debe de poner el tipo de dato que guardara (str, int, float)
// Ej:
var numeros int = 10
var decimal float64 = 1.22

// Boolean
var bolean bool = true

// O tambien puede ser contaste. Con comillas dobles
const nombre string = "emma"

//pero tambien se puede definir como "variables cortas"
// i := int64(1)

// iniciar valores en cero.
var apellido string
var edad int
var altura float64

//! Resumen para declarar una variable >> var (nombre) (tipoDeDato) = (valor)
/*
*Datos a tener en cuenta:
--------------------------------------------------------------------------------------------------------------
             DATO                                NO SE PUEDE                 SI SE PUEDE
	>	No se permiten guines.                   user-name                    userName
	>	No puede comenzar con un numero.          4name                         name4
	>	No pueden llevar simbolos.                $name                         name
	>	No puede tenes mas de una palabra.   nombre de usuario             nombreDeUsuario

	>> Tener en cuenta que distingue entre MAYUSCULAS y minusculas
--------------------------------------------------------------------------------------------------------------
*/

func main() {
	numeros = 15 //! Las variables si se pueden pisar, pero dentro de una funcion.
	i := 10
	i = 12
	apellido = "zaragoza"
	println(numeros, decimal, nombre, "zaragoza", apellido, i)

	//! OPERADORES MATEMATICO Y DE TIPO LOGICO

	//! cambiar a main y eliminar el de arriba
	suma := 10 + 15
	resta := 10 - 15
	multiplicacion := 10 * 2
	divicion := 10 / 2
	resto := 10 % 3

	println(suma, resta, multiplicacion, divicion, resto)

	suma++
	suma--
	println(suma)

	//! operadores logicos.
	//* Mayor - Igual, Mayor o Menor, Igualdad, And , Or
	println(10 > 5)
	println(10 < 5)

}
