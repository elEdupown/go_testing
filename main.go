package main

import (
<<<<<<< HEAD
	//"fmt"
	"github.com/elEdupown/go_testing/arreglos_slices"
=======
	"fmt"
	"github.com/elEdupown/go_testing/files"
>>>>>>> ac3b6d217ea8f9952864c245934a6e3160ea49a8
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

<<<<<<< HEAD
	ejercicios.Tabla()*/

	arreglos_slices.Capacidad()
}
=======
	fmt.Println(ejercicios.Tabla())
	files.GrabarTabla()
	files.SumaTabla() */
	fmt.Println(files.LeerArchivo())
}
>>>>>>> ac3b6d217ea8f9952864c245934a6e3160ea49a8
