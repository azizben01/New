//*package main

//import "fmt"

//func main() {
//	fmt.Printf("HOLLA")
//}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	number := rand.Intn(10)
	fmt.Println(number)
	fmt.Printf("The time is", time.Now())

}
