package funciones

import "fmt"

func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}

func CallClosure() {
	dos := 2
	LaTabla := tabla(dos)
	for i := 0; i < 10; i++ {
		fmt.Println(LaTabla())
	}
}
