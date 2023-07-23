package middleware

import "fmt"

func sumar(a int, b int) int {
	return a + b
}

func restar(a int, b int) int {
	return a - b
}

func multiplicar(a int, b int) int {
	return a * b
}

func MyMiddleware() {
	fmt.Println("1st")
	result := operacionMiddleware(sumar)(4,3)
	fmt.Println(result)
	fmt.Println("2nd")
	result = operacionMiddleware(restar)(5,7)
	fmt.Println(result)
	fmt.Println("3rd")
	result = operacionMiddleware(multiplicar)(6,4)
	fmt.Println(result)
}

func operacionMiddleware(f func(int, int) int) func(int, int) int {
	return func(a int, b int) int {
		fmt.Println("Pre proceso la información...")
		return f(a, b)
	}
}