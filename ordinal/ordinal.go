// Package ordinal provides English ordinals for numbers.
package ordinal

// For returns the English ordinal for the given number.
func For(n int) string {
	mod := n % 100
	if mod >= 4 && mod <= 20 {
		return "th"
	}
	switch n % 10 {
	case 1:
		return "st"
	case 2:
		return "nd"
	case 3:
		return "rd"
	default:
		return "th"
	}
}
