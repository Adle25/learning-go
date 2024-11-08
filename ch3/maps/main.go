package main

import "fmt"

func main() {
	// teams := map[string][]string{
	// 	"Orcas":   {"Fred", "Ralph", "Bijou"},
	// 	"Lions":   {"Sarah", "Peter", "Billie"},
	// 	"Kittens": {"Waldo", "Raul", "Ze"},
	// }

	// fmt.Println(teams)

	// totalWins := map[string]int{}
	// totalWins["Orcas"] = 1
	// totalWins["Lions"] = 2

	// fmt.Println(totalWins["Orcas"])
	// fmt.Println(totalWins["Kittens"])

	m := map[string]int{
		"hello": 5,
		"world": 0,
	}

	// v, ok := m["hello"]
	// fmt.Println(v, ok)

	delete(m, "hello")
	fmt.Println(m)

	clear(m)
	fmt.Println(m)
}
