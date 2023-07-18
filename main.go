package main

import (
	"github.com/elEdupown/go_testing/defer_panic"
)

func main() {
	/*estado, texto := variables.ConvertToText(1023)
	fmt.Println(estado, texto)

	if os:= runtime.GOOS; os == "linux" || os == "OS X."{
		fmt.Println("El sistema operativo es Linux")
	} else {
		fmt.Println("El sistema operativo es", os)
	}

	switch os:= runtime.GOOS; os {
	case "Linux":
		fmt.Println("El sistema operativo es Linux")
	case "darwin":
		fmt.Println("El sistema operativo es darwin")
	default:
		fmt.Printf("%s \n", os)
	}

	numero, texto := ejercicios.Ejercicio("100")
	fmt.Println(numero, texto)

	teclado.IngresoNumeros()

	iteraciones.Iterar()

	ejercicios.Tabla()

	arreglos_slices.Capacidad()
	fmt.Println(ejercicios.Tabla())
	files.GrabarTabla()
	files.SumaTabla()
	fmt.Println(files.LeerArchivo())

	mapas.ShowMap()

	users.MiUsuario()

	Eduardo := new(modelos.Hombre)
	ejer_interfaces.HumanosRespirando(Eduardo)

	Maria := new(modelos.Mujer)
	ejer_interfaces.HumanosRespirando(Maria)*/

	defer_panic.ExamplePanic()
}
