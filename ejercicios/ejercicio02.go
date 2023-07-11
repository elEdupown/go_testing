package ejercicios

import (
	"fmt"
	"strconv"
	"bufio"
	"os"
)


func Tabla() {
	scanner := bufio.NewScanner(os.Stdin)

	var err error
	var num int

	for {
		fmt.Println("Ingrese un número: ")
		if scanner.Scan(){
			num, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d \n", num, i, num*i)
	}
}