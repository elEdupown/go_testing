package mapas

import "fmt"

func ShowMap() {
	/*paises := make(map[string]string)
	fmt.Println(paises)

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises)
	fmt.Println(paises["Mexico"])*/

	lista_de_compras := map[string]int{
		"Leche":    2,
		"Pan":      1,
		"Ajo":      3,
		"Queso":    4,
		"Arroz":    5,
		"Cilantro": 6,
	}
	fmt.Println(lista_de_compras)

	/*for producto, cantidad := range lista_de_compras {
		fmt.Printf("Producto: %s, Cantidad: %d\n", producto, cantidad)
	}*/

	delete(lista_de_compras, "Ajo")
	fmt.Println(lista_de_compras)

	cantidad, existe := lista_de_compras["Queque"]
	fmt.Printf("Cantidad: %d, Existe: %t\n", cantidad, existe)
}
