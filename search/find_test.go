package search

import "fmt"

func ExampleFind() {
	l1 := []int{1, 3, 4, 6, 10, 21, 38}

	fmt.Println(Find(l1, 10))
	fmt.Println(Find(l1, 9))

	// Output:
	// 4
	// -1

}
