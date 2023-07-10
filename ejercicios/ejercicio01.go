package ejercicios

import (
	"strconv"
)

func Ejercicio(parametro string) (int, string) {
	numero, _ := strconv.Atoi(parametro)

	if numero > 100 {
		return numero, "El número es mayor a 100"
	} else {
		return numero, "El número es menor a 100"
	}
}