package main

import (
	"fmt"
	"github.com/elEdupown/go_testing/variables"
)

func main() {
	estado, texto := variables.ConvertToText(1023)
	fmt.Println(estado, texto)
}