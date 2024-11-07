package main

import "fmt"

func main() {
	// var x = []int{1, 2, 3}
	// var y [][]int
	// var z []int

	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(z)
	// fmt.Println(z == nil)

	// firstSlice := []int{1, 2, 3, 4, 5}
	// secondSlice := []int{1, 2, 3, 4, 5}
	// thirdSlice := []int{1, 2, 3, 4, 5, 6}

	// fmt.Println(slices.Equal(firstSlice, secondSlice))
	// fmt.Println(slices.Equal(firstSlice, thirdSlice))
	// fmt.Println(len(thirdSlice))

	// thirdSlice = append(thirdSlice, 10)
	// fmt.Println(thirdSlice)

	// fourthSlice := append(firstSlice, secondSlice...)
	// fmt.Println(fourthSlice)

	// var testSlice []int
	// fmt.Println(len(testSlice), cap(testSlice))

	// testSlice = append(testSlice, 2)
	// fmt.Println(len(testSlice), cap(testSlice))

	// testSlice = append(testSlice, 3)
	// fmt.Println(len(testSlice), cap(testSlice))

	// testSlice = append(testSlice, 4)
	// fmt.Println(len(testSlice), cap(testSlice))

	// testSlice = append(testSlice, 5)
	// fmt.Println(len(testSlice), cap(testSlice))

	// var mySlice = make([]int, 5)

	// fmt.Println(mySlice)

	// anotherSlice := []int{1, 2, 3}
	// clear(anotherSlice)
	// fmt.Println(anotherSlice)

	slicingSlice := []string{"a", "b", "c", "d", "e"}

	fmt.Println(slicingSlice[:2])
	fmt.Println(slicingSlice[1:])
	fmt.Println(slicingSlice[:])
}
