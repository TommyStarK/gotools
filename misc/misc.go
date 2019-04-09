package misc

func HaveOppositeSigns(a, b int) bool {
	return ((a ^ b) < 0)
}

func Min(a, b int) (min int) {
	min = b ^ ((a ^ b) & ((a - b) >> 31))
	return
}

func Max(a, b int) (max int) {
	max = a ^ ((b ^ a) & ((a - b) >> 31))
	return
}
