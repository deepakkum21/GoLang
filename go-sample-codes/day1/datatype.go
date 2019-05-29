package main

import "fmt"

// Go is a static-typed language, meaning you need to specify the type, and the type cannot change without you explicitly changing it.

func add(x float64, y float64) float64 {
	return x + y
}

func add1(x, y float64) float64 {
	return x + y
}

func add2(x, y float32) float32 {
	return x + y
}

// funstion with multiple return types
func multiple(a, b string) (string, string) {
	return a, b
}

func main() {
	var num1 float64 = 5.6
	var num2 float64 = 9.5
	var num3, num4 float64 = 15.6, 19.5

	// Outside of your functions, you will need to define explicitly data types. Inside your functions, Go can figure out your types when compiled, but the assigned types *cannot* be changed without being explicitly converted.
	num5, num6 := 5.6, 9.5

	fmt.Println(add(num1, num2))
	fmt.Println(add1(num3, num4))
	fmt.Print(add(num5, num6))

	// throws error --> cannot use num1 (type float64) as type float32 in argument to add
	// fmt.Print(add2(num5, num6))

	w1, w2 := multiple("Hey", "there")
	fmt.Println(w1, w2)

	// If you wanted to convert the types of a variable:
	var a int = 62
	var b float64 = float64(a)

	// type inference works in Go:
	var x float32
	y := x // y is float32 type
}

// bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte --> alias for uint8

// rune --> alias for int32
// represents a Unicode code point

// float32 float64

// complex64 complex128
