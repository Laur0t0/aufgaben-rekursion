package strings

// Contains prüft, ob der String s die Sequenz seq enthält.
func Contains(s, seq string) bool {
	n := 0
	m := n + len(seq)
	if len(s) < len(seq) {
		return false
	}
	if s[n:m] == seq {
		return true
	}
	return Contains(s[1:], seq)
}
