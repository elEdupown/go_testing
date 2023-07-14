package arreglos_slices

import "fmt"

var tabla [10]int = [10]int{5, 6, 7, 13, 14}
var matriz [10][20]int

func ShowArray() {
	tabla[0] = 1
	tabla[5] = 89

	tabla2 := [10]string{"Anita", "lava", "la", "tina"}
	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[5][13] = 45
	fmt.Println(matriz)
}
