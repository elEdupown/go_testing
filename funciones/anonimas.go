package funciones

import "fmt"

func Calculos() {
	var numero3, numero4 int = 45, 123

	calculo := func(numero1, numero2 int) int {
		return numero1 + numero2 + numero3 + numero4
	}

	fmt.Println(calculo(5, 6))

	calculo = func(numero1, numero2 int) int {
		return numero1 * numero2 * numero3
	}

	fmt.Println(calculo(5, 6))

}
