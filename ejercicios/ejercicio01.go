package ejercicios

import (
	"strconv"
)

func Ejercicio(parametro string) (int, string) {
	numero, err := strconv.Atoi(parametro)

	if err != nil {
		return 0, "Hubo un error" + err.Error()
	}

	if numero > 100 {
		return numero, "\n El número es mayor a 100"
	} else {
		return numero, "\n El número es menor a 100"
	}
}