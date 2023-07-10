package variables

import (
	"fmt"
	"time"
	"strconv"
)

var Nombre string // valor mínimo de string es ""
var Estado bool // valor mínimo de bool es false
var Sueldo float32 // valor mínimo de float es 0.0
var Fecha time.Time // valor mínimo de time.Time es 0001-01-01 00:00:00 +0000 UTC

func RestoVariables() {
	Nombre = "Eduardo"
	Estado = true
	Sueldo = 15000.00
	Fecha = time.Now()

	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}

func ConvertToText(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto
}