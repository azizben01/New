/*package main

import "fmt"

func swap(one, two string) (string, string) {
	return two, one
}
func main() {
	first, second := "Analized.", "Well"
	fmt.Println(swap(first, second))
}
*/

package main

import (
	"fmt"
)

func reverseString(input string) string {
	runes := []rune(input) // Convert the string to a slice of runes
	length := len(runes)

	// Swap characters from the beginning and end of the string
	for i := 0; i < length/2; i++ {
		runes[i], runes[length-i-1] = runes[length-i-1], runes[i]
	}

	return string(runes) // Convert the slice of runes back to a string
}

func main() {
	fmt.Print("Enter a string: ")
	var input string
	fmt.Scanln(&input)

	reversed := reverseString(input)
	fmt.Println("Reversed string:", reversed)
}
