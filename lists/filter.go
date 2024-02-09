package lists

// Liefert eine Liste mit allen Elementen aus list, die kleiner oder gleich key sind.
// Verwenden Sie Rekursion und benutzen Sie NICHT die len-Funktion.
// Sie können die Hilfsfunktion Empty aus empty.go verwenden.
func FilterLess(list []int, key int) []int {
	// Gehen Sie ähnlich wie bei Remove vor:
	// Wenn die Liste leer ist, ist das Ergebnis die leere Liste.
	// Wenn das erste Element größer als key ist, ist das Ergebnis die gefilterte Restliste.
	// Wenn das erste Element kleiner oder gleich key ist, ist das Ergebnis das erste Element
	// plus die gefilterte Restliste.

	if Empty(list) {
		return list
	}
	if key < list[0] {
		return FilterLess(list[1:], key)
	}
	//return list[0]+RemoveElement(list[1:], pos-1)
	return append([]int{list[0]}, FilterLess(list[1:], key)...)

}

// Liefert eine Liste mit allen Elementen aus list, die echt größer als key sind.
// Verwenden Sie Rekursion und benutzen Sie NICHT die len-Funktion.
// Sie können die Hilfsfunktion Empty aus empty.go verwenden.
func FilterGreater(list []int, key int) []int {

	if Empty(list) {
		return list
	}
	if key >= list[0] {
		return FilterGreater(list[1:], key)
	}
	//return list[0]+RemoveElement(list[1:], pos-1)
	return append([]int{list[0]}, FilterGreater(list[1:], key)...)
}
