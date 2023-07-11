package iteraciones

import (
	"fmt"
)

func Iterar() {
	for i := 0; i < 50; i+=5 {
		if i == 30 {
			continue
		}
		fmt.Println(i)
	}
}