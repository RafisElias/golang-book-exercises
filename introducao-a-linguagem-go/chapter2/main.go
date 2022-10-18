//* Tipos

package main

import (
	"fmt"
)

func main() {
	// Number
	// Int
	fmt.Println("1 + 1 =", 1+1)
	// Float
	fmt.Println("1.0 + 1.0 =", 1.0+1.0)

	// String
	fmt.Println(len("Hello, World"))
	fmt.Println("Hello, World"[1])
	fmt.Println("Hello" + " World")

	// Boolean
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && true)
	fmt.Println(false && false)

	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(false || false)

	fmt.Println(!true)
	fmt.Println(!false)
}
