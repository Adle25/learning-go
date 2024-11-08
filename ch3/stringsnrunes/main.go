package main

import "fmt"

func main() {
	var s string = "Hello there"
	var b byte = s[4]
	var exampleStr = "Hello"

	fmt.Println(b)
	fmt.Println(s[:4])
	fmt.Println(s[4:])
	fmt.Println(s[4:7])
	fmt.Println(len(exampleStr))
}
