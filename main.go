package main

import (
	"fmt"
	"github.com/elEdupown/go_testing/files"
	//"runtime"
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

	fmt.Println(ejercicios.Tabla())
	files.GrabarTabla()
	files.SumaTabla() */
	fmt.Println(files.LeerArchivo())
}