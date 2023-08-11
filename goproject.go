/*package main

import "fmt"

func Compare(a, b int) string {
	if a > b {
		return "a is greater than b"
	}
	return "b is greater than a"
}
func main() {
	fmt.Println(Compare(12, 40))
}
*/

package main

import "fmt"

func reverseStringWithIf(input string) string {
	reversed := ""
	length := len(input)

	for i := length - 1; i >= 0; i-- {
		reversed += string(input[i])
	}

	return reversed
}

func main() {
	input := "hello"
	reversed := reverseStringWithIf(input)
	fmt.Println(reversed) // Output: "olleh"
}
