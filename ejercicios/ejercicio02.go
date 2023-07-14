package ejercicios

import (
	"fmt"
	"strconv"
	"bufio"
	"os"
)


func Tabla() string{
	scanner := bufio.NewScanner(os.Stdin)

	var err error
	var num int
	var text string

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
		text += fmt.Sprintf("%d x %d = %d \n", num, i, num*i)
	}

	return text
}