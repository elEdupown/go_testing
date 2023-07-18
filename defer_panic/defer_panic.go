package defer_panic

import (
	"fmt"
)

func ShowDefer() {
	fmt.Println("First message")
	defer fmt.Println("Last message")

	fmt.Println("Third message")
}

func ExamplePanic() {
	a := 1
	if a == 1 {
		panic("Se encontro el valor de 1")
	}
}
