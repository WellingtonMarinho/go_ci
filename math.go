package main

import "fmt"

func main() {
	result := Soma(10, 10)

	fmt.Println(result)
}

func Soma(a int, b int) int {
	return a + b
}
