package search

// FindSorted sucht in einer sortierten Liste nach dem ersten Vorkommen von x.
// Falls x nicht gefunden wird, wird -1 zurückgegeben.
// Da die Liste sortiert ist, wird die binäre Suche verwendet.
func FindSorted(list []int, x int) int {
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
	offset := 0 // Offset für den Index in der Teilliste
	//Fall 3 (wenn x größer als wert an mitte)
	if x > list[c] {
		sublist = list[c+1:]
		offset = c + 1
	}
	return offset + FindSorted(sublist, x)
}
