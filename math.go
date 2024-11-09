package main

import "fmt"

func main() {
	result := soma(10, 10)

	fmt.Println(result)
}

func soma(a int, b int) int {
	return a + b
}
