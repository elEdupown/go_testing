package variables

import "fmt"

func ShowIntegers() {
	var intcommon int // int es 0 por defecto en go, el valor minimo de entero
	int32 := int32(15)
	int64 := int64(150)

	fmt.Println("intcommon = ", intcommon)
	fmt.Println("int32 = ", int32)
	fmt.Println("int64 = ", int64)
}