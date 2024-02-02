package calc

// Liefert das Produkt von x und y.
func Mult(x, y int) int {
	if y == 0 {
		return 0
	}

	if x < 0 {
		return -Mult(-x, y)
	}

	if y < 0 {
		return -Mult(x, -y)
	}

	return Mult(x, y-1) + x
}
