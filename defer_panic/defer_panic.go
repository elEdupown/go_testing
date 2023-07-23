package defer_panic

import (
	"fmt"
	"log"
)

func ShowDefer() {
	fmt.Println("First message")
	defer fmt.Println("Last message")

	fmt.Println("Third message")
}

func ExamplePanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrio un error que genero Panic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("Se encontro el valor de 1")
	}
}
