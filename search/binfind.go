package search

// FindSorted sucht in einer sortierten Liste nach dem ersten Vorkommen von x.
// Falls x nicht gefunden wird, wird -1 zurückgegeben.
// Da die Liste sortiert ist, wird die binäre Suche verwendet.
func FindSorted(list []int, x int) int {
	if len(list) == 0 {
		return -1
	}
	c := len(list) / 2
	if list[c] == x {
		return c
	}
	sublist := list[:c]
	offset := 0 // Offset für den Index in der Teilliste
	if x > list[c] {
		sublist = list[c+1:]
		offset = c + 1
	}
	pos := FindSorted(sublist, x)
	if pos == -1 {
		return -1
	}
	return pos + offset
}
