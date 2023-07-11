package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var num1 int
var num2 int
var legend string
var err error

func IngresoNumeros() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese el primer número: ")
	if scanner.Scan(){
		num1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Exploté jeje: "+ err.Error())
		}
	}

	fmt.Println("Ingrese el segundo número: ")
	if scanner.Scan(){
		num2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Exploté jeje: "+ err.Error())
		}
	}

	fmt.Println("Ingrese la leyenda: ")
	if scanner.Scan(){
		legend = scanner.Text()
	}

	fmt.Println(legend, num1*num2)
}