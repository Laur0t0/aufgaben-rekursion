package calc

// BinomialCoefficient erwartet zwei Zahlen n und k und liefert den
// Binomialkoeffizienten "n über k".
func BinomialCoefficient(n, k int) int {
	return fac(n) / (fac(k) * fac(n-k))
}

func fac(i int) int {
	if i == 0 {
		return 1
	}
	return i * fac(i-1)
}
