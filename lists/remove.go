package lists

// Liefert eine Liste, die alle Elemente aus list enthält,
// außer dem an Stelle pos.
// Wenn pos außerhalb des Bereichs der Liste liegt, wird die
// ursprüngliche Liste zurückgegeben.
// Verwenden Sie Rekursion und benutzen Sie NICHT die len-Funktion.
// Sie können die Hilfsfunktion Empty aus empty.go verwenden.
func RemoveElement(list []int, pos int) []int {
	if Empty(list) {
		return list
	}
	if pos == 0 {
		return list[1:]
	}
	//return list[0]+RemoveElement(list[1:], pos-1)
	return append([]int{list[0]},RemoveElement(list[1:], pos-1)...)	//... ist ein Shortcut um alle Elemente der Liste abzutasten
	
	//return append([]int{list[0]},[]int{list[1]},[]int{list[2]},[]int{list[3]},[]int{list[4]})
}
