package main

import "fmt"

func main() {
	// Boolean types
	var flag bool // No value assigned so it is false
	var isAvesome = true

	// Numeric types
	var smallInt int8 = -128
	var unsignedInt uint8 = 2
	var byteInt byte = 234

	var x int = 10
	x += 2

	var floatSmall float32 = 32.2323
	var floatBig float64 = 345.345345345345

	y := complex(2.5, 3.1)
	z := complex(10.2, 2)

	fmt.Println(y + z)
	fmt.Println(y - z)
	fmt.Println(y * z)
	fmt.Println(y / z)
	fmt.Println(y - z)
	fmt.Println(real(z))
	fmt.Println(imag(z))

	fmt.Println(flag)
	fmt.Println(isAvesome)
	fmt.Println(smallInt)
	fmt.Println(unsignedInt)
	fmt.Println(byteInt)
	fmt.Println(floatSmall)
	fmt.Println(floatBig)

}
