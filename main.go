package main

import (
	"fmt"
	"github.com/elEdupown/go_testing/ejericios"
	//"runtime"
	//"github.com/elEdupown/go_testing/variables"
)

func main() {
	/*estado, texto := variables.ConvertToText(1023)
	fmt.Println(estado, texto)*/

	/*if os:= runtime.GOOS; os == "linux" || os == "OS X."{
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
	} */

	numero, texto := ejercicios.Ejercicio("100")
	fmt.Println(numero, texto)
}