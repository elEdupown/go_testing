package ejer_interfaces

import (
	"fmt"

	"github.com/elEdupown/go_testing/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un/a %s, y ya estoy respirando \n", hu.Sexo())
}
