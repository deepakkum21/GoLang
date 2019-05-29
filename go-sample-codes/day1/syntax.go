package main

// use this style for the importing multiple library
import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func foo() {
	// the first letter of the function which we are using of the library is capital
	fmt.Println("The square root of 4 is", math.Sqrt(4))

	fmt.Println("It is currently:", time.Now())
	fmt.Println("A number from 1-100", rand.Intn(100))
}

func main() {
	foo()
}
