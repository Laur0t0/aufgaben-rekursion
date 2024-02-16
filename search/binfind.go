package search

// FindSorted sucht in einer sortierten Liste nach dem ersten Vorkommen von x.
// Falls x nicht gefunden wird, wird -1 zurückgegeben.
// Da die Liste sortiert ist, wird die binäre Suche verwendet.
func FindSorted(list []int, x int) int {
	var offset int
	//Fall 1 (wenn Liste leer ist)
	if len(list) == 0 {
		return -1
	}
	c := len(list) / 2
	//Fall 2 (wenn wert an mitte genau x)
	if list[c] == x {
		return c
	}
	sublist := list[:c]
	//offset = 0 // Offset für den Index in der Teilliste
	//Fall 3 (wenn x größer als wert an mitte)
	if x > list[c] {
		sublist = list[c+1:]
		offset = c + 1
	}
	pos := FindSorted(sublist, x)
	//Abbruchbedingung für pos < 0, pos wird -1, wenn es keine Schnittmenge zwischen der Teilliste (sublist) und x gibt
	if pos == -1 {
		return -1
	}
	return offset
}

// Test

// l1 := []int{1, 3, 4, 6, 10, 21, 38}

// 	fmt.Println(FindSorted(l1, 10))

// 	// Output:
// 	// 4
