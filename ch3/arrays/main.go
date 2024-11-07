package main

import "fmt"

func main() {
	// var x [3]int
	// var y = [3]int{10, 20, 30}
	// var z = [...]int{20, 30, 44, 40, 44}

	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(z)

	var firstArray = [...]int{1, 2, 3}
	var secondArray = [3]int{1, 2, 3}

	firstArray[0] = 3

	fmt.Println(firstArray == secondArray)
	fmt.Println(firstArray[0])
	fmt.Println(len(firstArray))
}
