package strings

// Chain erwartet einen String und hängt ihn n mal hintereinander.
// Liefert das Ergebnis.
func Chain(s string, n int) string {
	var chain string
	for i := 0; i < n; i++ {
		chain += s
	}
	return chain
}
