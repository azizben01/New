package main

import (
	"fmt"
)

func operation(a, b, c int) int {
	return (a + b) / c
}
func main() {
	fmt.Println(operation(5, 1, 3))
}
