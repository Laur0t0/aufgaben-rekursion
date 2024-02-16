package lists

// Liefert zwei Listen:
// - Eine, die alle Elemente aus list enthält, die kleiner oder gleich key sind.
// - Eine, die alle übrigen Elemente aus list enthält.
func Partition(list []int, key int) ([]int, []int) {
	// Verwende Kopien von list, damit die ursprüngliche Liste nicht verändert wird.
	l1 := append([]int{}, FilterLess(list, key)...)
	l2 := append([]int{}, FilterGreater(list, key)...)

	return l1, l2
}
