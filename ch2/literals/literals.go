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

	var firstValue string = "Askar"
	var secondValue string = "Adilet"
	var finalValue string = firstValue + secondValue

	var myFirstInitial rune = 'A'

	fmt.Println(finalValue[0])
	fmt.Println(myFirstInitial)

	// Type Conversions are not automatic
	var xt int = 10
	var yt float64 = 30.2

	var sum1 float64 = float64(xt) + yt
	var sum2 int = xt + int(yt)

	fmt.Println(sum1)
	fmt.Println(sum2)

	var xd int = 10
	var yd byte = 100

	var sum3 int = xd + int(yd)

	fmt.Println(sum3)

	xs := 10

	fmt.Println(xs)
}
