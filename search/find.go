package search

import "github.com/tel23a-inf/aufgaben-rekursion/lists"

// Find sucht in einer Liste nach dem ersten Vorkommen von x
// und gibt dessen Index zurück. Falls x nicht gefunden wird,
// wird -1 zurückgegeben.
func Find(list []int, x int) int {
	// for i := 0; i < len(list); i++ {
	// 	if list[i] == x {
	// 		return i
	// 	}
	// }
	// return -1

	if lists.Empty(list) {
		return -1
	}

	if list[0] == x {
		return 0
	}

	pos := Find(list[1:], x)
	//Abbruchbedingung für pos < 0, pos wird -1, wenn es keine Schnittmenge zwischen der liste und x gibt
	if pos == -1 {
		return -1
	}

	return 1 + Find(list[1:], x)
}
