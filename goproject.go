package main

import "fmt"

func swap(one, two string) (string, string) {
	return two, one
}
func main() {
	first, second := "Analized.", "Well"
	fmt.Println(swap(first, second))
}
